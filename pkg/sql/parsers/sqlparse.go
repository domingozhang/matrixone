// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parsers

import (
	"context"
	gotrace "runtime/trace"
	"strings"

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/dialect"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/dialect/mysql"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/dialect/postgresql"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/tree"
)

func Parse(ctx context.Context, dialectType dialect.DialectType, sql string, lower int64) ([]tree.Statement, error) {
	_, task := gotrace.NewTask(context.TODO(), "parser.Parse")
	defer task.End()
	switch dialectType {
	case dialect.MYSQL:
		return mysql.Parse(ctx, sql, lower)
	case dialect.POSTGRESQL:
		return postgresql.Parse(ctx, sql)
	default:
		return nil, moerr.NewInternalError(ctx, "type of dialect error")
	}
}

func ParseOne(ctx context.Context, dialectType dialect.DialectType, sql string, lower int64) (tree.Statement, error) {
	switch dialectType {
	case dialect.MYSQL:
		return mysql.ParseOne(ctx, sql, lower)
	case dialect.POSTGRESQL:
		return postgresql.ParseOne(ctx, sql)
	default:
		return nil, moerr.NewInternalError(ctx, "type of dialect error")
	}
}

const (
	stripCloudUser    = "/* cloud_user */"
	stripCloudNonUser = "/* cloud_nonuser */"
	stripSaveQuery    = "/* save_result */"
)

var HandleSqlForRecord = func(sql string) []string {
	split := SplitSqlBySemicolon(sql)
	for i := range split {
		//!!! remove method here assumes that the format of stripCloudUser or stripCloudNonUser
		// can not be changed, otherwise, the following code will not work.
		// It is case-sensitive and error-prone also.

		// Remove /* cloud_user */ prefix
		p0 := strings.Index(split[i], stripCloudUser)
		if p0 >= 0 {
			split[i] = split[i][0:p0] + split[i][p0+len(stripCloudUser):len(split[i])]
		}

		// remove /* cloud_nonuser */ prefix
		p0 = strings.Index(split[i], stripCloudNonUser)
		if p0 >= 0 {
			split[i] = split[i][0:p0] + split[i][p0+len(stripCloudNonUser):len(split[i])]
		}

		// remove /* save_query */ prefix
		p0 = strings.Index(split[i], stripSaveQuery)
		if p0 >= 0 {
			split[i] = split[i][0:p0] + split[i][p0+len(stripCloudNonUser):len(split[i])]
		}

		// Hide secret key for split[i],
		// for example:
		// before: create account nihao admin_name 'admin' identified with '123'
		// after: create account nihao admin_name 'admin' identified with '******'

		// Slice indexes helps to get the final ranges from split[i],
		// for example:
		// Secret keys' indexes ranges in split[i] are:
		// 1, 2, 3, 3
		// These mean [1, 2] and [3, 3] in split[i] are secret keys
		// And if len(split[i]) is 10, then we get slice indexes:
		// -1, 1, 2, 3, 3, 10
		// These mean we need to get (-1, 1), (2, 3), (3, 10) from split[i]
		scanner := mysql.NewScanner(dialect.MYSQL, split[i])
		indexes := []int{-1}
		eq := int('=')
		for scanner.Pos < len(split[i]) {
			typ, s := scanner.Scan()
			if typ == mysql.IDENTIFIED {
				typ, _ = scanner.Scan()
				if typ == mysql.BY || typ == mysql.WITH {
					typ, s = scanner.Scan()
					if typ != mysql.RANDOM {
						indexes = append(indexes, scanner.Pos-len(s)-1, scanner.Pos-2)
					}
				}
			} else if s == "access_key_id" || s == "secret_access_key" {
				typ, _ = scanner.Scan()
				if typ == eq {
					_, s = scanner.Scan()
					indexes = append(indexes, scanner.Pos-len(s)-1, scanner.Pos-2)
				}
			}
		}
		indexes = append(indexes, len(split[i]))

		if len(indexes) > 2 {
			var builder strings.Builder
			for j := 0; j < len(indexes); j += 2 {
				builder.WriteString(split[i][indexes[j]+1 : indexes[j+1]])
				if j < len(indexes)-2 {
					builder.WriteString("******")
				}
			}
			split[i] = builder.String()
		}
		split[i] = strings.TrimSpace(split[i])
	}
	return split
}

func SplitSqlBySemicolon(sql string) []string {
	var ret []string
	if len(sql) == 0 {
		// case 1 : "" => [""]
		return []string{sql}
	}
	scanner := mysql.NewScanner(dialect.MYSQL, sql)
	lastEnd := 0
	endWithSemicolon := false
	for scanner.Pos < len(sql) {
		typ, _ := scanner.Scan()
		for scanner.Pos < len(sql) && typ != ';' {
			typ, _ = scanner.Scan()
		}
		if typ == ';' {
			ret = append(ret, sql[lastEnd:scanner.Pos-1])
			lastEnd = scanner.Pos
			endWithSemicolon = true
		} else {
			ret = append(ret, sql[lastEnd:scanner.Pos])
			endWithSemicolon = false
		}
	}

	if len(ret) == 0 {
		//!!!NOTE there is at least one element in ret slice
		panic("there is at least one element")
	}
	//handle whitespace characters in the front and end of the sql
	for i := range ret {
		ret[i] = strings.TrimSpace(ret[i])
	}
	// do nothing
	//if len(ret) == 1 {
	//	//case 1 : "   " => [""]
	//	//case 2 : " abc " = > ["abc"]
	//	//case 3 : " /* abc */  " = > ["/* abc */"]
	//}
	if len(ret) > 1 {
		last := len(ret) - 1
		if !endWithSemicolon && len(ret[last]) == 0 {
			//case 3 : "abc;   " => ["abc"]
			//if the last one is end empty, remove it
			ret = ret[:last]
		}
		//case 4 : "abc; def; /* abc */  " => ["abc", "def", "/* abc */"]
	}

	return ret
}
