// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use crate::*;

#[derive(Clone)]
pub struct ImmutableFResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableFResults {
    pub fn n(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(RESULT_N))
	}
}

#[derive(Clone)]
pub struct MutableFResults {
	pub(crate) proxy: Proxy,
}

impl MutableFResults {
    pub fn n(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(RESULT_N))
	}
}
