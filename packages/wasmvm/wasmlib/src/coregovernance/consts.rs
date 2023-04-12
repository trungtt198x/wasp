// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use crate::*;

pub const SC_NAME        : &str = "governance";
pub const SC_DESCRIPTION : &str = "Governance contract";
pub const HSC_NAME       : ScHname = ScHname(0x17cf909f);

pub(crate) const PARAM_ACCESS_API  : &str = "ia";
pub(crate) const PARAM_ACCESS_ONLY : &str = "i";
pub(crate) const PARAM_ACTIONS     : &str = "n";
pub(crate) const PARAM_ADDRESS     : &str = "S";
pub(crate) const PARAM_CERTIFICATE : &str = "ic";
pub(crate) const PARAM_CHAIN_OWNER : &str = "o";
pub(crate) const PARAM_FEE_POLICY  : &str = "g";
pub(crate) const PARAM_GAS_LIMITS  : &str = "l";
pub(crate) const PARAM_GAS_RATIO   : &str = "e";
pub(crate) const PARAM_METADATA    : &str = "x";
pub(crate) const PARAM_PUB_KEY     : &str = "ip";

pub(crate) const RESULT_ACCESS_NODE_CANDIDATES : &str = "an";
pub(crate) const RESULT_ACCESS_NODES           : &str = "ac";
pub(crate) const RESULT_CHAIN_ID               : &str = "c";
pub(crate) const RESULT_CHAIN_OWNER            : &str = "o";
pub(crate) const RESULT_CHAIN_OWNER_ID         : &str = "o";
pub(crate) const RESULT_CONTROLLERS            : &str = "a";
pub(crate) const RESULT_CUSTOM_METADATA        : &str = "x";
pub(crate) const RESULT_FEE_POLICY             : &str = "g";
pub(crate) const RESULT_GAS_LIMITS             : &str = "l";
pub(crate) const RESULT_GAS_RATIO              : &str = "e";
pub(crate) const RESULT_METADATA               : &str = "x";
pub(crate) const RESULT_STATUS                 : &str = "m";

pub(crate) const FUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS    : &str = "addAllowedStateControllerAddress";
pub(crate) const FUNC_ADD_CANDIDATE_NODE                      : &str = "addCandidateNode";
pub(crate) const FUNC_CHANGE_ACCESS_NODES                     : &str = "changeAccessNodes";
pub(crate) const FUNC_CLAIM_CHAIN_OWNERSHIP                   : &str = "claimChainOwnership";
pub(crate) const FUNC_DELEGATE_CHAIN_OWNERSHIP                : &str = "delegateChainOwnership";
pub(crate) const FUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS : &str = "removeAllowedStateControllerAddress";
pub(crate) const FUNC_REVOKE_ACCESS_NODE                      : &str = "revokeAccessNode";
pub(crate) const FUNC_ROTATE_STATE_CONTROLLER                 : &str = "rotateStateController";
pub(crate) const FUNC_SET_CUSTOM_METADATA                     : &str = "setCustomMetadata";
pub(crate) const FUNC_SET_EVM_GAS_RATIO                       : &str = "setEVMGasRatio";
pub(crate) const FUNC_SET_FEE_POLICY                          : &str = "setFeePolicy";
pub(crate) const FUNC_SET_GAS_LIMITS                          : &str = "setGasLimits";
pub(crate) const FUNC_START_MAINTENANCE                       : &str = "startMaintenance";
pub(crate) const FUNC_STOP_MAINTENANCE                        : &str = "stopMaintenance";
pub(crate) const VIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES  : &str = "getAllowedStateControllerAddresses";
pub(crate) const VIEW_GET_CHAIN_INFO                          : &str = "getChainInfo";
pub(crate) const VIEW_GET_CHAIN_NODES                         : &str = "getChainNodes";
pub(crate) const VIEW_GET_CHAIN_OWNER                         : &str = "getChainOwner";
pub(crate) const VIEW_GET_CUSTOM_METADATA                     : &str = "getCustomMetadata";
pub(crate) const VIEW_GET_EVM_GAS_RATIO                       : &str = "getEVMGasRatio";
pub(crate) const VIEW_GET_FEE_POLICY                          : &str = "getFeePolicy";
pub(crate) const VIEW_GET_GAS_LIMITS                          : &str = "getGasLimits";
pub(crate) const VIEW_GET_MAINTENANCE_STATUS                  : &str = "getMaintenanceStatus";

pub(crate) const HFUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS    : ScHname = ScHname(0x9469d567);
pub(crate) const HFUNC_ADD_CANDIDATE_NODE                      : ScHname = ScHname(0xb745b382);
pub(crate) const HFUNC_CHANGE_ACCESS_NODES                     : ScHname = ScHname(0x7bca3700);
pub(crate) const HFUNC_CLAIM_CHAIN_OWNERSHIP                   : ScHname = ScHname(0x03ff0fc0);
pub(crate) const HFUNC_DELEGATE_CHAIN_OWNERSHIP                : ScHname = ScHname(0x93ecb6ad);
pub(crate) const HFUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS : ScHname = ScHname(0x31f69447);
pub(crate) const HFUNC_REVOKE_ACCESS_NODE                      : ScHname = ScHname(0x5459512d);
pub(crate) const HFUNC_ROTATE_STATE_CONTROLLER                 : ScHname = ScHname(0x244d1038);
pub(crate) const HFUNC_SET_CUSTOM_METADATA                     : ScHname = ScHname(0xaa71b0b3);
pub(crate) const HFUNC_SET_EVM_GAS_RATIO                       : ScHname = ScHname(0xaae22338);
pub(crate) const HFUNC_SET_FEE_POLICY                          : ScHname = ScHname(0x5b791c9f);
pub(crate) const HFUNC_SET_GAS_LIMITS                          : ScHname = ScHname(0xd72fb355);
pub(crate) const HFUNC_START_MAINTENANCE                       : ScHname = ScHname(0x742f0521);
pub(crate) const HFUNC_STOP_MAINTENANCE                        : ScHname = ScHname(0x4e017b6a);
pub(crate) const HVIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES  : ScHname = ScHname(0xf3505183);
pub(crate) const HVIEW_GET_CHAIN_INFO                          : ScHname = ScHname(0x434477e2);
pub(crate) const HVIEW_GET_CHAIN_NODES                         : ScHname = ScHname(0xe1832289);
pub(crate) const HVIEW_GET_CHAIN_OWNER                         : ScHname = ScHname(0x9b2ef0ac);
pub(crate) const HVIEW_GET_CUSTOM_METADATA                     : ScHname = ScHname(0x02d4bac9);
pub(crate) const HVIEW_GET_EVM_GAS_RATIO                       : ScHname = ScHname(0xb81c8c34);
pub(crate) const HVIEW_GET_FEE_POLICY                          : ScHname = ScHname(0xf8c89790);
pub(crate) const HVIEW_GET_GAS_LIMITS                          : ScHname = ScHname(0x3a493455);
pub(crate) const HVIEW_GET_MAINTENANCE_STATUS                  : ScHname = ScHname(0x61fe5443);
