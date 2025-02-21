// Copyright 2021 - 2022 Matrix Origin
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

package plan

import (
	"context"
	gotrace "runtime/trace"

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/pb/plan"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/dialect"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/tree"
)

func runBuildSelectByBinder(stmtType plan.Query_StatementType, ctx CompilerContext, stmt *tree.Select, isPrepareStmt bool) (*Plan, error) {
	builder := NewQueryBuilder(stmtType, ctx, isPrepareStmt)
	bindCtx := NewBindContext(builder, nil)
	rootId, err := builder.buildSelect(stmt, bindCtx, true)
	builder.qry.Steps = append(builder.qry.Steps, rootId)
	if err != nil {
		return nil, err
	}
	query, err := builder.createQuery()
	if err != nil {
		return nil, err
	}
	return &Plan{
		Plan: &plan.Plan_Query{
			Query: query,
		},
	}, err
}

func buildExplainAnalyze(ctx CompilerContext, stmt *tree.ExplainAnalyze, isPrepareStmt bool) (*Plan, error) {
	//get query optimizer and execute Optimize
	plan, err := BuildPlan(ctx, stmt.Statement, isPrepareStmt)
	if err != nil {
		return nil, err
	}
	if plan.GetQuery() == nil {
		return nil, moerr.NewNotSupported(ctx.GetContext(), "the sql query plan does not support explain.")
	}
	return plan, nil
}

func BuildPlan(ctx CompilerContext, stmt tree.Statement, isPrepareStmt bool) (*Plan, error) {
	_, task := gotrace.NewTask(context.TODO(), "plan.BuildPlan")
	defer task.End()
	switch stmt := stmt.(type) {
	case *tree.Select:
		return runBuildSelectByBinder(plan.Query_SELECT, ctx, stmt, isPrepareStmt)
	case *tree.ParenSelect:
		return runBuildSelectByBinder(plan.Query_SELECT, ctx, stmt.Select, isPrepareStmt)
	case *tree.ExplainAnalyze:
		return buildExplainAnalyze(ctx, stmt, isPrepareStmt)
	case *tree.Insert:
		return buildInsert(stmt, ctx, false, isPrepareStmt)
	case *tree.Replace:
		return buildReplace(stmt, ctx, isPrepareStmt)
	case *tree.Update:
		return buildTableUpdate(stmt, ctx, isPrepareStmt)
	case *tree.Delete:
		return buildDelete(stmt, ctx, isPrepareStmt)
	case *tree.BeginTransaction:
		return buildBeginTransaction(stmt, ctx)
	case *tree.CommitTransaction:
		return buildCommitTransaction(stmt, ctx)
	case *tree.RollbackTransaction:
		return buildRollbackTransaction(stmt, ctx)
	case *tree.CreateDatabase:
		return buildCreateDatabase(stmt, ctx)
	case *tree.DropDatabase:
		return buildDropDatabase(stmt, ctx)
	case *tree.CreateTable:
		return buildCreateTable(stmt, ctx)
	case *tree.DropTable:
		return buildDropTable(stmt, ctx)
	case *tree.TruncateTable:
		return buildTruncateTable(stmt, ctx)
	case *tree.CreateSequence:
		return buildCreateSequence(stmt, ctx)
	case *tree.DropSequence:
		return buildDropSequence(stmt, ctx)
	case *tree.AlterSequence:
		return buildAlterSequence(stmt, ctx)
	case *tree.DropView:
		return buildDropView(stmt, ctx)
	case *tree.CreateView:
		return buildCreateView(stmt, ctx)
	case *tree.CreateStream:
		return buildCreateStream(stmt, ctx)
	case *tree.AlterView:
		return buildAlterView(stmt, ctx)
	case *tree.AlterTable:
		return buildAlterTable(stmt, ctx)
	case *tree.CreateIndex:
		return buildCreateIndex(stmt, ctx)
	case *tree.DropIndex:
		return buildDropIndex(stmt, ctx)
	case *tree.ShowCreateDatabase:
		return buildShowCreateDatabase(stmt, ctx)
	case *tree.ShowCreateTable:
		return buildShowCreateTable(stmt, ctx)
	case *tree.ShowCreateView:
		return buildShowCreateView(stmt, ctx)
	case *tree.ShowDatabases:
		return buildShowDatabases(stmt, ctx)
	case *tree.ShowTables:
		return buildShowTables(stmt, ctx)
	case *tree.ShowSequences:
		return buildShowSequences(stmt, ctx)
	case *tree.ShowColumns:
		return buildShowColumns(stmt, ctx)
	case *tree.ShowTableStatus:
		return buildShowTableStatus(stmt, ctx)
	case *tree.ShowTarget:
		return buildShowTarget(stmt, ctx)
	case *tree.ShowIndex:
		return buildShowIndex(stmt, ctx)
	case *tree.ShowGrants:
		return buildShowGrants(stmt, ctx)
	case *tree.ShowCollation:
		return buildShowCollation(stmt, ctx)
	case *tree.ShowVariables:
		return buildShowVariables(stmt, ctx)
	case *tree.ShowStatus:
		return buildShowStatus(stmt, ctx)
	case *tree.ShowProcessList:
		return buildShowProcessList(ctx)
	case *tree.ShowLocks:
		return buildShowLocks(stmt, ctx)
	case *tree.ShowNodeList:
		return buildShowNodeList(stmt, ctx)
	case *tree.ShowFunctionOrProcedureStatus:
		return buildShowFunctionOrProcedureStatus(stmt, ctx)
	case *tree.ShowTableNumber:
		return buildShowTableNumber(stmt, ctx)
	case *tree.ShowColumnNumber:
		return buildShowColumnNumber(stmt, ctx)
	case *tree.ShowTableValues:
		return buildShowTableValues(stmt, ctx)
	case *tree.ShowRolesStmt:
		return buildShowRoles(stmt, ctx)
	case *tree.SetVar:
		return buildSetVariables(stmt, ctx)
	case *tree.Execute:
		return buildExecute(stmt, ctx)
	case *tree.Deallocate:
		return buildDeallocate(stmt, ctx)
	case *tree.Load:
		return buildLoad(stmt, ctx, isPrepareStmt)
	case *tree.PrepareStmt, *tree.PrepareString:
		return buildPrepare(stmt, ctx)
	case *tree.Do, *tree.Declare:
		return nil, moerr.NewNotSupported(ctx.GetContext(), tree.String(stmt, dialect.MYSQL))
	case *tree.ValuesStatement:
		return buildValues(stmt, ctx, isPrepareStmt)
	case *tree.LockTableStmt:
		return buildLockTables(stmt, ctx)
	case *tree.UnLockTableStmt:
		return buildUnLockTables(stmt, ctx)
	case *tree.ShowPublications:
		return buildShowPublication(stmt, ctx)
	case *tree.ShowCreatePublications:
		return buildShowCreatePublications(stmt, ctx)
	case *tree.ShowStages:
		return buildShowStages(stmt, ctx)
	default:
		return nil, moerr.NewInternalError(ctx.GetContext(), "statement: '%v'", tree.String(stmt, dialect.MYSQL))
	}
}

// GetExecType get executor will execute base AP or TP
func GetExecTypeFromPlan(pn *Plan) ExecInfo {
	defInfo := ExecInfo{
		Typ:        ExecTypeAP,
		WithGPU:    false,
		WithBigMem: false,
		CnNumbers:  2,
	}
	if IsTpQuery(pn.GetQuery()) {
		defInfo.Typ = ExecTypeTP
	}
	return defInfo
}

// GetResultColumnsFromPlan
func GetResultColumnsFromPlan(p *Plan) []*ColDef {
	getResultColumnsByProjectionlist := func(query *Query) []*ColDef {
		lastNode := query.Nodes[query.Steps[len(query.Steps)-1]]
		columns := make([]*ColDef, len(lastNode.ProjectList))
		for idx, expr := range lastNode.ProjectList {
			columns[idx] = &ColDef{
				Name: query.Headings[idx],
				Typ:  expr.Typ,
			}
		}

		return columns
	}

	switch logicPlan := p.Plan.(type) {
	case *plan.Plan_Query:
		switch logicPlan.Query.StmtType {
		case plan.Query_SELECT:
			return getResultColumnsByProjectionlist(logicPlan.Query)
		default:
			// insert/update/delete statement will return nil
			return nil
		}
	case *plan.Plan_Tcl:
		// begin/commmit/rollback statement will return nil
		return nil
	case *plan.Plan_Ddl:
		switch logicPlan.Ddl.DdlType {
		case plan.DataDefinition_SHOW_VARIABLES:
			typ := &plan.Type{
				Id:    int32(types.T_varchar),
				Width: 1024,
			}
			return []*ColDef{
				{Typ: typ, Name: "Variable_name"},
				{Typ: typ, Name: "Value"},
			}
		default:
			// show statement(except show variables) will return a query
			if logicPlan.Ddl.Query != nil {
				return getResultColumnsByProjectionlist(logicPlan.Ddl.Query)
			}
			return nil
		}
	}
	return nil
}
