// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use testwasmlib::*;
use wasmlib::*;

use crate::consts::*;
use crate::events::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;
use crate::typedefs::*;

mod consts;
mod contract;
mod events;
mod params;
mod results;
mod state;
mod typedefs;

mod testwasmlib;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_ARRAY_APPEND,
    	FUNC_ARRAY_CLEAR,
    	FUNC_ARRAY_SET,
    	FUNC_MAP_CLEAR,
    	FUNC_MAP_SET,
    	FUNC_PARAM_TYPES,
    	FUNC_RANDOM,
    	FUNC_TRIGGER_EVENT,
    	VIEW_ARRAY_LENGTH,
    	VIEW_ARRAY_VALUE,
    	VIEW_BLOCK_RECORD,
    	VIEW_BLOCK_RECORDS,
    	VIEW_GET_RANDOM,
    	VIEW_IOTA_BALANCE,
    	VIEW_MAP_VALUE,
	],
    funcs: &[
    	func_array_append_thunk,
    	func_array_clear_thunk,
    	func_array_set_thunk,
    	func_map_clear_thunk,
    	func_map_set_thunk,
    	func_param_types_thunk,
    	func_random_thunk,
    	func_trigger_event_thunk,
	],
    views: &[
    	view_array_length_thunk,
    	view_array_value_thunk,
    	view_block_record_thunk,
    	view_block_records_thunk,
    	view_get_random_thunk,
    	view_iota_balance_thunk,
    	view_map_value_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct ArrayAppendContext {
	events:  TestWasmLibEvents,
	params: ImmutableArrayAppendParams,
	state: MutableTestWasmLibState,
}

fn func_array_append_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcArrayAppend");
	let f = ArrayAppendContext {
		events:  TestWasmLibEvents {},
		params: ImmutableArrayAppendParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.value().exists(), "missing mandatory value");
	func_array_append(ctx, &f);
	ctx.log("testwasmlib.funcArrayAppend ok");
}

pub struct ArrayClearContext {
	events:  TestWasmLibEvents,
	params: ImmutableArrayClearParams,
	state: MutableTestWasmLibState,
}

fn func_array_clear_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcArrayClear");
	let f = ArrayClearContext {
		events:  TestWasmLibEvents {},
		params: ImmutableArrayClearParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	func_array_clear(ctx, &f);
	ctx.log("testwasmlib.funcArrayClear ok");
}

pub struct ArraySetContext {
	events:  TestWasmLibEvents,
	params: ImmutableArraySetParams,
	state: MutableTestWasmLibState,
}

fn func_array_set_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcArraySet");
	let f = ArraySetContext {
		events:  TestWasmLibEvents {},
		params: ImmutableArraySetParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.index().exists(), "missing mandatory index");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.value().exists(), "missing mandatory value");
	func_array_set(ctx, &f);
	ctx.log("testwasmlib.funcArraySet ok");
}

pub struct MapClearContext {
	events:  TestWasmLibEvents,
	params: ImmutableMapClearParams,
	state: MutableTestWasmLibState,
}

fn func_map_clear_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcMapClear");
	let f = MapClearContext {
		events:  TestWasmLibEvents {},
		params: ImmutableMapClearParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	func_map_clear(ctx, &f);
	ctx.log("testwasmlib.funcMapClear ok");
}

pub struct MapSetContext {
	events:  TestWasmLibEvents,
	params: ImmutableMapSetParams,
	state: MutableTestWasmLibState,
}

fn func_map_set_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcMapSet");
	let f = MapSetContext {
		events:  TestWasmLibEvents {},
		params: ImmutableMapSetParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.key().exists(), "missing mandatory key");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.value().exists(), "missing mandatory value");
	func_map_set(ctx, &f);
	ctx.log("testwasmlib.funcMapSet ok");
}

pub struct ParamTypesContext {
	events:  TestWasmLibEvents,
	params: ImmutableParamTypesParams,
	state: MutableTestWasmLibState,
}

fn func_param_types_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcParamTypes");
	let f = ParamTypesContext {
		events:  TestWasmLibEvents {},
		params: ImmutableParamTypesParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	func_param_types(ctx, &f);
	ctx.log("testwasmlib.funcParamTypes ok");
}

pub struct RandomContext {
	events:  TestWasmLibEvents,
	state: MutableTestWasmLibState,
}

fn func_random_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcRandom");
	let f = RandomContext {
		events:  TestWasmLibEvents {},
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	func_random(ctx, &f);
	ctx.log("testwasmlib.funcRandom ok");
}

pub struct TriggerEventContext {
	events:  TestWasmLibEvents,
	params: ImmutableTriggerEventParams,
	state: MutableTestWasmLibState,
}

fn func_trigger_event_thunk(ctx: &ScFuncContext) {
	ctx.log("testwasmlib.funcTriggerEvent");
	let f = TriggerEventContext {
		events:  TestWasmLibEvents {},
		params: ImmutableTriggerEventParams { proxy: params_proxy() },
		state: MutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.address().exists(), "missing mandatory address");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	func_trigger_event(ctx, &f);
	ctx.log("testwasmlib.funcTriggerEvent ok");
}

pub struct ArrayLengthContext {
	params: ImmutableArrayLengthParams,
	results: MutableArrayLengthResults,
	state: ImmutableTestWasmLibState,
}

fn view_array_length_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewArrayLength");
	let f = ArrayLengthContext {
		params: ImmutableArrayLengthParams { proxy: params_proxy() },
		results: MutableArrayLengthResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	view_array_length(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewArrayLength ok");
}

pub struct ArrayValueContext {
	params: ImmutableArrayValueParams,
	results: MutableArrayValueResults,
	state: ImmutableTestWasmLibState,
}

fn view_array_value_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewArrayValue");
	let f = ArrayValueContext {
		params: ImmutableArrayValueParams { proxy: params_proxy() },
		results: MutableArrayValueResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.index().exists(), "missing mandatory index");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	view_array_value(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewArrayValue ok");
}

pub struct BlockRecordContext {
	params: ImmutableBlockRecordParams,
	results: MutableBlockRecordResults,
	state: ImmutableTestWasmLibState,
}

fn view_block_record_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewBlockRecord");
	let f = BlockRecordContext {
		params: ImmutableBlockRecordParams { proxy: params_proxy() },
		results: MutableBlockRecordResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.block_index().exists(), "missing mandatory blockIndex");
	ctx.require(f.params.record_index().exists(), "missing mandatory recordIndex");
	view_block_record(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewBlockRecord ok");
}

pub struct BlockRecordsContext {
	params: ImmutableBlockRecordsParams,
	results: MutableBlockRecordsResults,
	state: ImmutableTestWasmLibState,
}

fn view_block_records_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewBlockRecords");
	let f = BlockRecordsContext {
		params: ImmutableBlockRecordsParams { proxy: params_proxy() },
		results: MutableBlockRecordsResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.block_index().exists(), "missing mandatory blockIndex");
	view_block_records(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewBlockRecords ok");
}

pub struct GetRandomContext {
	results: MutableGetRandomResults,
	state: ImmutableTestWasmLibState,
}

fn view_get_random_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewGetRandom");
	let f = GetRandomContext {
		results: MutableGetRandomResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	view_get_random(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewGetRandom ok");
}

pub struct IotaBalanceContext {
	results: MutableIotaBalanceResults,
	state: ImmutableTestWasmLibState,
}

fn view_iota_balance_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewIotaBalance");
	let f = IotaBalanceContext {
		results: MutableIotaBalanceResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	view_iota_balance(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewIotaBalance ok");
}

pub struct MapValueContext {
	params: ImmutableMapValueParams,
	results: MutableMapValueResults,
	state: ImmutableTestWasmLibState,
}

fn view_map_value_thunk(ctx: &ScViewContext) {
	ctx.log("testwasmlib.viewMapValue");
	let f = MapValueContext {
		params: ImmutableMapValueParams { proxy: params_proxy() },
		results: MutableMapValueResults { proxy: results_proxy() },
		state: ImmutableTestWasmLibState { proxy: state_proxy() },
	};
	ctx.require(f.params.key().exists(), "missing mandatory key");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	view_map_value(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("testwasmlib.viewMapValue ok");
}
