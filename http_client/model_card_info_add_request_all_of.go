/*
Veritrance API

Veritrance API Requestの形式はcomponentsを参照。文字列にしたものを各Requestのparamsに用いる 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package veritrance_paynow

import (
	"encoding/json"
)

// CardInfoAddRequestAllOf struct for CardInfoAddRequestAllOf
type CardInfoAddRequestAllOf struct {
	AccountId *string `json:"accountId,omitempty"`
	Token *string `json:"token,omitempty"`
	CardNumber *string `json:"cardNumber,omitempty"`
	CardExpire *string `json:"cardExpire,omitempty"`
	SecurityCode *string `json:"securityCode,omitempty"`
}

// NewCardInfoAddRequestAllOf instantiates a new CardInfoAddRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoAddRequestAllOf() *CardInfoAddRequestAllOf {
	this := CardInfoAddRequestAllOf{}
	return &this
}

// NewCardInfoAddRequestAllOfWithDefaults instantiates a new CardInfoAddRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoAddRequestAllOfWithDefaults() *CardInfoAddRequestAllOf {
	this := CardInfoAddRequestAllOf{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CardInfoAddRequestAllOf) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestAllOf) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CardInfoAddRequestAllOf) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CardInfoAddRequestAllOf) SetAccountId(v string) {
	o.AccountId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardInfoAddRequestAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardInfoAddRequestAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardInfoAddRequestAllOf) SetToken(v string) {
	o.Token = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CardInfoAddRequestAllOf) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestAllOf) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CardInfoAddRequestAllOf) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CardInfoAddRequestAllOf) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardExpire returns the CardExpire field value if set, zero value otherwise.
func (o *CardInfoAddRequestAllOf) GetCardExpire() string {
	if o == nil || o.CardExpire == nil {
		var ret string
		return ret
	}
	return *o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestAllOf) GetCardExpireOk() (*string, bool) {
	if o == nil || o.CardExpire == nil {
		return nil, false
	}
	return o.CardExpire, true
}

// HasCardExpire returns a boolean if a field has been set.
func (o *CardInfoAddRequestAllOf) HasCardExpire() bool {
	if o != nil && o.CardExpire != nil {
		return true
	}

	return false
}

// SetCardExpire gets a reference to the given string and assigns it to the CardExpire field.
func (o *CardInfoAddRequestAllOf) SetCardExpire(v string) {
	o.CardExpire = &v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *CardInfoAddRequestAllOf) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestAllOf) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *CardInfoAddRequestAllOf) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *CardInfoAddRequestAllOf) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

func (o CardInfoAddRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.CardNumber != nil {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if o.CardExpire != nil {
		toSerialize["cardExpire"] = o.CardExpire
	}
	if o.SecurityCode != nil {
		toSerialize["securityCode"] = o.SecurityCode
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoAddRequestAllOf struct {
	value *CardInfoAddRequestAllOf
	isSet bool
}

func (v NullableCardInfoAddRequestAllOf) Get() *CardInfoAddRequestAllOf {
	return v.value
}

func (v *NullableCardInfoAddRequestAllOf) Set(val *CardInfoAddRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoAddRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoAddRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoAddRequestAllOf(val *CardInfoAddRequestAllOf) *NullableCardInfoAddRequestAllOf {
	return &NullableCardInfoAddRequestAllOf{value: val, isSet: true}
}

func (v NullableCardInfoAddRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoAddRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


