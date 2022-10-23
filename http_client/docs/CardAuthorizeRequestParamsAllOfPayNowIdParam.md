# CardAuthorizeRequestParamsAllOfPayNowIdParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountParam** | [**AccountAddRequestParamsAllOfPayNowIdParamAccountParam**](AccountAddRequestParamsAllOfPayNowIdParamAccountParam.md) |  | 
**OrderId** | **string** |  | 
**Amount** | **string** |  | 
**CurrencyUnit** | [**CurrencyUnit**](CurrencyUnit.md) |  | 
**Token** | Pointer to **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | [optional] 
**CardNumber** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardExpire** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**SecurityCode** | Pointer to **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | [optional] 
**CardOptionType** | Pointer to **string** | カードオプションタイプ （MPI 有り/無し） | [optional] 
**Jpo** | Pointer to **string** | 支払種別 \&quot;10\&quot;： 一括払い \&quot;21\&quot;： ボーナス一括 \&quot;61Cxx\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\&quot;10\&quot;（一括払い）が適用されます。  | [optional] 
**WithCapture** | Pointer to **bool** | 売上フラグ \&quot;true\&quot;： 与信・売上 \&quot;false\&quot;： 与信のみ  | [optional] 

## Methods

### NewCardAuthorizeRequestParamsAllOfPayNowIdParam

`func NewCardAuthorizeRequestParamsAllOfPayNowIdParam(accountParam AccountAddRequestParamsAllOfPayNowIdParamAccountParam, orderId string, amount string, currencyUnit CurrencyUnit, ) *CardAuthorizeRequestParamsAllOfPayNowIdParam`

NewCardAuthorizeRequestParamsAllOfPayNowIdParam instantiates a new CardAuthorizeRequestParamsAllOfPayNowIdParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults

`func NewCardAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults() *CardAuthorizeRequestParamsAllOfPayNowIdParam`

NewCardAuthorizeRequestParamsAllOfPayNowIdParamWithDefaults instantiates a new CardAuthorizeRequestParamsAllOfPayNowIdParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountParam

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountParam() AccountAddRequestParamsAllOfPayNowIdParamAccountParam`

GetAccountParam returns the AccountParam field if non-nil, zero value otherwise.

### GetAccountParamOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetAccountParamOk() (*AccountAddRequestParamsAllOfPayNowIdParamAccountParam, bool)`

GetAccountParamOk returns a tuple with the AccountParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountParam

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetAccountParam(v AccountAddRequestParamsAllOfPayNowIdParamAccountParam)`

SetAccountParam sets AccountParam field to given value.


### GetOrderId

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyUnit

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCurrencyUnit() CurrencyUnit`

GetCurrencyUnit returns the CurrencyUnit field if non-nil, zero value otherwise.

### GetCurrencyUnitOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCurrencyUnitOk() (*CurrencyUnit, bool)`

GetCurrencyUnitOk returns a tuple with the CurrencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyUnit

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetCurrencyUnit(v CurrencyUnit)`

SetCurrencyUnit sets CurrencyUnit field to given value.


### GetToken

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardOptionType

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardOptionType() string`

GetCardOptionType returns the CardOptionType field if non-nil, zero value otherwise.

### GetCardOptionTypeOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetCardOptionTypeOk() (*string, bool)`

GetCardOptionTypeOk returns a tuple with the CardOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptionType

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetCardOptionType(v string)`

SetCardOptionType sets CardOptionType field to given value.

### HasCardOptionType

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasCardOptionType() bool`

HasCardOptionType returns a boolean if a field has been set.

### GetJpo

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetJpo() string`

GetJpo returns the Jpo field if non-nil, zero value otherwise.

### GetJpoOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetJpoOk() (*string, bool)`

GetJpoOk returns a tuple with the Jpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpo

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetJpo(v string)`

SetJpo sets Jpo field to given value.

### HasJpo

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasJpo() bool`

HasJpo returns a boolean if a field has been set.

### GetWithCapture

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetWithCapture() bool`

GetWithCapture returns the WithCapture field if non-nil, zero value otherwise.

### GetWithCaptureOk

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) GetWithCaptureOk() (*bool, bool)`

GetWithCaptureOk returns a tuple with the WithCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCapture

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) SetWithCapture(v bool)`

SetWithCapture sets WithCapture field to given value.

### HasWithCapture

`func (o *CardAuthorizeRequestParamsAllOfPayNowIdParam) HasWithCapture() bool`

HasWithCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


