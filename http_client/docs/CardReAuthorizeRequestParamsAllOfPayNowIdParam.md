# CardReAuthorizeRequestParamsAllOfPayNowIdParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountParam** | **map[string]interface{}** |  | 
**AccountId** | Pointer to **string** |  | [optional] 
**OrderId** | **string** |  | 
**OriginalOrderId** | **string** |  | 
**Amount** | **string** |  | 
**Token** | Pointer to **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | [optional] 
**CardNumber** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardExpire** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**SecurityCode** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardOptionType** | Pointer to **string** | カードオプションタイプ （MPI 有り/無し） | [optional] 
**Jpo** | Pointer to **string** | 支払種別 \&quot;10\&quot;： 一括払い \&quot;21\&quot;： ボーナス一括 \&quot;61Cxx\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\&quot;10\&quot;（一括払い）が適用されます。  | [optional] 
**WithCapture** | Pointer to **bool** | 売上フラグ \&quot;true\&quot;： 与信・売上 \&quot;false\&quot;： 与信のみ  | [optional] 

## Methods

### NewCardReAuthorizeRequestParamsAllOfPayNowIdParam

`func NewCardReAuthorizeRequestParamsAllOfPayNowIdParam(accountParam map[string]interface{}, orderId string, originalOrderId string, amount string, ) *CardReAuthorizeRequestParamsAllOfPayNowIdParam`

NewCardReAuthorizeRequestParamsAllOfPayNowIdParam instantiates a new CardReAuthorizeRequestParamsAllOfPayNowIdParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardReAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults

`func NewCardReAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults() *CardReAuthorizeRequestParamsAllOfPayNowIdParam`

NewCardReAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults instantiates a new CardReAuthorizeRequestParamsAllOfPayNowIdParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountParam

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountParam() map[string]interface{}`

GetAccountParam returns the AccountParam field if non-nil, zero value otherwise.

### GetAccountParamOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountParamOk() (*map[string]interface{}, bool)`

GetAccountParamOk returns a tuple with the AccountParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountParam

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetAccountParam(v map[string]interface{})`

SetAccountParam sets AccountParam field to given value.


### GetAccountId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetOrderId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetOriginalOrderId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetOriginalOrderId() string`

GetOriginalOrderId returns the OriginalOrderId field if non-nil, zero value otherwise.

### GetOriginalOrderIdOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetOriginalOrderIdOk() (*string, bool)`

GetOriginalOrderIdOk returns a tuple with the OriginalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderId

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetOriginalOrderId(v string)`

SetOriginalOrderId sets OriginalOrderId field to given value.


### GetAmount

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToken

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardOptionType

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardOptionType() string`

GetCardOptionType returns the CardOptionType field if non-nil, zero value otherwise.

### GetCardOptionTypeOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetCardOptionTypeOk() (*string, bool)`

GetCardOptionTypeOk returns a tuple with the CardOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptionType

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetCardOptionType(v string)`

SetCardOptionType sets CardOptionType field to given value.

### HasCardOptionType

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasCardOptionType() bool`

HasCardOptionType returns a boolean if a field has been set.

### GetJpo

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetJpo() string`

GetJpo returns the Jpo field if non-nil, zero value otherwise.

### GetJpoOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetJpoOk() (*string, bool)`

GetJpoOk returns a tuple with the Jpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpo

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetJpo(v string)`

SetJpo sets Jpo field to given value.

### HasJpo

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasJpo() bool`

HasJpo returns a boolean if a field has been set.

### GetWithCapture

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetWithCapture() bool`

GetWithCapture returns the WithCapture field if non-nil, zero value otherwise.

### GetWithCaptureOk

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) GetWithCaptureOk() (*bool, bool)`

GetWithCaptureOk returns a tuple with the WithCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCapture

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) SetWithCapture(v bool)`

SetWithCapture sets WithCapture field to given value.

### HasWithCapture

`func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParam) HasWithCapture() bool`

HasWithCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


