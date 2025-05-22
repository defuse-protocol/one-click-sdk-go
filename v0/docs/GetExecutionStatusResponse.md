# GetExecutionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteResponse** | [**QuoteResponse**](QuoteResponse.md) | Quote response from original request | 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** | Last time the state was updated | 
**SwapDetails** | [**SwapDetails**](SwapDetails.md) | Details of actual swaps and withdrawals | 

## Methods

### NewGetExecutionStatusResponse

`func NewGetExecutionStatusResponse(quoteResponse QuoteResponse, status string, updatedAt time.Time, swapDetails SwapDetails, ) *GetExecutionStatusResponse`

NewGetExecutionStatusResponse instantiates a new GetExecutionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExecutionStatusResponseWithDefaults

`func NewGetExecutionStatusResponseWithDefaults() *GetExecutionStatusResponse`

NewGetExecutionStatusResponseWithDefaults instantiates a new GetExecutionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteResponse

`func (o *GetExecutionStatusResponse) GetQuoteResponse() QuoteResponse`

GetQuoteResponse returns the QuoteResponse field if non-nil, zero value otherwise.

### GetQuoteResponseOk

`func (o *GetExecutionStatusResponse) GetQuoteResponseOk() (*QuoteResponse, bool)`

GetQuoteResponseOk returns a tuple with the QuoteResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteResponse

`func (o *GetExecutionStatusResponse) SetQuoteResponse(v QuoteResponse)`

SetQuoteResponse sets QuoteResponse field to given value.


### GetStatus

`func (o *GetExecutionStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExecutionStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExecutionStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *GetExecutionStatusResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetExecutionStatusResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetExecutionStatusResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSwapDetails

`func (o *GetExecutionStatusResponse) GetSwapDetails() SwapDetails`

GetSwapDetails returns the SwapDetails field if non-nil, zero value otherwise.

### GetSwapDetailsOk

`func (o *GetExecutionStatusResponse) GetSwapDetailsOk() (*SwapDetails, bool)`

GetSwapDetailsOk returns a tuple with the SwapDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapDetails

`func (o *GetExecutionStatusResponse) SetSwapDetails(v SwapDetails)`

SetSwapDetails sets SwapDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


