// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package schemacomment

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableTestFunc1Results struct {
	proxy wasmtypes.Proxy
}

// comment for length
func (s ImmutableTestFunc1Results) Length() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultLength))
}

type MutableTestFunc1Results struct {
	proxy wasmtypes.Proxy
}

// comment for length
func (s MutableTestFunc1Results) Length() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultLength))
}

type ImmutableTestView1Results struct {
	proxy wasmtypes.Proxy
}

// comment for length
func (s ImmutableTestView1Results) Length() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultLength))
}

type MutableTestView1Results struct {
	proxy wasmtypes.Proxy
}

// comment for length
func (s MutableTestView1Results) Length() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultLength))
}
