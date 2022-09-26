# CardInfoV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | 決済サーバーにて付与されたカード ID が返戻されます。 | 
**CardNumber** | **string** | カード番号の先頭 6 桁と下 2 桁のみ数字表記され、その他は \&quot;*\&quot;（アスタリスク）に変換されます。 | 
**CardExpire** | **string** |  | 
**DefaultCard** | **int32** | カード情報を明示的に指定せず決済する場合に使用するカードか否かを示すフラグ \&quot;1\&quot;：決済時、カード情報を明示的に指定しない場合に使用されるカード \&quot;0\&quot;：決済時、カード情報を明示的に指定しない場合に使用されるカードではない  | 
**CardholderName** | Pointer to **string** | カードの名義人 | [optional] 

## Methods

### NewCardInfoV1

`func NewCardInfoV1(cardId string, cardNumber string, cardExpire string, defaultCard int32, ) *CardInfoV1`

NewCardInfoV1 instantiates a new CardInfoV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoV1WithDefaults

`func NewCardInfoV1WithDefaults() *CardInfoV1`

NewCardInfoV1WithDefaults instantiates a new CardInfoV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *CardInfoV1) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CardInfoV1) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CardInfoV1) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCardNumber

`func (o *CardInfoV1) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardInfoV1) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardInfoV1) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardExpire

`func (o *CardInfoV1) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardInfoV1) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardInfoV1) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.


### GetDefaultCard

`func (o *CardInfoV1) GetDefaultCard() int32`

GetDefaultCard returns the DefaultCard field if non-nil, zero value otherwise.

### GetDefaultCardOk

`func (o *CardInfoV1) GetDefaultCardOk() (*int32, bool)`

GetDefaultCardOk returns a tuple with the DefaultCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCard

`func (o *CardInfoV1) SetDefaultCard(v int32)`

SetDefaultCard sets DefaultCard field to given value.


### GetCardholderName

`func (o *CardInfoV1) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *CardInfoV1) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *CardInfoV1) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *CardInfoV1) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


