// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use helloworld::*;
use consts::*;
use wasmlib::*;

mod helloworld;
mod consts;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_HELLO_WORLD, func_hello_world_thunk);
    exports.add_view(VIEW_GET_HELLO_WORLD, view_get_hello_world_thunk);
}

pub struct FuncHelloWorldParams {}

fn func_hello_world_thunk(ctx: &ScFuncContext) {
    let params = FuncHelloWorldParams {};
    func_hello_world(ctx, &params);
}

pub struct ViewGetHelloWorldParams {}

fn view_get_hello_world_thunk(ctx: &ScViewContext) {
    let params = ViewGetHelloWorldParams {};
    view_get_hello_world(ctx, &params);
}
