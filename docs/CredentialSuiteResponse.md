# CredentialSuiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthId** | **string** | The identity that the device should be authenticated as. | [optional] 
**Enabled** | **bool** | If set to false the credentials are not supposed to be used to authenticate devices any longer. This may e.g. be used to disable a particular mechanism for authenticating the device. | [optional] [default to true]
**DeviceId** | **string** | The ID of the device to which the credentials belong. | [optional] 
**Type** | **string** |  | 
**Secrets** | [**[]OneOfsecretPsk**](oneOf&lt;secretPsk&gt;.md) | A list of secrets scoped to a particular time period. This array must contain at least one element - an empty array is handled as an error. Note: PSK is only supported in some protocol adapters. For details check the [Hub documentation](https://docs.bosch-iot-suite.com/hub/general-concepts/deviceauthentication.html) | 
**Adapters** | [**[]CredentialSuiteResponseAllOfAdapters**](credential_suite_response_allOf_adapters.md) | A list of protocol adapters that can be used with these credentials | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


