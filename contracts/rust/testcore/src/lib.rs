// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use testcore::*;
use consts::*;
use wasmlib::*;

mod testcore;
mod consts;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_CALL_ON_CHAIN, func_call_on_chain_thunk);
    exports.add_func(FUNC_CHECK_CONTEXT_FROM_FULL_EP, func_check_context_from_full_ep_thunk);
    exports.add_func(FUNC_DO_NOTHING, func_do_nothing_thunk);
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_PASS_TYPES_FULL, func_pass_types_full_thunk);
    exports.add_func(FUNC_RUN_RECURSION, func_run_recursion_thunk);
    exports.add_func(FUNC_SEND_TO_ADDRESS, func_send_to_address_thunk);
    exports.add_func(FUNC_SET_INT, func_set_int_thunk);
    exports.add_func(FUNC_TEST_CALL_PANIC_FULL_EP, func_test_call_panic_full_ep_thunk);
    exports.add_func(FUNC_TEST_CALL_PANIC_VIEW_EPFROM_FULL, func_test_call_panic_view_epfrom_full_thunk);
    exports.add_func(FUNC_TEST_CHAIN_OWNER_IDFULL, func_test_chain_owner_idfull_thunk);
    exports.add_func(FUNC_TEST_CONTRACT_IDFULL, func_test_contract_idfull_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_DEPLOY, func_test_event_log_deploy_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_EVENT_DATA, func_test_event_log_event_data_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_GENERIC_DATA, func_test_event_log_generic_data_thunk);
    exports.add_func(FUNC_TEST_PANIC_FULL_EP, func_test_panic_full_ep_thunk);
    exports.add_func(FUNC_WITHDRAW_TO_CHAIN, func_withdraw_to_chain_thunk);
    exports.add_view(VIEW_CHECK_CONTEXT_FROM_VIEW_EP, view_check_context_from_view_ep_thunk);
    exports.add_view(VIEW_FIBONACCI, view_fibonacci_thunk);
    exports.add_view(VIEW_GET_COUNTER, view_get_counter_thunk);
    exports.add_view(VIEW_GET_INT, view_get_int_thunk);
    exports.add_view(VIEW_JUST_VIEW, view_just_view_thunk);
    exports.add_view(VIEW_PASS_TYPES_VIEW, view_pass_types_view_thunk);
    exports.add_view(VIEW_TEST_CALL_PANIC_VIEW_EPFROM_VIEW, view_test_call_panic_view_epfrom_view_thunk);
    exports.add_view(VIEW_TEST_CHAIN_OWNER_IDVIEW, view_test_chain_owner_idview_thunk);
    exports.add_view(VIEW_TEST_CONTRACT_IDVIEW, view_test_contract_idview_thunk);
    exports.add_view(VIEW_TEST_PANIC_VIEW_EP, view_test_panic_view_ep_thunk);
    exports.add_view(VIEW_TEST_SANDBOX_CALL, view_test_sandbox_call_thunk);
}

//@formatter:off
pub struct FuncCallOnChainParams {
    pub hname_contract: ScImmutableHname,
    pub hname_ep:       ScImmutableHname,
    pub int_value:      ScImmutableInt,
}
//@formatter:on

fn func_call_on_chain_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncCallOnChainParams {
        hname_contract: p.get_hname(PARAM_HNAME_CONTRACT),
        hname_ep: p.get_hname(PARAM_HNAME_EP),
        int_value: p.get_int(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    func_call_on_chain(ctx, &params);
}

//@formatter:off
pub struct FuncCheckContextFromFullEPParams {
    pub agent_id:         ScImmutableAgentId,
    pub caller:           ScImmutableAgentId,
    pub chain_id:         ScImmutableChainId,
    pub chain_owner_id:   ScImmutableAgentId,
    pub contract_creator: ScImmutableAgentId,
    pub contract_id:      ScImmutableContractId,
}
//@formatter:on

fn func_check_context_from_full_ep_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncCheckContextFromFullEPParams {
        agent_id: p.get_agent_id(PARAM_AGENT_ID),
        caller: p.get_agent_id(PARAM_CALLER),
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
        chain_owner_id: p.get_agent_id(PARAM_CHAIN_OWNER_ID),
        contract_creator: p.get_agent_id(PARAM_CONTRACT_CREATOR),
        contract_id: p.get_contract_id(PARAM_CONTRACT_ID),
    };
    ctx.require(params.agent_id.exists(), "missing mandatory agentId");
    ctx.require(params.caller.exists(), "missing mandatory caller");
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    ctx.require(params.chain_owner_id.exists(), "missing mandatory chainOwnerId");
    ctx.require(params.contract_creator.exists(), "missing mandatory contractCreator");
    ctx.require(params.contract_id.exists(), "missing mandatory contractId");
    func_check_context_from_full_ep(ctx, &params);
}

pub struct FuncDoNothingParams {}

fn func_do_nothing_thunk(ctx: &ScFuncContext) {
    let params = FuncDoNothingParams {};
    func_do_nothing(ctx, &params);
}

pub struct FuncInitParams {}

fn func_init_thunk(ctx: &ScFuncContext) {
    let params = FuncInitParams {};
    func_init(ctx, &params);
}

//@formatter:off
pub struct FuncPassTypesFullParams {
    pub hash:        ScImmutableHash,
    pub hname:       ScImmutableHname,
    pub hname_zero:  ScImmutableHname,
    pub int64:       ScImmutableInt,
    pub int64_zero:  ScImmutableInt,
    pub string:      ScImmutableString,
    pub string_zero: ScImmutableString,
}
//@formatter:on

fn func_pass_types_full_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncPassTypesFullParams {
        hash: p.get_hash(PARAM_HASH),
        hname: p.get_hname(PARAM_HNAME),
        hname_zero: p.get_hname(PARAM_HNAME_ZERO),
        int64: p.get_int(PARAM_INT64),
        int64_zero: p.get_int(PARAM_INT64_ZERO),
        string: p.get_string(PARAM_STRING),
        string_zero: p.get_string(PARAM_STRING_ZERO),
    };
    ctx.require(params.hash.exists(), "missing mandatory hash");
    ctx.require(params.hname.exists(), "missing mandatory hname");
    ctx.require(params.hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(params.int64.exists(), "missing mandatory int64");
    ctx.require(params.int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(params.string.exists(), "missing mandatory string");
    ctx.require(params.string_zero.exists(), "missing mandatory stringZero");
    func_pass_types_full(ctx, &params);
}

pub struct FuncRunRecursionParams {
    pub int_value: ScImmutableInt,
}

fn func_run_recursion_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncRunRecursionParams {
        int_value: p.get_int(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    func_run_recursion(ctx, &params);
}

pub struct FuncSendToAddressParams {
    pub address: ScImmutableAddress,
}

fn func_send_to_address_thunk(ctx: &ScFuncContext) {
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncSendToAddressParams {
        address: p.get_address(PARAM_ADDRESS),
    };
    ctx.require(params.address.exists(), "missing mandatory address");
    func_send_to_address(ctx, &params);
}

//@formatter:off
pub struct FuncSetIntParams {
    pub int_value: ScImmutableInt,
    pub name:      ScImmutableString,
}
//@formatter:on

fn func_set_int_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncSetIntParams {
        int_value: p.get_int(PARAM_INT_VALUE),
        name: p.get_string(PARAM_NAME),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    ctx.require(params.name.exists(), "missing mandatory name");
    func_set_int(ctx, &params);
}

pub struct FuncTestCallPanicFullEPParams {}

fn func_test_call_panic_full_ep_thunk(ctx: &ScFuncContext) {
    let params = FuncTestCallPanicFullEPParams {};
    func_test_call_panic_full_ep(ctx, &params);
}

pub struct FuncTestCallPanicViewEPFromFullParams {}

fn func_test_call_panic_view_epfrom_full_thunk(ctx: &ScFuncContext) {
    let params = FuncTestCallPanicViewEPFromFullParams {};
    func_test_call_panic_view_epfrom_full(ctx, &params);
}

pub struct FuncTestChainOwnerIDFullParams {}

fn func_test_chain_owner_idfull_thunk(ctx: &ScFuncContext) {
    let params = FuncTestChainOwnerIDFullParams {};
    func_test_chain_owner_idfull(ctx, &params);
}

pub struct FuncTestContractIDFullParams {}

fn func_test_contract_idfull_thunk(ctx: &ScFuncContext) {
    let params = FuncTestContractIDFullParams {};
    func_test_contract_idfull(ctx, &params);
}

pub struct FuncTestEventLogDeployParams {}

fn func_test_event_log_deploy_thunk(ctx: &ScFuncContext) {
    let params = FuncTestEventLogDeployParams {};
    func_test_event_log_deploy(ctx, &params);
}

pub struct FuncTestEventLogEventDataParams {}

fn func_test_event_log_event_data_thunk(ctx: &ScFuncContext) {
    let params = FuncTestEventLogEventDataParams {};
    func_test_event_log_event_data(ctx, &params);
}

pub struct FuncTestEventLogGenericDataParams {
    pub counter: ScImmutableInt,
}

fn func_test_event_log_generic_data_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncTestEventLogGenericDataParams {
        counter: p.get_int(PARAM_COUNTER),
    };
    ctx.require(params.counter.exists(), "missing mandatory counter");
    func_test_event_log_generic_data(ctx, &params);
}

pub struct FuncTestPanicFullEPParams {}

fn func_test_panic_full_ep_thunk(ctx: &ScFuncContext) {
    let params = FuncTestPanicFullEPParams {};
    func_test_panic_full_ep(ctx, &params);
}

pub struct FuncWithdrawToChainParams {
    pub chain_id: ScImmutableChainId,
}

fn func_withdraw_to_chain_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncWithdrawToChainParams {
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
    };
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    func_withdraw_to_chain(ctx, &params);
}

//@formatter:off
pub struct ViewCheckContextFromViewEPParams {
    pub agent_id:         ScImmutableAgentId,
    pub chain_id:         ScImmutableChainId,
    pub chain_owner_id:   ScImmutableAgentId,
    pub contract_creator: ScImmutableAgentId,
    pub contract_id:      ScImmutableContractId,
}
//@formatter:on

fn view_check_context_from_view_ep_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewCheckContextFromViewEPParams {
        agent_id: p.get_agent_id(PARAM_AGENT_ID),
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
        chain_owner_id: p.get_agent_id(PARAM_CHAIN_OWNER_ID),
        contract_creator: p.get_agent_id(PARAM_CONTRACT_CREATOR),
        contract_id: p.get_contract_id(PARAM_CONTRACT_ID),
    };
    ctx.require(params.agent_id.exists(), "missing mandatory agentId");
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    ctx.require(params.chain_owner_id.exists(), "missing mandatory chainOwnerId");
    ctx.require(params.contract_creator.exists(), "missing mandatory contractCreator");
    ctx.require(params.contract_id.exists(), "missing mandatory contractId");
    view_check_context_from_view_ep(ctx, &params);
}

pub struct ViewFibonacciParams {
    pub int_value: ScImmutableInt,
}

fn view_fibonacci_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewFibonacciParams {
        int_value: p.get_int(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    view_fibonacci(ctx, &params);
}

pub struct ViewGetCounterParams {}

fn view_get_counter_thunk(ctx: &ScViewContext) {
    let params = ViewGetCounterParams {};
    view_get_counter(ctx, &params);
}

pub struct ViewGetIntParams {
    pub name: ScImmutableString,
}

fn view_get_int_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewGetIntParams {
        name: p.get_string(PARAM_NAME),
    };
    ctx.require(params.name.exists(), "missing mandatory name");
    view_get_int(ctx, &params);
}

pub struct ViewJustViewParams {}

fn view_just_view_thunk(ctx: &ScViewContext) {
    let params = ViewJustViewParams {};
    view_just_view(ctx, &params);
}

//@formatter:off
pub struct ViewPassTypesViewParams {
    pub hash:        ScImmutableHash,
    pub hname:       ScImmutableHname,
    pub hname_zero:  ScImmutableHname,
    pub int64:       ScImmutableInt,
    pub int64_zero:  ScImmutableInt,
    pub string:      ScImmutableString,
    pub string_zero: ScImmutableString,
}
//@formatter:on

fn view_pass_types_view_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewPassTypesViewParams {
        hash: p.get_hash(PARAM_HASH),
        hname: p.get_hname(PARAM_HNAME),
        hname_zero: p.get_hname(PARAM_HNAME_ZERO),
        int64: p.get_int(PARAM_INT64),
        int64_zero: p.get_int(PARAM_INT64_ZERO),
        string: p.get_string(PARAM_STRING),
        string_zero: p.get_string(PARAM_STRING_ZERO),
    };
    ctx.require(params.hash.exists(), "missing mandatory hash");
    ctx.require(params.hname.exists(), "missing mandatory hname");
    ctx.require(params.hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(params.int64.exists(), "missing mandatory int64");
    ctx.require(params.int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(params.string.exists(), "missing mandatory string");
    ctx.require(params.string_zero.exists(), "missing mandatory stringZero");
    view_pass_types_view(ctx, &params);
}

pub struct ViewTestCallPanicViewEPFromViewParams {}

fn view_test_call_panic_view_epfrom_view_thunk(ctx: &ScViewContext) {
    let params = ViewTestCallPanicViewEPFromViewParams {};
    view_test_call_panic_view_epfrom_view(ctx, &params);
}

pub struct ViewTestChainOwnerIDViewParams {}

fn view_test_chain_owner_idview_thunk(ctx: &ScViewContext) {
    let params = ViewTestChainOwnerIDViewParams {};
    view_test_chain_owner_idview(ctx, &params);
}

pub struct ViewTestContractIDViewParams {}

fn view_test_contract_idview_thunk(ctx: &ScViewContext) {
    let params = ViewTestContractIDViewParams {};
    view_test_contract_idview(ctx, &params);
}

pub struct ViewTestPanicViewEPParams {}

fn view_test_panic_view_ep_thunk(ctx: &ScViewContext) {
    let params = ViewTestPanicViewEPParams {};
    view_test_panic_view_ep(ctx, &params);
}

pub struct ViewTestSandboxCallParams {}

fn view_test_sandbox_call_thunk(ctx: &ScViewContext) {
    let params = ViewTestSandboxCallParams {};
    view_test_sandbox_call(ctx, &params);
}
