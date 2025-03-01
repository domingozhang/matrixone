// Copyright 2022 Matrix Origin
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

package onduplicatekey

import (
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/sql/plan"
	"github.com/matrixorigin/matrixone/pkg/vm/engine"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const (
	Build = iota
	Eval
	End
)

type proc = process.Process

type container struct {
	colexec.ReceiverOperator

	state            int
	checkConflictBat *batch.Batch //batch to check conflict
	insertBat        *batch.Batch //the final batch
}

type Argument struct {
	// Ts is not used
	Ts       uint64
	Affected uint64
	Engine   engine.Engine

	// Source       engine.Relation
	// UniqueSource []engine.Relation
	// Ref          *plan.ObjectRef
	TableDef        *plan.TableDef
	OnDuplicateIdx  []int32
	OnDuplicateExpr map[string]*plan.Expr

	IdxIdx []int32

	ctr *container

	IsIgnore bool
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	if arg.ctr != nil {
		arg.ctr.FreeMergeTypeOperator(pipelineFailed)

		if arg.ctr.insertBat != nil {
			arg.ctr.insertBat.Clean(proc.GetMPool())
		}
		if arg.ctr.checkConflictBat != nil {
			arg.ctr.checkConflictBat.Clean(proc.GetMPool())
		}
	}
}
