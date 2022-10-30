# CardCancelRequestParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**Amount** | **string** |  | 

## Methods

### NewCardCancelRequestParam

`func NewCardCancelRequestParam(orderId string, amount string, ) *CardCancelRequestParam`

NewCardCancelRequestParam instantiates a new CardCancelRequestParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCancelRequestParamWithDefaults

`func NewCardCancelRequestParamWithDefaults() *CardCancelRequestParam`

NewCardCancelRequestParamWithDefaults instantiates a new CardCancelRequestParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CardCancelRequestParam) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCancelRequestParam) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCancelRequestParam) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardCancelRequestParam) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardCancelRequestParam) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardCancelRequestParam) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


