// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::typedefs::*;

#[derive(Clone, Copy)]
pub struct ImmutableArrayLengthResults {
    pub(crate) id: i32,
}

impl ImmutableArrayLengthResults {
    pub fn length(&self) -> ScImmutableInt32 {
		ScImmutableInt32::new(self.id, RESULT_LENGTH.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableArrayLengthResults {
    pub(crate) id: i32,
}

impl MutableArrayLengthResults {
    pub fn length(&self) -> ScMutableInt32 {
		ScMutableInt32::new(self.id, RESULT_LENGTH.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableArrayValueResults {
    pub(crate) id: i32,
}

impl ImmutableArrayValueResults {
    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, RESULT_VALUE.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableArrayValueResults {
    pub(crate) id: i32,
}

impl MutableArrayValueResults {
    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.id, RESULT_VALUE.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableBlockRecordResults {
    pub(crate) id: i32,
}

impl ImmutableBlockRecordResults {
    pub fn record(&self) -> ScImmutableBytes {
		ScImmutableBytes::new(self.id, RESULT_RECORD.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableBlockRecordResults {
    pub(crate) id: i32,
}

impl MutableBlockRecordResults {
    pub fn record(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.id, RESULT_RECORD.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableBlockRecordsResults {
    pub(crate) id: i32,
}

impl ImmutableBlockRecordsResults {
    pub fn count(&self) -> ScImmutableInt32 {
		ScImmutableInt32::new(self.id, RESULT_COUNT.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableBlockRecordsResults {
    pub(crate) id: i32,
}

impl MutableBlockRecordsResults {
    pub fn count(&self) -> ScMutableInt32 {
		ScMutableInt32::new(self.id, RESULT_COUNT.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableGetRandomResults {
    pub(crate) id: i32,
}

impl ImmutableGetRandomResults {
    pub fn random(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_RANDOM.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableGetRandomResults {
    pub(crate) id: i32,
}

impl MutableGetRandomResults {
    pub fn random(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_RANDOM.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableIotaBalanceResults {
    pub(crate) id: i32,
}

impl ImmutableIotaBalanceResults {
    pub fn iotas(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_IOTAS.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableIotaBalanceResults {
    pub(crate) id: i32,
}

impl MutableIotaBalanceResults {
    pub fn iotas(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_IOTAS.get_key_id())
	}
}
