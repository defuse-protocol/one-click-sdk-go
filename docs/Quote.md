# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositAddress** | Pointer to **string** | The deposit address on the chain of &#x60;originAsset&#x60; in case if &#x60;depositType&#x60; is &#x60;ORIGIN_CHAIN&#x60;.  The deposit address inside of near intents (the verifier smart contract) in case if &#x60;depositType&#x60; is &#x60;INTENTS&#x60;. | [optional] 
**AmountIn** | **string** | Amount of the origin asset | 
**AmountInFormatted** | **string** | Amount of the origin asset in readable format | 
**AmountInUsd** | **string** | Amount of the origin assets equivalent in USD | 
**MinAmountIn** | **string** | Minimum amount of the origin asset that will be used for swap | 
**AmountOut** | **string** | Amount of the destination asset | 
**AmountOutFormatted** | **string** | Amount of the destination asset in readable format | 
**AmountOutUsd** | **string** | Amount of the destination asset equivalent in USD | 
**MinAmountOut** | **string** | Minimum amount with slippage taken into account | 
**Deadline** | Pointer to **time.Time** | Time when the deposit address will become inactive and funds might be lost | [optional] 
**TimeWhenInactive** | Pointer to **time.Time** | Time when the deposit address will become cold and swap processing will take more time | [optional] 
**TimeEstimate** | Pointer to **float32** | Estimated time in seconds for swap to be executed after the deposit transaction is confirmed | [optional] 

## Methods

### NewQuote

`func NewQuote(amountIn string, amountInFormatted string, amountInUsd string, minAmountIn string, amountOut string, amountOutFormatted string, amountOutUsd string, minAmountOut string, ) *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositAddress

`func (o *Quote) GetDepositAddress() string`

GetDepositAddress returns the DepositAddress field if non-nil, zero value otherwise.

### GetDepositAddressOk

`func (o *Quote) GetDepositAddressOk() (*string, bool)`

GetDepositAddressOk returns a tuple with the DepositAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddress

`func (o *Quote) SetDepositAddress(v string)`

SetDepositAddress sets DepositAddress field to given value.

### HasDepositAddress

`func (o *Quote) HasDepositAddress() bool`

HasDepositAddress returns a boolean if a field has been set.

### GetAmountIn

`func (o *Quote) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *Quote) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *Quote) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.


### GetAmountInFormatted

`func (o *Quote) GetAmountInFormatted() string`

GetAmountInFormatted returns the AmountInFormatted field if non-nil, zero value otherwise.

### GetAmountInFormattedOk

`func (o *Quote) GetAmountInFormattedOk() (*string, bool)`

GetAmountInFormattedOk returns a tuple with the AmountInFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInFormatted

`func (o *Quote) SetAmountInFormatted(v string)`

SetAmountInFormatted sets AmountInFormatted field to given value.


### GetAmountInUsd

`func (o *Quote) GetAmountInUsd() string`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *Quote) GetAmountInUsdOk() (*string, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *Quote) SetAmountInUsd(v string)`

SetAmountInUsd sets AmountInUsd field to given value.


### GetMinAmountIn

`func (o *Quote) GetMinAmountIn() string`

GetMinAmountIn returns the MinAmountIn field if non-nil, zero value otherwise.

### GetMinAmountInOk

`func (o *Quote) GetMinAmountInOk() (*string, bool)`

GetMinAmountInOk returns a tuple with the MinAmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmountIn

`func (o *Quote) SetMinAmountIn(v string)`

SetMinAmountIn sets MinAmountIn field to given value.


### GetAmountOut

`func (o *Quote) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *Quote) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *Quote) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.


### GetAmountOutFormatted

`func (o *Quote) GetAmountOutFormatted() string`

GetAmountOutFormatted returns the AmountOutFormatted field if non-nil, zero value otherwise.

### GetAmountOutFormattedOk

`func (o *Quote) GetAmountOutFormattedOk() (*string, bool)`

GetAmountOutFormattedOk returns a tuple with the AmountOutFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutFormatted

`func (o *Quote) SetAmountOutFormatted(v string)`

SetAmountOutFormatted sets AmountOutFormatted field to given value.


### GetAmountOutUsd

`func (o *Quote) GetAmountOutUsd() string`

GetAmountOutUsd returns the AmountOutUsd field if non-nil, zero value otherwise.

### GetAmountOutUsdOk

`func (o *Quote) GetAmountOutUsdOk() (*string, bool)`

GetAmountOutUsdOk returns a tuple with the AmountOutUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutUsd

`func (o *Quote) SetAmountOutUsd(v string)`

SetAmountOutUsd sets AmountOutUsd field to given value.


### GetMinAmountOut

`func (o *Quote) GetMinAmountOut() string`

GetMinAmountOut returns the MinAmountOut field if non-nil, zero value otherwise.

### GetMinAmountOutOk

`func (o *Quote) GetMinAmountOutOk() (*string, bool)`

GetMinAmountOutOk returns a tuple with the MinAmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmountOut

`func (o *Quote) SetMinAmountOut(v string)`

SetMinAmountOut sets MinAmountOut field to given value.


### GetDeadline

`func (o *Quote) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *Quote) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *Quote) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *Quote) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetTimeWhenInactive

`func (o *Quote) GetTimeWhenInactive() time.Time`

GetTimeWhenInactive returns the TimeWhenInactive field if non-nil, zero value otherwise.

### GetTimeWhenInactiveOk

`func (o *Quote) GetTimeWhenInactiveOk() (*time.Time, bool)`

GetTimeWhenInactiveOk returns a tuple with the TimeWhenInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWhenInactive

`func (o *Quote) SetTimeWhenInactive(v time.Time)`

SetTimeWhenInactive sets TimeWhenInactive field to given value.

### HasTimeWhenInactive

`func (o *Quote) HasTimeWhenInactive() bool`

HasTimeWhenInactive returns a boolean if a field has been set.

### GetTimeEstimate

`func (o *Quote) GetTimeEstimate() float32`

GetTimeEstimate returns the TimeEstimate field if non-nil, zero value otherwise.

### GetTimeEstimateOk

`func (o *Quote) GetTimeEstimateOk() (*float32, bool)`

GetTimeEstimateOk returns a tuple with the TimeEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEstimate

`func (o *Quote) SetTimeEstimate(v float32)`

SetTimeEstimate sets TimeEstimate field to given value.

### HasTimeEstimate

`func (o *Quote) HasTimeEstimate() bool`

HasTimeEstimate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


