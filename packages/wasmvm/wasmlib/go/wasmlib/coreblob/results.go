// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableStoreBlobResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStoreBlobResults) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ResultHash))
}

type MutableStoreBlobResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableStoreBlobResults) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ResultHash))
}

type ImmutableGetBlobFieldResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetBlobFieldResults) Bytes() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultBytes))
}

type MutableGetBlobFieldResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetBlobFieldResults) Bytes() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultBytes))
}

type MapStringToImmutableInt32 struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableInt32) GetInt32(key string) wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableGetBlobInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetBlobInfoResults) BlobSizes() MapStringToImmutableInt32 {
	//nolint:gosimple
	return MapStringToImmutableInt32{proxy: s.proxy}
}

type MapStringToMutableInt32 struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableInt32) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableInt32) GetInt32(key string) wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableGetBlobInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetBlobInfoResults) BlobSizes() MapStringToMutableInt32 {
	//nolint:gosimple
	return MapStringToMutableInt32{proxy: s.proxy}
}

type MapHashToImmutableInt32 struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableInt32) GetInt32(key wasmtypes.ScHash) wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(m.proxy.Key(wasmtypes.HashToBytes(key)))
}

type ImmutableListBlobsResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableListBlobsResults) BlobSizes() MapHashToImmutableInt32 {
	//nolint:gosimple
	return MapHashToImmutableInt32{proxy: s.proxy}
}

type MapHashToMutableInt32 struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableInt32) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableInt32) GetInt32(key wasmtypes.ScHash) wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(m.proxy.Key(wasmtypes.HashToBytes(key)))
}

type MutableListBlobsResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableListBlobsResults) BlobSizes() MapHashToMutableInt32 {
	//nolint:gosimple
	return MapHashToMutableInt32{proxy: s.proxy}
}
