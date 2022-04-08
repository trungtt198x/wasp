// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package schemacomment

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type TestFunc1Call struct {
	Func    *wasmlib.ScFunc
	Params  MutableTestFunc1Params
	Results ImmutableTestFunc1Results
}

type TestView1Call struct {
	Func    *wasmlib.ScView
	Params  MutableTestView1Params
	Results ImmutableTestView1Results
}

type Funcs struct{}

var ScFuncs Funcs

// comment for TestFunc1 1
func (sc Funcs) TestFunc1(ctx wasmlib.ScFuncCallContext) *TestFunc1Call {
	f := &TestFunc1Call{Func: wasmlib.NewScFunc(ctx, HScName, HFuncTestFunc1)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	wasmlib.NewCallResultsProxy(&f.Func.ScView, &f.Results.proxy)
	return f
}

// comment for TestView1
func (sc Funcs) TestView1(ctx wasmlib.ScViewCallContext) *TestView1Call {
	f := &TestView1Call{Func: wasmlib.NewScView(ctx, HScName, HViewTestView1)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}
