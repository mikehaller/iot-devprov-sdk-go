# X509Cert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthId** | **string** | The identity that the device should be authenticated as. | [optional] 
**Enabled** | **bool** | If set to false the credentials are not supposed to be used to authenticate devices any longer. This may e.g. be used to disable a particular mechanism for authenticating the device. | [optional] [default to true]
**DeviceId** | **string** | The ID of the device to which the credentials belong. | [optional] 
**Type** | **string** |  | 
**Secrets** | [**[]OneOfsecretCertificate**](oneOf&lt;secretCertificate&gt;.md) | A list of secrets scoped to a particular time period. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


