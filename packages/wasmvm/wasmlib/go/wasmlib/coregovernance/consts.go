// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coregovernance

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "governance"
	ScDescription = "Core governance contract"
	HScName       = wasmtypes.ScHname(0x17cf909f)
)

const (
	ParamChainOwner               = "c"
	ParamChangeAccessNodesActions = "p"
	ParamFeePolicyBytes           = "d"
	ParamMaxBlobSize              = "g"
	ParamMaxEventSize             = "h"
	ParamMaxEventsPerReq          = "i"
	ParamStateControllerAddress   = "S"
)

const (
	ResultAccessNodeCandidates            = "ac"
	ResultAccessNodes                     = "an"
	ResultAllowedStateControllerAddresses = "b"
	ResultChainID                         = "c"
	ResultChainOwner                      = "c"
	ResultChainOwnerID                    = "o"
	ResultDescription                     = "d"
	ResultFeePolicyBytes                  = "d"
	ResultGasFeePolicyBytes               = "g"
	ResultMaxBlobSize                     = "mb"
	ResultMaxEventSize                    = "me"
	ResultMaxEventsPerReq                 = "mr"
)

const (
	FuncAddAllowedStateControllerAddress    = "addAllowedStateControllerAddress"
	FuncAddCandidateNode                    = "addCandidateNode"
	FuncChangeAccessNodes                   = "changeAccessNodes"
	FuncClaimChainOwnership                 = "claimChainOwnership"
	FuncDelegateChainOwnership              = "delegateChainOwnership"
	FuncRemoveAllowedStateControllerAddress = "removeAllowedStateControllerAddress"
	FuncRevokeAccessNode                    = "revokeAccessNode"
	FuncRotateStateController               = "rotateStateController"
	FuncSetChainInfo                        = "setChainInfo"
	FuncSetFeePolicy                        = "setFeePolicy"
	ViewGetAllowedStateControllerAddresses  = "getAllowedStateControllerAddresses"
	ViewGetChainInfo                        = "getChainInfo"
	ViewGetChainNodes                       = "getChainNodes"
	ViewGetChainOwner                       = "getChainOwner"
	ViewGetFeePolicy                        = "getFeePolicy"
	ViewGetMaxBlobSize                      = "getMaxBlobSize"
)

const (
	HFuncAddAllowedStateControllerAddress    = wasmtypes.ScHname(0x9469d567)
	HFuncAddCandidateNode                    = wasmtypes.ScHname(0xb745b382)
	HFuncChangeAccessNodes                   = wasmtypes.ScHname(0x7bca3700)
	HFuncClaimChainOwnership                 = wasmtypes.ScHname(0x03ff0fc0)
	HFuncDelegateChainOwnership              = wasmtypes.ScHname(0x93ecb6ad)
	HFuncRemoveAllowedStateControllerAddress = wasmtypes.ScHname(0x31f69447)
	HFuncRevokeAccessNode                    = wasmtypes.ScHname(0x5459512d)
	HFuncRotateStateController               = wasmtypes.ScHname(0x244d1038)
	HFuncSetChainInfo                        = wasmtypes.ScHname(0x702f5d2b)
	HFuncSetFeePolicy                        = wasmtypes.ScHname(0x5b791c9f)
	HViewGetAllowedStateControllerAddresses  = wasmtypes.ScHname(0xf3505183)
	HViewGetChainInfo                        = wasmtypes.ScHname(0x434477e2)
	HViewGetChainNodes                       = wasmtypes.ScHname(0xe1832289)
	HViewGetChainOwner                       = wasmtypes.ScHname(0x9b2ef0ac)
	HViewGetFeePolicy                        = wasmtypes.ScHname(0xf8c89790)
	HViewGetMaxBlobSize                      = wasmtypes.ScHname(0xe1db3d28)
)
