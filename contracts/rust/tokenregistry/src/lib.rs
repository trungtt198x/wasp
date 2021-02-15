// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use tokenregistry::*;
use consts::*;
use wasmlib::*;

mod tokenregistry;
mod consts;
mod types;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_MINT_SUPPLY, func_mint_supply_thunk);
    exports.add_func(FUNC_TRANSFER_OWNERSHIP, func_transfer_ownership_thunk);
    exports.add_func(FUNC_UPDATE_METADATA, func_update_metadata_thunk);
    exports.add_view(VIEW_GET_INFO, view_get_info_thunk);
}

//@formatter:off
pub struct FuncMintSupplyParams {
    pub description:  ScImmutableString, // description what minted token represents
    pub user_defined: ScImmutableString, // any user defined text
}
//@formatter:on

fn func_mint_supply_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncMintSupplyParams {
        description: p.get_string(PARAM_DESCRIPTION),
        user_defined: p.get_string(PARAM_USER_DEFINED),
    };
    func_mint_supply(ctx, &params);
}

pub struct FuncTransferOwnershipParams {
    pub color: ScImmutableColor, // color of token to transfer ownership of
}

fn func_transfer_ownership_thunk(ctx: &ScFuncContext) {
    //TODO the one who can transfer token ownership
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncTransferOwnershipParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    func_transfer_ownership(ctx, &params);
}

pub struct FuncUpdateMetadataParams {
    pub color: ScImmutableColor, // color of token to update metadata for
}

fn func_update_metadata_thunk(ctx: &ScFuncContext) {
    //TODO the one who can change the token info
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncUpdateMetadataParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    func_update_metadata(ctx, &params);
}

pub struct ViewGetInfoParams {
    pub color: ScImmutableColor, // color of token to view registry info of
}

fn view_get_info_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewGetInfoParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    view_get_info(ctx, &params);
}
