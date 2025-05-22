# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | Unique asset identifier | 
**Decimals** | **float32** | Number of decimals for the token | 
**Blockchain** | **string** | Blockchain associated with the token | 
**Symbol** | **string** | Token symbol (e.g. BTC, ETH) | 
**Price** | **float32** | Current price of the token in USD | 
**PriceUpdatedAt** | **time.Time** | Date when the token price was last updated | 
**ContractAddress** | Pointer to **string** | Contract address of the token | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse(assetId string, decimals float32, blockchain string, symbol string, price float32, priceUpdatedAt time.Time, ) *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *TokenResponse) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TokenResponse) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TokenResponse) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetDecimals

`func (o *TokenResponse) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenResponse) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenResponse) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.


### GetBlockchain

`func (o *TokenResponse) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TokenResponse) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TokenResponse) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetSymbol

`func (o *TokenResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetPrice

`func (o *TokenResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TokenResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TokenResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPriceUpdatedAt

`func (o *TokenResponse) GetPriceUpdatedAt() time.Time`

GetPriceUpdatedAt returns the PriceUpdatedAt field if non-nil, zero value otherwise.

### GetPriceUpdatedAtOk

`func (o *TokenResponse) GetPriceUpdatedAtOk() (*time.Time, bool)`

GetPriceUpdatedAtOk returns a tuple with the PriceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUpdatedAt

`func (o *TokenResponse) SetPriceUpdatedAt(v time.Time)`

SetPriceUpdatedAt sets PriceUpdatedAt field to given value.


### GetContractAddress

`func (o *TokenResponse) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokenResponse) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokenResponse) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *TokenResponse) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


