// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"
import * as sc from "./index";

export class ArrayOfImmutableBytes {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getBytes(index: i32): wasmlib.ScImmutableBytes {
        return new wasmlib.ScImmutableBytes(this.objID, new wasmlib.Key32(index));
    }
}

export class ImmutableGetAllowedStateControllerAddressesResults extends wasmlib.ScMapID {

    allowedStateControllerAddresses(): sc.ArrayOfImmutableBytes {
        let arrID = wasmlib.getObjectID(this.mapID, wasmlib.Key32.fromString(sc.ResultAllowedStateControllerAddresses), wasmlib.TYPE_ARRAY16|wasmlib.TYPE_BYTES);
        return new sc.ArrayOfImmutableBytes(arrID)
    }
}

export class ArrayOfMutableBytes {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getBytes(index: i32): wasmlib.ScMutableBytes {
        return new wasmlib.ScMutableBytes(this.objID, new wasmlib.Key32(index));
    }
}

export class MutableGetAllowedStateControllerAddressesResults extends wasmlib.ScMapID {

    allowedStateControllerAddresses(): sc.ArrayOfMutableBytes {
        let arrID = wasmlib.getObjectID(this.mapID, wasmlib.Key32.fromString(sc.ResultAllowedStateControllerAddresses), wasmlib.TYPE_ARRAY16|wasmlib.TYPE_BYTES);
        return new sc.ArrayOfMutableBytes(arrID)
    }
}

export class ImmutableGetChainInfoResults extends wasmlib.ScMapID {

    chainID(): wasmlib.ScImmutableChainID {
        return new wasmlib.ScImmutableChainID(this.mapID, wasmlib.Key32.fromString(sc.ResultChainID));
    }

    chainOwnerID(): wasmlib.ScImmutableAgentID {
        return new wasmlib.ScImmutableAgentID(this.mapID, wasmlib.Key32.fromString(sc.ResultChainOwnerID));
    }

    defaultOwnerFee(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultDefaultOwnerFee));
    }

    defaultValidatorFee(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultDefaultValidatorFee));
    }

    description(): wasmlib.ScImmutableString {
        return new wasmlib.ScImmutableString(this.mapID, wasmlib.Key32.fromString(sc.ResultDescription));
    }

    feeColor(): wasmlib.ScImmutableColor {
        return new wasmlib.ScImmutableColor(this.mapID, wasmlib.Key32.fromString(sc.ResultFeeColor));
    }

    maxBlobSize(): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxBlobSize));
    }

    maxEventSize(): wasmlib.ScImmutableInt16 {
        return new wasmlib.ScImmutableInt16(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxEventSize));
    }

    maxEventsPerReq(): wasmlib.ScImmutableInt16 {
        return new wasmlib.ScImmutableInt16(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxEventsPerReq));
    }
}

export class MutableGetChainInfoResults extends wasmlib.ScMapID {

    chainID(): wasmlib.ScMutableChainID {
        return new wasmlib.ScMutableChainID(this.mapID, wasmlib.Key32.fromString(sc.ResultChainID));
    }

    chainOwnerID(): wasmlib.ScMutableAgentID {
        return new wasmlib.ScMutableAgentID(this.mapID, wasmlib.Key32.fromString(sc.ResultChainOwnerID));
    }

    defaultOwnerFee(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultDefaultOwnerFee));
    }

    defaultValidatorFee(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultDefaultValidatorFee));
    }

    description(): wasmlib.ScMutableString {
        return new wasmlib.ScMutableString(this.mapID, wasmlib.Key32.fromString(sc.ResultDescription));
    }

    feeColor(): wasmlib.ScMutableColor {
        return new wasmlib.ScMutableColor(this.mapID, wasmlib.Key32.fromString(sc.ResultFeeColor));
    }

    maxBlobSize(): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxBlobSize));
    }

    maxEventSize(): wasmlib.ScMutableInt16 {
        return new wasmlib.ScMutableInt16(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxEventSize));
    }

    maxEventsPerReq(): wasmlib.ScMutableInt16 {
        return new wasmlib.ScMutableInt16(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxEventsPerReq));
    }
}

export class ImmutableGetFeeInfoResults extends wasmlib.ScMapID {

    feeColor(): wasmlib.ScImmutableColor {
        return new wasmlib.ScImmutableColor(this.mapID, wasmlib.Key32.fromString(sc.ResultFeeColor));
    }

    ownerFee(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultOwnerFee));
    }

    validatorFee(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultValidatorFee));
    }
}

export class MutableGetFeeInfoResults extends wasmlib.ScMapID {

    feeColor(): wasmlib.ScMutableColor {
        return new wasmlib.ScMutableColor(this.mapID, wasmlib.Key32.fromString(sc.ResultFeeColor));
    }

    ownerFee(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultOwnerFee));
    }

    validatorFee(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultValidatorFee));
    }
}

export class ImmutableGetMaxBlobSizeResults extends wasmlib.ScMapID {

    maxBlobSize(): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxBlobSize));
    }
}

export class MutableGetMaxBlobSizeResults extends wasmlib.ScMapID {

    maxBlobSize(): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.mapID, wasmlib.Key32.fromString(sc.ResultMaxBlobSize));
    }
}
