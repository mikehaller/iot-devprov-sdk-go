# \DeProvisioningApi

All URIs are relative to *https://deviceprovisioning.eu-1.bosch-iot-suite.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceInstanceIdDevicesDeviceIdDelete**](DeProvisioningApi.md#ServiceInstanceIdDevicesDeviceIdDelete) | **Delete** /{service-instance-id}/devices/{device-id} | De-provisioning a device: Deletes the device, thing, policy and credentials by ID in the underlying Bosch IoT Suite services. 



## ServiceInstanceIdDevicesDeviceIdDelete

> DeprovisioningResponsePayload ServiceInstanceIdDevicesDeviceIdDelete(ctx, serviceInstanceId, deviceId, optional)

De-provisioning a device: Deletes the device, thing, policy and credentials by ID in the underlying Bosch IoT Suite services. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceInstanceId** | **string**| The ID of the Bosch IoT Suite for Asset Communication subscription for which you want to de-provision a device. Not the IDs of the underlying Bosch IoT Hub or Bosch IoT Things service instance.  | 
**deviceId** | **string**| The ID of the device you want to de-provision | 
 **optional** | ***ServiceInstanceIdDevicesDeviceIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServiceInstanceIdDevicesDeviceIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **keepCredentials** | **optional.Bool**| Pass &#39;true&#39; if you wish to keep the device&#39;s credentials | 
 **keepPolicy** | **optional.Bool**| Pass &#39;true&#39; if you wish to keep the thing&#39;s policy | 

### Return type

[**DeprovisioningResponsePayload**](DeprovisioningResponsePayload.md)

### Authorization

[SuiteAuth](../README.md#SuiteAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

