// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type Token struct {
	// creation timestamp
	Created     uint64
	// description what minted token represents
	Description string
	// original minter
	MintedBy    wasmtypes.ScAgentID
	// current owner
	Owner       wasmtypes.ScAgentID
	// amount of tokens originally minted
	Supply      uint64
	// last update timestamp
	Updated     uint64
	// any user defined text
	UserDefined string
}

func NewTokenFromBytes(buf []byte) *Token {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Token{}
	data.Created     = wasmtypes.Uint64Decode(dec)
	data.Description = wasmtypes.StringDecode(dec)
	data.MintedBy    = wasmtypes.AgentIDDecode(dec)
	data.Owner       = wasmtypes.AgentIDDecode(dec)
	data.Supply      = wasmtypes.Uint64Decode(dec)
	data.Updated     = wasmtypes.Uint64Decode(dec)
	data.UserDefined = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *Token) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
	wasmtypes.Uint64Encode(enc, o.Created)
	wasmtypes.StringEncode(enc, o.Description)
	wasmtypes.AgentIDEncode(enc, o.MintedBy)
	wasmtypes.AgentIDEncode(enc, o.Owner)
	wasmtypes.Uint64Encode(enc, o.Supply)
	wasmtypes.Uint64Encode(enc, o.Updated)
	wasmtypes.StringEncode(enc, o.UserDefined)
	return enc.Buf()
}

type ImmutableToken struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableToken) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableToken) Value() *Token {
	return NewTokenFromBytes(o.proxy.Get())
}

type MutableToken struct {
	proxy wasmtypes.Proxy
}

func (o MutableToken) Delete() {
	o.proxy.Delete()
}

func (o MutableToken) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableToken) SetValue(value *Token) {
	o.proxy.Set(value.Bytes())
}

func (o MutableToken) Value() *Token {
	return NewTokenFromBytes(o.proxy.Get())
}
