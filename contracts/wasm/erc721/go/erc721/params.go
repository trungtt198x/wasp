// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc721

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableApproveParams struct {
	id int32
}

func (s ImmutableApproveParams) Approved() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamApproved))
}

func (s ImmutableApproveParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableApproveParams struct {
	id int32
}

func (s MutableApproveParams) Approved() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamApproved))
}

func (s MutableApproveParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableBurnParams struct {
	id int32
}

func (s ImmutableBurnParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableBurnParams struct {
	id int32
}

func (s MutableBurnParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Name() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamName])
}

func (s ImmutableInitParams) Symbol() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamSymbol])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Name() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamName])
}

func (s MutableInitParams) Symbol() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamSymbol])
}

type ImmutableMintParams struct {
	id int32
}

func (s ImmutableMintParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableMintParams struct {
	id int32
}

func (s MutableMintParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableSafeTransferFromParams struct {
	id int32
}

func (s ImmutableSafeTransferFromParams) Data() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, wasmlib.KeyID(ParamData))
}

func (s ImmutableSafeTransferFromParams) From() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamFrom))
}

func (s ImmutableSafeTransferFromParams) To() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamTo))
}

func (s ImmutableSafeTransferFromParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableSafeTransferFromParams struct {
	id int32
}

func (s MutableSafeTransferFromParams) Data() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, wasmlib.KeyID(ParamData))
}

func (s MutableSafeTransferFromParams) From() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamFrom))
}

func (s MutableSafeTransferFromParams) To() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamTo))
}

func (s MutableSafeTransferFromParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableSetApprovalForAllParams struct {
	id int32
}

func (s ImmutableSetApprovalForAllParams) Approval() wasmlib.ScImmutableBool {
	return wasmlib.NewScImmutableBool(s.id, wasmlib.KeyID(ParamApproval))
}

func (s ImmutableSetApprovalForAllParams) Operator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamOperator))
}

type MutableSetApprovalForAllParams struct {
	id int32
}

func (s MutableSetApprovalForAllParams) Approval() wasmlib.ScMutableBool {
	return wasmlib.NewScMutableBool(s.id, wasmlib.KeyID(ParamApproval))
}

func (s MutableSetApprovalForAllParams) Operator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamOperator))
}

type ImmutableTransferFromParams struct {
	id int32
}

func (s ImmutableTransferFromParams) From() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamFrom))
}

func (s ImmutableTransferFromParams) To() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamTo))
}

func (s ImmutableTransferFromParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableTransferFromParams struct {
	id int32
}

func (s MutableTransferFromParams) From() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamFrom))
}

func (s MutableTransferFromParams) To() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamTo))
}

func (s MutableTransferFromParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableBalanceOfParams struct {
	id int32
}

func (s ImmutableBalanceOfParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamOwner))
}

type MutableBalanceOfParams struct {
	id int32
}

func (s MutableBalanceOfParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamOwner))
}

type ImmutableGetApprovedParams struct {
	id int32
}

func (s ImmutableGetApprovedParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableGetApprovedParams struct {
	id int32
}

func (s MutableGetApprovedParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableIsApprovedForAllParams struct {
	id int32
}

func (s ImmutableIsApprovedForAllParams) Operator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamOperator))
}

func (s ImmutableIsApprovedForAllParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamOwner))
}

type MutableIsApprovedForAllParams struct {
	id int32
}

func (s MutableIsApprovedForAllParams) Operator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamOperator))
}

func (s MutableIsApprovedForAllParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamOwner))
}

type ImmutableOwnerOfParams struct {
	id int32
}

func (s ImmutableOwnerOfParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableOwnerOfParams struct {
	id int32
}

func (s MutableOwnerOfParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type ImmutableTokenURIParams struct {
	id int32
}

func (s ImmutableTokenURIParams) TokenID() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}

type MutableTokenURIParams struct {
	id int32
}

func (s MutableTokenURIParams) TokenID() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ParamTokenID))
}
