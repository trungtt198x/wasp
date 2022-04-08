// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreaccounts

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "accounts"
	ScDescription = "Core chain account ledger contract"
	HScName       = wasmtypes.ScHname(0x3c4b5e02)
)

const (
	ParamAgentID        = "a"
	ParamWithdrawAmount = "m"
)

const (
	ResultAccountNonce = "n"
	ResultAgents       = "this"
)

const (
	FuncDeposit         = "deposit"
	FuncHarvest         = "harvest"
	FuncWithdraw        = "withdraw"
	ViewAccounts        = "accounts"
	ViewGetAccountNonce = "getAccountNonce"
)

const (
	HFuncDeposit         = wasmtypes.ScHname(0xbdc9102d)
	HFuncHarvest         = wasmtypes.ScHname(0x7b40efbd)
	HFuncWithdraw        = wasmtypes.ScHname(0x9dcc0f41)
	HViewAccounts        = wasmtypes.ScHname(0x3c4b5e02)
	HViewGetAccountNonce = wasmtypes.ScHname(0x529d7df9)
)
