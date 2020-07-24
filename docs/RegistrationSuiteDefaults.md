# RegistrationSuiteDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | The content-type to use for northbound messages published by this device. This can be useful for protocols like MQTT where the content-type cannot be set directly in a message. The content-type is specified in RFC-2046. Example value: &#39;application/json&#39; | [optional] 
**ContentEncoding** | **string** | The content-encoding can be set optionally in addition to the content-type. It defines what additional content encodings have to be applied to the message content to receive the media-type referenced by the content-type. The main usage of the content-encoding is to allow compression without losing information about the compressed data type. The content-encoding is specified in section 3.5 of RFC 2616. | [optional] 
**Ttl** | **int32** | The time-to-live (in seconds) to use for events published by this device. This value is limited by the booked service plan&#39;s event storage time. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


