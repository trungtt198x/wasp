// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package tokenregistry

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		FuncMintSupply,
		FuncTransferOwnership,
		FuncUpdateMetadata,
		ViewGetInfo,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
		funcMintSupplyThunk,
		funcTransferOwnershipThunk,
		funcUpdateMetadataThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
		viewGetInfoThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type MintSupplyContext struct {
	Params  ImmutableMintSupplyParams
	State   MutableTokenRegistryState
}

func funcMintSupplyThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcMintSupply")
	f := &MintSupplyContext{
		Params: ImmutableMintSupplyParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTokenRegistryState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcMintSupply(ctx, f)
	ctx.Log("tokenregistry.funcMintSupply ok")
}

type TransferOwnershipContext struct {
	Params  ImmutableTransferOwnershipParams
	State   MutableTokenRegistryState
}

func funcTransferOwnershipThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcTransferOwnership")
	f := &TransferOwnershipContext{
		Params: ImmutableTransferOwnershipParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTokenRegistryState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// TODO the one who can transfer token ownership
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	ctx.Require(f.Params.Token().Exists(), "missing mandatory token")
	funcTransferOwnership(ctx, f)
	ctx.Log("tokenregistry.funcTransferOwnership ok")
}

type UpdateMetadataContext struct {
	Params  ImmutableUpdateMetadataParams
	State   MutableTokenRegistryState
}

func funcUpdateMetadataThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcUpdateMetadata")
	f := &UpdateMetadataContext{
		Params: ImmutableUpdateMetadataParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTokenRegistryState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// TODO the one who can change the token info
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	ctx.Require(f.Params.Token().Exists(), "missing mandatory token")
	funcUpdateMetadata(ctx, f)
	ctx.Log("tokenregistry.funcUpdateMetadata ok")
}

type GetInfoContext struct {
	Params  ImmutableGetInfoParams
	State   ImmutableTokenRegistryState
}

func viewGetInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("tokenregistry.viewGetInfo")
	f := &GetInfoContext{
		Params: ImmutableGetInfoParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: ImmutableTokenRegistryState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Token().Exists(), "missing mandatory token")
	viewGetInfo(ctx, f)
	ctx.Log("tokenregistry.viewGetInfo ok")
}
