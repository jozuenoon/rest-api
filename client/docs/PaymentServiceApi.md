# \PaymentServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentServiceApi.md#CreatePayment) | **Post** /v1/transaction/payments | 
[**DeletePayment**](PaymentServiceApi.md#DeletePayment) | **Delete** /v1/transaction/payments/{id} | 
[**GetPayment**](PaymentServiceApi.md#GetPayment) | **Get** /v1/transaction/payments/{id} | 
[**GetPayment2**](PaymentServiceApi.md#GetPayment2) | **Get** /v1/transaction/payments | 
[**UpdatePayment**](PaymentServiceApi.md#UpdatePayment) | **Patch** /v1/transaction/payments/{id} | 


# **CreatePayment**
> PaymentapiPaymentServiceResponse CreatePayment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PaymentapiPaymentAttributes**](PaymentapiPaymentAttributes.md)|  | 

### Return type

[**PaymentapiPaymentServiceResponse**](paymentapiPaymentServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePayment**
> PaymentapiPaymentRequest DeletePayment(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**PaymentapiPaymentRequest**](paymentapiPaymentRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayment**
> PaymentapiPaymentServiceResponse GetPayment(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**PaymentapiPaymentServiceResponse**](paymentapiPaymentServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayment2**
> PaymentapiPaymentServiceResponse GetPayment2(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPayment2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPayment2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 

### Return type

[**PaymentapiPaymentServiceResponse**](paymentapiPaymentServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePayment**
> PaymentapiPaymentServiceResponse UpdatePayment(ctx, id, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**PaymentapiPaymentAttributes**](PaymentapiPaymentAttributes.md)|  | 

### Return type

[**PaymentapiPaymentServiceResponse**](paymentapiPaymentServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

