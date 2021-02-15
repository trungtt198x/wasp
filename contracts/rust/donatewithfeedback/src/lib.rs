// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use donatewithfeedback::*;
use consts::*;
use wasmlib::*;

mod donatewithfeedback;
mod consts;
mod types;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DONATE, func_donate_thunk);
    exports.add_func(FUNC_WITHDRAW, func_withdraw_thunk);
    exports.add_view(VIEW_DONATIONS, view_donations_thunk);
}

pub struct FuncDonateParams {
    pub feedback: ScImmutableString, // feedback for the person you donate to
}

fn func_donate_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncDonateParams {
        feedback: p.get_string(PARAM_FEEDBACK),
    };
    func_donate(ctx, &params);
}

pub struct FuncWithdrawParams {
    pub amount: ScImmutableInt, // amount to withdraw
}

fn func_withdraw_thunk(ctx: &ScFuncContext) {
    // only SC creator can withdraw donated funds
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncWithdrawParams {
        amount: p.get_int(PARAM_AMOUNT),
    };
    func_withdraw(ctx, &params);
}

pub struct ViewDonationsParams {}

fn view_donations_thunk(ctx: &ScViewContext) {
    let params = ViewDonationsParams {};
    view_donations(ctx, &params);
}
