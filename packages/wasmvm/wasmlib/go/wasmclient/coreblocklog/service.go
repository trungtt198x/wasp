// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblocklogclient

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmclient"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

///////////////////////////// CoreBlockLogService /////////////////////////////

type CoreBlockLogService struct {
	wasmclient.Service
}

func NewCoreBlockLogService(cl *wasmclient.ServiceClient, chainID *wasmtypes.ScChainID) (*CoreBlockLogService, error) {
	s := &CoreBlockLogService{}
	err := s.Service.Init(cl, chainID, 0xf538ef2b)
	return s, err
}
