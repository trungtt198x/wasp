// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableGetFactorResults struct {
	proxy wasmtypes.Proxy
}

// relative division factor
func (s ImmutableGetFactorResults) Factor() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultFactor))
}

type MutableGetFactorResults struct {
	proxy wasmtypes.Proxy
}

// relative division factor
func (s MutableGetFactorResults) Factor() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultFactor))
}

type ImmutableGetOwnerResults struct {
	proxy wasmtypes.Proxy
}

// current owner of this smart contract
func (s ImmutableGetOwnerResults) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultOwner))
}

type MutableGetOwnerResults struct {
	proxy wasmtypes.Proxy
}

// current owner of this smart contract
func (s MutableGetOwnerResults) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultOwner))
}
