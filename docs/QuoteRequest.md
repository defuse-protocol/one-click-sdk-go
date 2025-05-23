# QuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dry** | **bool** | Flag indicating whether this is a dry run request. If &#x60;true&#x60;, the response will **NOT** contain the following fields: - &#x60;depositAddress&#x60; - &#x60;timeWhenInactive&#x60; - &#x60;deadline&#x60; | 
**SwapType** | **string** | Whether to use the amount as the output or the input for the basis of the swap: - &#x60;EXACT_INPUT&#x60; - request output amount for exact input. - &#x60;EXACT_OUTPUT&#x60; - request output amount for exact output. The &#x60;refundTo&#x60; address will always receive excess tokens back even after the swap is complete. | 
**SlippageTolerance** | **float32** | Slippage tolerance for the swap. This value is in basis points (1/100th of a percent), e.g. 100 for 1% slippage. | 
**OriginAsset** | **string** | ID of the origin asset. | 
**DepositType** | **string** | Type of the deposit address: - &#x60;ORIGIN_CHAIN&#x60; - deposit address on the origin chain - &#x60;INTENTS&#x60; - **account ID** inside near intents to which you should transfer assets inside intents. | 
**DestinationAsset** | **string** | ID of the destination asset. | 
**Amount** | **string** | Amount to swap as the base amount (can be switched to exact input/output using the dedicated flag), denoted in the smallest unit of the specified currency (e.g., wei for ETH). | 
**RefundTo** | **string** | Address for user refund. | 
**RefundType** | **string** | Type of refund address: - &#x60;ORIGIN_CHAIN&#x60; - assets will be refunded to &#x60;refundTo&#x60; address on the origin chain - &#x60;INTENTS&#x60; - assets will be refunded to &#x60;refundTo&#x60; intents account | 
**Recipient** | **string** | Recipient address. The format should match &#x60;recipientType&#x60;. | 
**RecipientType** | **string** | Type of recipient address: - &#x60;DESTINATION_CHAIN&#x60; - assets will be transferred to chain of &#x60;destinationAsset&#x60; - &#x60;INTENTS&#x60; - assets will be transferred to account inside intents | 
**Deadline** | **time.Time** | Timestamp in ISO format, that identifies when user refund will begin if the swap isn&#39;t completed by then. | 
**Referral** | Pointer to **string** | Referral identifier(lower case only) | [optional] 
**QuoteWaitingTimeMs** | Pointer to **float32** | Time in milliseconds user is willing to wait for quote from relay. | [optional] [default to 3000]
**AppFees** | Pointer to [**[]AppFee**](AppFee.md) | List of recipients and their fees | [optional] 

## Methods

### NewQuoteRequest

`func NewQuoteRequest(dry bool, swapType string, slippageTolerance float32, originAsset string, depositType string, destinationAsset string, amount string, refundTo string, refundType string, recipient string, recipientType string, deadline time.Time, ) *QuoteRequest`

NewQuoteRequest instantiates a new QuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRequestWithDefaults

`func NewQuoteRequestWithDefaults() *QuoteRequest`

NewQuoteRequestWithDefaults instantiates a new QuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDry

`func (o *QuoteRequest) GetDry() bool`

GetDry returns the Dry field if non-nil, zero value otherwise.

### GetDryOk

`func (o *QuoteRequest) GetDryOk() (*bool, bool)`

GetDryOk returns a tuple with the Dry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDry

`func (o *QuoteRequest) SetDry(v bool)`

SetDry sets Dry field to given value.


### GetSwapType

`func (o *QuoteRequest) GetSwapType() string`

GetSwapType returns the SwapType field if non-nil, zero value otherwise.

### GetSwapTypeOk

`func (o *QuoteRequest) GetSwapTypeOk() (*string, bool)`

GetSwapTypeOk returns a tuple with the SwapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapType

`func (o *QuoteRequest) SetSwapType(v string)`

SetSwapType sets SwapType field to given value.


### GetSlippageTolerance

`func (o *QuoteRequest) GetSlippageTolerance() float32`

GetSlippageTolerance returns the SlippageTolerance field if non-nil, zero value otherwise.

### GetSlippageToleranceOk

`func (o *QuoteRequest) GetSlippageToleranceOk() (*float32, bool)`

GetSlippageToleranceOk returns a tuple with the SlippageTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageTolerance

`func (o *QuoteRequest) SetSlippageTolerance(v float32)`

SetSlippageTolerance sets SlippageTolerance field to given value.


### GetOriginAsset

`func (o *QuoteRequest) GetOriginAsset() string`

GetOriginAsset returns the OriginAsset field if non-nil, zero value otherwise.

### GetOriginAssetOk

`func (o *QuoteRequest) GetOriginAssetOk() (*string, bool)`

GetOriginAssetOk returns a tuple with the OriginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAsset

`func (o *QuoteRequest) SetOriginAsset(v string)`

SetOriginAsset sets OriginAsset field to given value.


### GetDepositType

`func (o *QuoteRequest) GetDepositType() string`

GetDepositType returns the DepositType field if non-nil, zero value otherwise.

### GetDepositTypeOk

`func (o *QuoteRequest) GetDepositTypeOk() (*string, bool)`

GetDepositTypeOk returns a tuple with the DepositType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositType

`func (o *QuoteRequest) SetDepositType(v string)`

SetDepositType sets DepositType field to given value.


### GetDestinationAsset

`func (o *QuoteRequest) GetDestinationAsset() string`

GetDestinationAsset returns the DestinationAsset field if non-nil, zero value otherwise.

### GetDestinationAssetOk

`func (o *QuoteRequest) GetDestinationAssetOk() (*string, bool)`

GetDestinationAssetOk returns a tuple with the DestinationAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAsset

`func (o *QuoteRequest) SetDestinationAsset(v string)`

SetDestinationAsset sets DestinationAsset field to given value.


### GetAmount

`func (o *QuoteRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuoteRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuoteRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRefundTo

`func (o *QuoteRequest) GetRefundTo() string`

GetRefundTo returns the RefundTo field if non-nil, zero value otherwise.

### GetRefundToOk

`func (o *QuoteRequest) GetRefundToOk() (*string, bool)`

GetRefundToOk returns a tuple with the RefundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTo

`func (o *QuoteRequest) SetRefundTo(v string)`

SetRefundTo sets RefundTo field to given value.


### GetRefundType

`func (o *QuoteRequest) GetRefundType() string`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *QuoteRequest) GetRefundTypeOk() (*string, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *QuoteRequest) SetRefundType(v string)`

SetRefundType sets RefundType field to given value.


### GetRecipient

`func (o *QuoteRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *QuoteRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *QuoteRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetRecipientType

`func (o *QuoteRequest) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *QuoteRequest) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *QuoteRequest) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.


### GetDeadline

`func (o *QuoteRequest) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *QuoteRequest) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *QuoteRequest) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetReferral

`func (o *QuoteRequest) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *QuoteRequest) GetReferralOk() (*string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *QuoteRequest) SetReferral(v string)`

SetReferral sets Referral field to given value.

### HasReferral

`func (o *QuoteRequest) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### GetQuoteWaitingTimeMs

`func (o *QuoteRequest) GetQuoteWaitingTimeMs() float32`

GetQuoteWaitingTimeMs returns the QuoteWaitingTimeMs field if non-nil, zero value otherwise.

### GetQuoteWaitingTimeMsOk

`func (o *QuoteRequest) GetQuoteWaitingTimeMsOk() (*float32, bool)`

GetQuoteWaitingTimeMsOk returns a tuple with the QuoteWaitingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteWaitingTimeMs

`func (o *QuoteRequest) SetQuoteWaitingTimeMs(v float32)`

SetQuoteWaitingTimeMs sets QuoteWaitingTimeMs field to given value.

### HasQuoteWaitingTimeMs

`func (o *QuoteRequest) HasQuoteWaitingTimeMs() bool`

HasQuoteWaitingTimeMs returns a boolean if a field has been set.

### GetAppFees

`func (o *QuoteRequest) GetAppFees() []AppFee`

GetAppFees returns the AppFees field if non-nil, zero value otherwise.

### GetAppFeesOk

`func (o *QuoteRequest) GetAppFeesOk() (*[]AppFee, bool)`

GetAppFeesOk returns a tuple with the AppFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppFees

`func (o *QuoteRequest) SetAppFees(v []AppFee)`

SetAppFees sets AppFees field to given value.

### HasAppFees

`func (o *QuoteRequest) HasAppFees() bool`

HasAppFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


