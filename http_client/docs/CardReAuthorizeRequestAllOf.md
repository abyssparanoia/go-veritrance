# CardReAuthorizeRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | [optional] 
**CardNumber** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardExpire** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**SecurityCode** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardOptionType** | Pointer to **string** | カードオプションタイプ （MPI 有り/無し） | [optional] 
**Jpo** | Pointer to **string** | 支払種別 \&quot;10\&quot;： 一括払い \&quot;21\&quot;： ボーナス一括 \&quot;61Cxx\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\&quot;10\&quot;（一括払い）が適用されます。  | [optional] 
**WithCapture** | Pointer to **bool** | 売上フラグ \&quot;true\&quot;： 与信・売上 \&quot;false\&quot;： 与信のみ  | [optional] 

## Methods

### NewCardReAuthorizeRequestAllOf

`func NewCardReAuthorizeRequestAllOf() *CardReAuthorizeRequestAllOf`

NewCardReAuthorizeRequestAllOf instantiates a new CardReAuthorizeRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardReAuthorizeRequestAllOfWithDefaults

`func NewCardReAuthorizeRequestAllOfWithDefaults() *CardReAuthorizeRequestAllOf`

NewCardReAuthorizeRequestAllOfWithDefaults instantiates a new CardReAuthorizeRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CardReAuthorizeRequestAllOf) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardReAuthorizeRequestAllOf) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardReAuthorizeRequestAllOf) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CardReAuthorizeRequestAllOf) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAmount

`func (o *CardReAuthorizeRequestAllOf) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardReAuthorizeRequestAllOf) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardReAuthorizeRequestAllOf) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CardReAuthorizeRequestAllOf) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetToken

`func (o *CardReAuthorizeRequestAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardReAuthorizeRequestAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardReAuthorizeRequestAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardReAuthorizeRequestAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardReAuthorizeRequestAllOf) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardReAuthorizeRequestAllOf) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardReAuthorizeRequestAllOf) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardReAuthorizeRequestAllOf) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardReAuthorizeRequestAllOf) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardReAuthorizeRequestAllOf) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardReAuthorizeRequestAllOf) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardReAuthorizeRequestAllOf) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardReAuthorizeRequestAllOf) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardReAuthorizeRequestAllOf) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardReAuthorizeRequestAllOf) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardReAuthorizeRequestAllOf) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardOptionType

`func (o *CardReAuthorizeRequestAllOf) GetCardOptionType() string`

GetCardOptionType returns the CardOptionType field if non-nil, zero value otherwise.

### GetCardOptionTypeOk

`func (o *CardReAuthorizeRequestAllOf) GetCardOptionTypeOk() (*string, bool)`

GetCardOptionTypeOk returns a tuple with the CardOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptionType

`func (o *CardReAuthorizeRequestAllOf) SetCardOptionType(v string)`

SetCardOptionType sets CardOptionType field to given value.

### HasCardOptionType

`func (o *CardReAuthorizeRequestAllOf) HasCardOptionType() bool`

HasCardOptionType returns a boolean if a field has been set.

### GetJpo

`func (o *CardReAuthorizeRequestAllOf) GetJpo() string`

GetJpo returns the Jpo field if non-nil, zero value otherwise.

### GetJpoOk

`func (o *CardReAuthorizeRequestAllOf) GetJpoOk() (*string, bool)`

GetJpoOk returns a tuple with the Jpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpo

`func (o *CardReAuthorizeRequestAllOf) SetJpo(v string)`

SetJpo sets Jpo field to given value.

### HasJpo

`func (o *CardReAuthorizeRequestAllOf) HasJpo() bool`

HasJpo returns a boolean if a field has been set.

### GetWithCapture

`func (o *CardReAuthorizeRequestAllOf) GetWithCapture() bool`

GetWithCapture returns the WithCapture field if non-nil, zero value otherwise.

### GetWithCaptureOk

`func (o *CardReAuthorizeRequestAllOf) GetWithCaptureOk() (*bool, bool)`

GetWithCaptureOk returns a tuple with the WithCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCapture

`func (o *CardReAuthorizeRequestAllOf) SetWithCapture(v bool)`

SetWithCapture sets WithCapture field to given value.

### HasWithCapture

`func (o *CardReAuthorizeRequestAllOf) HasWithCapture() bool`

HasWithCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


