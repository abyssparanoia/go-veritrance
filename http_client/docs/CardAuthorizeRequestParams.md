# CardAuthorizeRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**PayNowIdParam** | [**CardAuthorizeRequestParamsAllOfPayNowIdParam**](CardAuthorizeRequestParamsAllOfPayNowIdParam.md) |  | 
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

### NewCardAuthorizeRequestParams

`func NewCardAuthorizeRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardAuthorizeRequestParamsAllOfPayNowIdParam, orderId string, amount string, currencyUnit CurrencyUnit, ) *CardAuthorizeRequestParams`

NewCardAuthorizeRequestParams instantiates a new CardAuthorizeRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeRequestParamsWithDefaults

`func NewCardAuthorizeRequestParamsWithDefaults() *CardAuthorizeRequestParams`

NewCardAuthorizeRequestParamsWithDefaults instantiates a new CardAuthorizeRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardAuthorizeRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardAuthorizeRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardAuthorizeRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardAuthorizeRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardAuthorizeRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardAuthorizeRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardAuthorizeRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardAuthorizeRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardAuthorizeRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetPayNowIdParam

`func (o *CardAuthorizeRequestParams) GetPayNowIdParam() CardAuthorizeRequestParamsAllOfPayNowIdParam`

GetPayNowIdParam returns the PayNowIdParam field if non-nil, zero value otherwise.

### GetPayNowIdParamOk

`func (o *CardAuthorizeRequestParams) GetPayNowIdParamOk() (*CardAuthorizeRequestParamsAllOfPayNowIdParam, bool)`

GetPayNowIdParamOk returns a tuple with the PayNowIdParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdParam

`func (o *CardAuthorizeRequestParams) SetPayNowIdParam(v CardAuthorizeRequestParamsAllOfPayNowIdParam)`

SetPayNowIdParam sets PayNowIdParam field to given value.


### GetOrderId

`func (o *CardAuthorizeRequestParams) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardAuthorizeRequestParams) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardAuthorizeRequestParams) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardAuthorizeRequestParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardAuthorizeRequestParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardAuthorizeRequestParams) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyUnit

`func (o *CardAuthorizeRequestParams) GetCurrencyUnit() CurrencyUnit`

GetCurrencyUnit returns the CurrencyUnit field if non-nil, zero value otherwise.

### GetCurrencyUnitOk

`func (o *CardAuthorizeRequestParams) GetCurrencyUnitOk() (*CurrencyUnit, bool)`

GetCurrencyUnitOk returns a tuple with the CurrencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyUnit

`func (o *CardAuthorizeRequestParams) SetCurrencyUnit(v CurrencyUnit)`

SetCurrencyUnit sets CurrencyUnit field to given value.


### GetToken

`func (o *CardAuthorizeRequestParams) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardAuthorizeRequestParams) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardAuthorizeRequestParams) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardAuthorizeRequestParams) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardAuthorizeRequestParams) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardAuthorizeRequestParams) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardAuthorizeRequestParams) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardAuthorizeRequestParams) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardAuthorizeRequestParams) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardAuthorizeRequestParams) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardAuthorizeRequestParams) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardAuthorizeRequestParams) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardAuthorizeRequestParams) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardAuthorizeRequestParams) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardAuthorizeRequestParams) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardAuthorizeRequestParams) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardOptionType

`func (o *CardAuthorizeRequestParams) GetCardOptionType() string`

GetCardOptionType returns the CardOptionType field if non-nil, zero value otherwise.

### GetCardOptionTypeOk

`func (o *CardAuthorizeRequestParams) GetCardOptionTypeOk() (*string, bool)`

GetCardOptionTypeOk returns a tuple with the CardOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptionType

`func (o *CardAuthorizeRequestParams) SetCardOptionType(v string)`

SetCardOptionType sets CardOptionType field to given value.

### HasCardOptionType

`func (o *CardAuthorizeRequestParams) HasCardOptionType() bool`

HasCardOptionType returns a boolean if a field has been set.

### GetJpo

`func (o *CardAuthorizeRequestParams) GetJpo() string`

GetJpo returns the Jpo field if non-nil, zero value otherwise.

### GetJpoOk

`func (o *CardAuthorizeRequestParams) GetJpoOk() (*string, bool)`

GetJpoOk returns a tuple with the Jpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpo

`func (o *CardAuthorizeRequestParams) SetJpo(v string)`

SetJpo sets Jpo field to given value.

### HasJpo

`func (o *CardAuthorizeRequestParams) HasJpo() bool`

HasJpo returns a boolean if a field has been set.

### GetWithCapture

`func (o *CardAuthorizeRequestParams) GetWithCapture() bool`

GetWithCapture returns the WithCapture field if non-nil, zero value otherwise.

### GetWithCaptureOk

`func (o *CardAuthorizeRequestParams) GetWithCaptureOk() (*bool, bool)`

GetWithCaptureOk returns a tuple with the WithCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCapture

`func (o *CardAuthorizeRequestParams) SetWithCapture(v bool)`

SetWithCapture sets WithCapture field to given value.

### HasWithCapture

`func (o *CardAuthorizeRequestParams) HasWithCapture() bool`

HasWithCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


