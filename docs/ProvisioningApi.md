# \ProvisioningApi

All URIs are relative to *https://deviceprovisioning.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceInstanceIdDevicesPost**](ProvisioningApi.md#ServiceInstanceIdDevicesPost) | **Post** /{service-instance-id}/devices | Provision a device: Creates the required resources in the underlying Bosch IoT Suite services. 



## ServiceInstanceIdDevicesPost

> ProvisioningResponsePayload ServiceInstanceIdDevicesPost(ctx, serviceInstanceId, optional)

Provision a device: Creates the required resources in the underlying Bosch IoT Suite services. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceInstanceId** | **string**| The ID of the Bosch IoT Suite for Asset Communication subscription for which you want to provision a device. Not the IDs of the underlying Bosch IoT Hub or Bosch IoT Things service instance.  | 
 **optional** | ***ServiceInstanceIdDevicesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServiceInstanceIdDevicesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authScope** | **optional.String**| Controls who owns the newly created Thing (i.e. which subject is used for the default entry of the created Policy). If this parameter is omitted the authenticated subject of the user who made the request is used (Things default).  Supported scopes: * &#x60;subscription&#x60;: full access is granted to all users who have access to the same IoT Suite subscription  | 
 **skipVorto** | **optional.Bool**| If set to true, the resolution of the Vorto model from the definition field of the Thing will not be attempted.  | 
 **provisioningRequestPayload** | [**optional.Interface of ProvisioningRequestPayload**](ProvisioningRequestPayload.md)|  | 

### Return type

[**ProvisioningResponsePayload**](ProvisioningResponsePayload.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

