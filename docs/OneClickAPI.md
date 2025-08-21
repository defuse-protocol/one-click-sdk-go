# \OneClickAPI

All URIs are relative to *https://1click.chaindefuser.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutionStatus**](OneClickAPI.md#GetExecutionStatus) | **Get** /v0/status | Check swap execution status
[**GetQuote**](OneClickAPI.md#GetQuote) | **Post** /v0/quote | Request a swap quote
[**GetTokens**](OneClickAPI.md#GetTokens) | **Get** /v0/tokens | Get supported tokens
[**SubmitDepositTx**](OneClickAPI.md#SubmitDepositTx) | **Post** /v0/deposit/submit | Submit deposit transaction hash



## GetExecutionStatus

> GetExecutionStatusResponse GetExecutionStatus(ctx).DepositAddress(depositAddress).DepositMemo(depositMemo).Execute()

Check swap execution status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/defuse-protocol/one-click-sdk-go"
)

func main() {
	depositAddress := "depositAddress_example" // string | 
	depositMemo := "depositMemo_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneClickAPI.GetExecutionStatus(context.Background()).DepositAddress(depositAddress).DepositMemo(depositMemo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneClickAPI.GetExecutionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutionStatus`: GetExecutionStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `OneClickAPI.GetExecutionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositAddress** | **string** |  | 
 **depositMemo** | **string** |  | 

### Return type

[**GetExecutionStatusResponse**](GetExecutionStatusResponse.md)

### Authorization

[JWT-auth](../README.md#JWT-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> QuoteResponse GetQuote(ctx).QuoteRequest(quoteRequest).Execute()

Request a swap quote



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/defuse-protocol/one-click-sdk-go"
)

func main() {
	quoteRequest := *openapiclient.NewQuoteRequest(true, "SwapType_example", float32(100), "nep141:arb-0xaf88d065e77c8cc2239327c5edb3a432268e5831.omft.near", "DepositType_example", "nep141:sol-5ce3bf3a31af18be40ba30f721101b4341690186.omft.near", "1000", "0x2527D02599Ba641c19FEa793cD0F167589a0f10D", "RefundType_example", "13QkxhNMrTPxoCkRdYdJ65tFuwXPhL5gLS2Z5Nr6gjRK", "RecipientType_example", time.Now()) // QuoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneClickAPI.GetQuote(context.Background()).QuoteRequest(quoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneClickAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `OneClickAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteRequest** | [**QuoteRequest**](QuoteRequest.md) |  | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[JWT-auth](../README.md#JWT-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> []TokenResponse GetTokens(ctx).Execute()

Get supported tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/defuse-protocol/one-click-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneClickAPI.GetTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneClickAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: []TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OneClickAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitDepositTx

> SubmitDepositTxResponse SubmitDepositTx(ctx).SubmitDepositTxRequest(submitDepositTxRequest).Execute()

Submit deposit transaction hash



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/defuse-protocol/one-click-sdk-go"
)

func main() {
	submitDepositTxRequest := *openapiclient.NewSubmitDepositTxRequest("0x123abc456def789", "0x2527D02599Ba641c19FEa793cD0F167589a0f10D") // SubmitDepositTxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneClickAPI.SubmitDepositTx(context.Background()).SubmitDepositTxRequest(submitDepositTxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneClickAPI.SubmitDepositTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitDepositTx`: SubmitDepositTxResponse
	fmt.Fprintf(os.Stdout, "Response from `OneClickAPI.SubmitDepositTx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitDepositTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitDepositTxRequest** | [**SubmitDepositTxRequest**](SubmitDepositTxRequest.md) |  | 

### Return type

[**SubmitDepositTxResponse**](SubmitDepositTxResponse.md)

### Authorization

[JWT-auth](../README.md#JWT-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

