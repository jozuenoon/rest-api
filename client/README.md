# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/jozuenoone/rest-api](https://github.com/jozuenoone/rest-api)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PaymentServiceApi* | [**CreatePayment**](docs/PaymentServiceApi.md#createpayment) | **Post** /v1/transaction/payments | 
*PaymentServiceApi* | [**DeletePayment**](docs/PaymentServiceApi.md#deletepayment) | **Delete** /v1/transaction/payments/{id} | 
*PaymentServiceApi* | [**GetPayment**](docs/PaymentServiceApi.md#getpayment) | **Get** /v1/transaction/payments/{id} | 
*PaymentServiceApi* | [**GetPayment2**](docs/PaymentServiceApi.md#getpayment2) | **Get** /v1/transaction/payments | 
*PaymentServiceApi* | [**UpdatePayment**](docs/PaymentServiceApi.md#updatepayment) | **Patch** /v1/transaction/payments/{id} | 


## Documentation For Models

 - [ChargersInformationSenderCharges](docs/ChargersInformationSenderCharges.md)
 - [PaymentapiBeneficiaryParty](docs/PaymentapiBeneficiaryParty.md)
 - [PaymentapiChargersInformation](docs/PaymentapiChargersInformation.md)
 - [PaymentapiDebtorParty](docs/PaymentapiDebtorParty.md)
 - [PaymentapiError](docs/PaymentapiError.md)
 - [PaymentapiFx](docs/PaymentapiFx.md)
 - [PaymentapiPayment](docs/PaymentapiPayment.md)
 - [PaymentapiPaymentAttributes](docs/PaymentapiPaymentAttributes.md)
 - [PaymentapiPaymentRequest](docs/PaymentapiPaymentRequest.md)
 - [PaymentapiPaymentServiceResponse](docs/PaymentapiPaymentServiceResponse.md)
 - [PaymentapiSponsorParty](docs/PaymentapiSponsorParty.md)


## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

none@example.com

