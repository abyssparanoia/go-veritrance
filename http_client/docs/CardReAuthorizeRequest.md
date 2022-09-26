# CardReAuthorizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**OrderId** | **string** |  | 
**Amount** | **string** |  | 
**Token** | Pointer to **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | [optional] 
**CardNumber** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardExpire** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**SecurityCode** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardOptionType** | Pointer to **string** | カードオプションタイプ （MPI 有り/無し） | [optional] 
**Jpo** | Pointer to **string** | 支払種別 \&quot;10\&quot;： 一括払い \&quot;21\&quot;： ボーナス一括 \&quot;61Cxx\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\&quot;10\&quot;（一括払い）が適用されます。  | [optional] 
**WithCapture** | Pointer to **bool** | 売上フラグ \&quot;true\&quot;： 与信・売上 \&quot;false\&quot;： 与信のみ  | [optional] 

## Methods

### NewCardReAuthorizeRequest

`func NewCardReAuthorizeRequest(txnVersion string, dummyRequest string, merchantCcid string, orderId string, amount string, ) *CardReAuthorizeRequest`

NewCardReAuthorizeRequest instantiates a new CardReAuthorizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardReAuthorizeRequestWithDefaults

`func NewCardReAuthorizeRequestWithDefaults() *CardReAuthorizeRequest`

NewCardReAuthorizeRequestWithDefaults instantiates a new CardReAuthorizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardReAuthorizeRequest) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardReAuthorizeRequest) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardReAuthorizeRequest) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardReAuthorizeRequest) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardReAuthorizeRequest) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardReAuthorizeRequest) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardReAuthorizeRequest) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardReAuthorizeRequest) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardReAuthorizeRequest) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetOrderId

`func (o *CardReAuthorizeRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardReAuthorizeRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardReAuthorizeRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardReAuthorizeRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardReAuthorizeRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardReAuthorizeRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToken

`func (o *CardReAuthorizeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardReAuthorizeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardReAuthorizeRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardReAuthorizeRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardReAuthorizeRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardReAuthorizeRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardReAuthorizeRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardReAuthorizeRequest) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardReAuthorizeRequest) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardReAuthorizeRequest) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardReAuthorizeRequest) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardReAuthorizeRequest) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardReAuthorizeRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardReAuthorizeRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardReAuthorizeRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardReAuthorizeRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardOptionType

`func (o *CardReAuthorizeRequest) GetCardOptionType() string`

GetCardOptionType returns the CardOptionType field if non-nil, zero value otherwise.

### GetCardOptionTypeOk

`func (o *CardReAuthorizeRequest) GetCardOptionTypeOk() (*string, bool)`

GetCardOptionTypeOk returns a tuple with the CardOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptionType

`func (o *CardReAuthorizeRequest) SetCardOptionType(v string)`

SetCardOptionType sets CardOptionType field to given value.

### HasCardOptionType

`func (o *CardReAuthorizeRequest) HasCardOptionType() bool`

HasCardOptionType returns a boolean if a field has been set.

### GetJpo

`func (o *CardReAuthorizeRequest) GetJpo() string`

GetJpo returns the Jpo field if non-nil, zero value otherwise.

### GetJpoOk

`func (o *CardReAuthorizeRequest) GetJpoOk() (*string, bool)`

GetJpoOk returns a tuple with the Jpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpo

`func (o *CardReAuthorizeRequest) SetJpo(v string)`

SetJpo sets Jpo field to given value.

### HasJpo

`func (o *CardReAuthorizeRequest) HasJpo() bool`

HasJpo returns a boolean if a field has been set.

### GetWithCapture

`func (o *CardReAuthorizeRequest) GetWithCapture() bool`

GetWithCapture returns the WithCapture field if non-nil, zero value otherwise.

### GetWithCaptureOk

`func (o *CardReAuthorizeRequest) GetWithCaptureOk() (*bool, bool)`

GetWithCaptureOk returns a tuple with the WithCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCapture

`func (o *CardReAuthorizeRequest) SetWithCapture(v bool)`

SetWithCapture sets WithCapture field to given value.

### HasWithCapture

`func (o *CardReAuthorizeRequest) HasWithCapture() bool`

HasWithCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


