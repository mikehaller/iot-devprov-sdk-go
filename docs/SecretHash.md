# SecretHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotBefore** | **string** | The point in time from which the secret may be used to authenticate devices. This is optional and can be ignored if the validity period is the same as the period indicated by the client certificate’s corresponding properties. If not null, the value MUST be an ISO 8601 compliant combined date and time representation. | [optional] 
**NotAfter** | **string** | The point in time until which the secret may be used to authenticate devices. This is optional and can be ignored if the validity period is the same as the period indicated by the client certificate’s corresponding properties. If not null, the value MUST be an ISO 8601 compliant combined date and time representation. | [optional] 
**PwdHash** | **string** | __DEPRECATED!__ The Base64 encoded bytes representing the hashed password. The password hash MUST be computed by applying the hash function to the byte array consisting of the salt bytes (if a salt is used) and the UTF-8 encoding of the clear text password. | 
**Salt** | **string** | __DEPRECATED!__ The Base64 encoded bytes used as salt for the password hash. If not set then the password hash has been created without salt. | [optional] 
**HashFunction** | **string** | __DEPRECATED!__ The name of the hash function used to create the password hash. At the moment sha-256 and sha-512 are supported. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


