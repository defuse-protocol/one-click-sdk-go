# SwapDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntentHashes** | **[]string** | All intent hashes that took part in this swap | 
**NearTxHashes** | **[]string** | All Near transactions executed for this swap | 
**AmountIn** | Pointer to **string** | Exact amount of &#x60;originToken&#x60; after trade was settled | [optional] 
**AmountInFormatted** | Pointer to **string** | Exact amount of &#x60;originToken&#x60; after trade was settled in readable format | [optional] 
**AmountInUsd** | Pointer to **string** | Exact amount of &#x60;originToken&#x60; equivalent in USD | [optional] 
**AmountOut** | Pointer to **string** | Exact amount of &#x60;destinationToken&#x60; after trade was settled | [optional] 
**AmountOutFormatted** | Pointer to **string** | Exact amount of &#x60;destinationToken&#x60; in readable format | [optional] 
**AmountOutUsd** | Pointer to **string** | Exact amount of &#x60;destinationToken&#x60; equivalent in USD | [optional] 
**Slippage** | Pointer to **float32** | Actual slippage | [optional] 
**OriginChainTxHashes** | [**[]TransactionDetails**](TransactionDetails.md) | Hashes and explorer URLs for all transactions on the origin chain | 
**DestinationChainTxHashes** | [**[]TransactionDetails**](TransactionDetails.md) | Hashes and explorer URLs for all transactions on the destination chain | 
**RefundedAmount** | Pointer to **string** | Amount of &#x60;originAsset&#x60; that got transferred to &#x60;refundTo&#x60; | [optional] 
**RefundedAmountFormatted** | Pointer to **string** | Refunded amount in readable format | [optional] 
**RefundedAmountUsd** | Pointer to **string** | Refunded amount equivalent in USD | [optional] 

## Methods

### NewSwapDetails

`func NewSwapDetails(intentHashes []string, nearTxHashes []string, originChainTxHashes []TransactionDetails, destinationChainTxHashes []TransactionDetails, ) *SwapDetails`

NewSwapDetails instantiates a new SwapDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapDetailsWithDefaults

`func NewSwapDetailsWithDefaults() *SwapDetails`

NewSwapDetailsWithDefaults instantiates a new SwapDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntentHashes

`func (o *SwapDetails) GetIntentHashes() []string`

GetIntentHashes returns the IntentHashes field if non-nil, zero value otherwise.

### GetIntentHashesOk

`func (o *SwapDetails) GetIntentHashesOk() (*[]string, bool)`

GetIntentHashesOk returns a tuple with the IntentHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentHashes

`func (o *SwapDetails) SetIntentHashes(v []string)`

SetIntentHashes sets IntentHashes field to given value.


### GetNearTxHashes

`func (o *SwapDetails) GetNearTxHashes() []string`

GetNearTxHashes returns the NearTxHashes field if non-nil, zero value otherwise.

### GetNearTxHashesOk

`func (o *SwapDetails) GetNearTxHashesOk() (*[]string, bool)`

GetNearTxHashesOk returns a tuple with the NearTxHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearTxHashes

`func (o *SwapDetails) SetNearTxHashes(v []string)`

SetNearTxHashes sets NearTxHashes field to given value.


### GetAmountIn

`func (o *SwapDetails) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *SwapDetails) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *SwapDetails) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.

### HasAmountIn

`func (o *SwapDetails) HasAmountIn() bool`

HasAmountIn returns a boolean if a field has been set.

### GetAmountInFormatted

`func (o *SwapDetails) GetAmountInFormatted() string`

GetAmountInFormatted returns the AmountInFormatted field if non-nil, zero value otherwise.

### GetAmountInFormattedOk

`func (o *SwapDetails) GetAmountInFormattedOk() (*string, bool)`

GetAmountInFormattedOk returns a tuple with the AmountInFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInFormatted

`func (o *SwapDetails) SetAmountInFormatted(v string)`

SetAmountInFormatted sets AmountInFormatted field to given value.

### HasAmountInFormatted

`func (o *SwapDetails) HasAmountInFormatted() bool`

HasAmountInFormatted returns a boolean if a field has been set.

### GetAmountInUsd

`func (o *SwapDetails) GetAmountInUsd() string`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *SwapDetails) GetAmountInUsdOk() (*string, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *SwapDetails) SetAmountInUsd(v string)`

SetAmountInUsd sets AmountInUsd field to given value.

### HasAmountInUsd

`func (o *SwapDetails) HasAmountInUsd() bool`

HasAmountInUsd returns a boolean if a field has been set.

### GetAmountOut

`func (o *SwapDetails) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *SwapDetails) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *SwapDetails) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.

### HasAmountOut

`func (o *SwapDetails) HasAmountOut() bool`

HasAmountOut returns a boolean if a field has been set.

### GetAmountOutFormatted

`func (o *SwapDetails) GetAmountOutFormatted() string`

GetAmountOutFormatted returns the AmountOutFormatted field if non-nil, zero value otherwise.

### GetAmountOutFormattedOk

`func (o *SwapDetails) GetAmountOutFormattedOk() (*string, bool)`

GetAmountOutFormattedOk returns a tuple with the AmountOutFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutFormatted

`func (o *SwapDetails) SetAmountOutFormatted(v string)`

SetAmountOutFormatted sets AmountOutFormatted field to given value.

### HasAmountOutFormatted

`func (o *SwapDetails) HasAmountOutFormatted() bool`

HasAmountOutFormatted returns a boolean if a field has been set.

### GetAmountOutUsd

`func (o *SwapDetails) GetAmountOutUsd() string`

GetAmountOutUsd returns the AmountOutUsd field if non-nil, zero value otherwise.

### GetAmountOutUsdOk

`func (o *SwapDetails) GetAmountOutUsdOk() (*string, bool)`

GetAmountOutUsdOk returns a tuple with the AmountOutUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOutUsd

`func (o *SwapDetails) SetAmountOutUsd(v string)`

SetAmountOutUsd sets AmountOutUsd field to given value.

### HasAmountOutUsd

`func (o *SwapDetails) HasAmountOutUsd() bool`

HasAmountOutUsd returns a boolean if a field has been set.

### GetSlippage

`func (o *SwapDetails) GetSlippage() float32`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *SwapDetails) GetSlippageOk() (*float32, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *SwapDetails) SetSlippage(v float32)`

SetSlippage sets Slippage field to given value.

### HasSlippage

`func (o *SwapDetails) HasSlippage() bool`

HasSlippage returns a boolean if a field has been set.

### GetOriginChainTxHashes

`func (o *SwapDetails) GetOriginChainTxHashes() []TransactionDetails`

GetOriginChainTxHashes returns the OriginChainTxHashes field if non-nil, zero value otherwise.

### GetOriginChainTxHashesOk

`func (o *SwapDetails) GetOriginChainTxHashesOk() (*[]TransactionDetails, bool)`

GetOriginChainTxHashesOk returns a tuple with the OriginChainTxHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginChainTxHashes

`func (o *SwapDetails) SetOriginChainTxHashes(v []TransactionDetails)`

SetOriginChainTxHashes sets OriginChainTxHashes field to given value.


### GetDestinationChainTxHashes

`func (o *SwapDetails) GetDestinationChainTxHashes() []TransactionDetails`

GetDestinationChainTxHashes returns the DestinationChainTxHashes field if non-nil, zero value otherwise.

### GetDestinationChainTxHashesOk

`func (o *SwapDetails) GetDestinationChainTxHashesOk() (*[]TransactionDetails, bool)`

GetDestinationChainTxHashesOk returns a tuple with the DestinationChainTxHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationChainTxHashes

`func (o *SwapDetails) SetDestinationChainTxHashes(v []TransactionDetails)`

SetDestinationChainTxHashes sets DestinationChainTxHashes field to given value.


### GetRefundedAmount

`func (o *SwapDetails) GetRefundedAmount() string`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *SwapDetails) GetRefundedAmountOk() (*string, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *SwapDetails) SetRefundedAmount(v string)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *SwapDetails) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetRefundedAmountFormatted

`func (o *SwapDetails) GetRefundedAmountFormatted() string`

GetRefundedAmountFormatted returns the RefundedAmountFormatted field if non-nil, zero value otherwise.

### GetRefundedAmountFormattedOk

`func (o *SwapDetails) GetRefundedAmountFormattedOk() (*string, bool)`

GetRefundedAmountFormattedOk returns a tuple with the RefundedAmountFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmountFormatted

`func (o *SwapDetails) SetRefundedAmountFormatted(v string)`

SetRefundedAmountFormatted sets RefundedAmountFormatted field to given value.

### HasRefundedAmountFormatted

`func (o *SwapDetails) HasRefundedAmountFormatted() bool`

HasRefundedAmountFormatted returns a boolean if a field has been set.

### GetRefundedAmountUsd

`func (o *SwapDetails) GetRefundedAmountUsd() string`

GetRefundedAmountUsd returns the RefundedAmountUsd field if non-nil, zero value otherwise.

### GetRefundedAmountUsdOk

`func (o *SwapDetails) GetRefundedAmountUsdOk() (*string, bool)`

GetRefundedAmountUsdOk returns a tuple with the RefundedAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmountUsd

`func (o *SwapDetails) SetRefundedAmountUsd(v string)`

SetRefundedAmountUsd sets RefundedAmountUsd field to given value.

### HasRefundedAmountUsd

`func (o *SwapDetails) HasRefundedAmountUsd() bool`

HasRefundedAmountUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


