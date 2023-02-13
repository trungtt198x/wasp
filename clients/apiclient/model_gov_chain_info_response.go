/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the GovChainInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovChainInfoResponse{}

// GovChainInfoResponse struct for GovChainInfoResponse
type GovChainInfoResponse struct {
	// ChainID (Bech32-encoded).
	ChainID string `json:"chainID"`
	// The chain owner address (Bech32-encoded).
	ChainOwnerId string `json:"chainOwnerId"`
	// The description of the chain.
	Description string `json:"description"`
	GasFeePolicy GasFeePolicy `json:"gasFeePolicy"`
	// The maximum contract blob size.
	MaxBlobSize int32 `json:"maxBlobSize"`
	// The maximum event size.
	MaxEventSize int32 `json:"maxEventSize"`
	// The maximum amount of events per request.
	MaxEventsPerReq int32 `json:"maxEventsPerReq"`
}

// NewGovChainInfoResponse instantiates a new GovChainInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovChainInfoResponse(chainID string, chainOwnerId string, description string, gasFeePolicy GasFeePolicy, maxBlobSize int32, maxEventSize int32, maxEventsPerReq int32) *GovChainInfoResponse {
	this := GovChainInfoResponse{}
	this.ChainID = chainID
	this.ChainOwnerId = chainOwnerId
	this.Description = description
	this.GasFeePolicy = gasFeePolicy
	this.MaxBlobSize = maxBlobSize
	this.MaxEventSize = maxEventSize
	this.MaxEventsPerReq = maxEventsPerReq
	return &this
}

// NewGovChainInfoResponseWithDefaults instantiates a new GovChainInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovChainInfoResponseWithDefaults() *GovChainInfoResponse {
	this := GovChainInfoResponse{}
	return &this
}

// GetChainID returns the ChainID field value
func (o *GovChainInfoResponse) GetChainID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainID
}

// GetChainIDOk returns a tuple with the ChainID field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetChainIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainID, true
}

// SetChainID sets field value
func (o *GovChainInfoResponse) SetChainID(v string) {
	o.ChainID = v
}

// GetChainOwnerId returns the ChainOwnerId field value
func (o *GovChainInfoResponse) GetChainOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainOwnerId
}

// GetChainOwnerIdOk returns a tuple with the ChainOwnerId field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetChainOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainOwnerId, true
}

// SetChainOwnerId sets field value
func (o *GovChainInfoResponse) SetChainOwnerId(v string) {
	o.ChainOwnerId = v
}

// GetDescription returns the Description field value
func (o *GovChainInfoResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GovChainInfoResponse) SetDescription(v string) {
	o.Description = v
}

// GetGasFeePolicy returns the GasFeePolicy field value
func (o *GovChainInfoResponse) GetGasFeePolicy() GasFeePolicy {
	if o == nil {
		var ret GasFeePolicy
		return ret
	}

	return o.GasFeePolicy
}

// GetGasFeePolicyOk returns a tuple with the GasFeePolicy field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetGasFeePolicyOk() (*GasFeePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasFeePolicy, true
}

// SetGasFeePolicy sets field value
func (o *GovChainInfoResponse) SetGasFeePolicy(v GasFeePolicy) {
	o.GasFeePolicy = v
}

// GetMaxBlobSize returns the MaxBlobSize field value
func (o *GovChainInfoResponse) GetMaxBlobSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxBlobSize
}

// GetMaxBlobSizeOk returns a tuple with the MaxBlobSize field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetMaxBlobSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBlobSize, true
}

// SetMaxBlobSize sets field value
func (o *GovChainInfoResponse) SetMaxBlobSize(v int32) {
	o.MaxBlobSize = v
}

// GetMaxEventSize returns the MaxEventSize field value
func (o *GovChainInfoResponse) GetMaxEventSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxEventSize
}

// GetMaxEventSizeOk returns a tuple with the MaxEventSize field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetMaxEventSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxEventSize, true
}

// SetMaxEventSize sets field value
func (o *GovChainInfoResponse) SetMaxEventSize(v int32) {
	o.MaxEventSize = v
}

// GetMaxEventsPerReq returns the MaxEventsPerReq field value
func (o *GovChainInfoResponse) GetMaxEventsPerReq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxEventsPerReq
}

// GetMaxEventsPerReqOk returns a tuple with the MaxEventsPerReq field value
// and a boolean to check if the value has been set.
func (o *GovChainInfoResponse) GetMaxEventsPerReqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxEventsPerReq, true
}

// SetMaxEventsPerReq sets field value
func (o *GovChainInfoResponse) SetMaxEventsPerReq(v int32) {
	o.MaxEventsPerReq = v
}

func (o GovChainInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovChainInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainID"] = o.ChainID
	toSerialize["chainOwnerId"] = o.ChainOwnerId
	toSerialize["description"] = o.Description
	toSerialize["gasFeePolicy"] = o.GasFeePolicy
	toSerialize["maxBlobSize"] = o.MaxBlobSize
	toSerialize["maxEventSize"] = o.MaxEventSize
	toSerialize["maxEventsPerReq"] = o.MaxEventsPerReq
	return toSerialize, nil
}

type NullableGovChainInfoResponse struct {
	value *GovChainInfoResponse
	isSet bool
}

func (v NullableGovChainInfoResponse) Get() *GovChainInfoResponse {
	return v.value
}

func (v *NullableGovChainInfoResponse) Set(val *GovChainInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGovChainInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGovChainInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovChainInfoResponse(val *GovChainInfoResponse) *NullableGovChainInfoResponse {
	return &NullableGovChainInfoResponse{value: val, isSet: true}
}

func (v NullableGovChainInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovChainInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

