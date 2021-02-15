// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use dividend::*;
use consts::*;
use wasmlib::*;

mod dividend;
mod consts;
mod types;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DIVIDE, func_divide_thunk);
    exports.add_func(FUNC_MEMBER, func_member_thunk);
}

pub struct FuncDivideParams {}

fn func_divide_thunk(ctx: &ScFuncContext) {
    let params = FuncDivideParams {};
    func_divide(ctx, &params);
}

//@formatter:off
pub struct FuncMemberParams {
    pub address: ScImmutableAddress, // address of dividend recipient
    pub factor:  ScImmutableInt,     // relative division factor
}
//@formatter:on

fn func_member_thunk(ctx: &ScFuncContext) {
    // only creator can add members
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncMemberParams {
        address: p.get_address(PARAM_ADDRESS),
        factor: p.get_int(PARAM_FACTOR),
    };
    ctx.require(params.address.exists(), "missing mandatory address");
    ctx.require(params.factor.exists(), "missing mandatory factor");
    func_member(ctx, &params);
}
