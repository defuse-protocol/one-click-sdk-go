# QuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | Timestamp in ISO format that took part in deposit address derivation | 
**Signature** | **string** | Signature that means 1click service commit to quote and deposit address | 
**QuoteRequest** | [**QuoteRequest**](QuoteRequest.md) | User request | 
**Quote** | [**Quote**](Quote.md) | Response that contains deposit address for user to send \&quot;amount\&quot; of **originAsset** and possible output amount | 

## Methods

### NewQuoteResponse

`func NewQuoteResponse(timestamp time.Time, signature string, quoteRequest QuoteRequest, quote Quote, ) *QuoteResponse`

NewQuoteResponse instantiates a new QuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseWithDefaults

`func NewQuoteResponseWithDefaults() *QuoteResponse`

NewQuoteResponseWithDefaults instantiates a new QuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *QuoteResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QuoteResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QuoteResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSignature

`func (o *QuoteResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *QuoteResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *QuoteResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetQuoteRequest

`func (o *QuoteResponse) GetQuoteRequest() QuoteRequest`

GetQuoteRequest returns the QuoteRequest field if non-nil, zero value otherwise.

### GetQuoteRequestOk

`func (o *QuoteResponse) GetQuoteRequestOk() (*QuoteRequest, bool)`

GetQuoteRequestOk returns a tuple with the QuoteRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteRequest

`func (o *QuoteResponse) SetQuoteRequest(v QuoteRequest)`

SetQuoteRequest sets QuoteRequest field to given value.


### GetQuote

`func (o *QuoteResponse) GetQuote() Quote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *QuoteResponse) GetQuoteOk() (*Quote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *QuoteResponse) SetQuote(v Quote)`

SetQuote sets Quote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


