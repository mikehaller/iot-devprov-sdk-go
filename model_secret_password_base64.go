/*
 * Bosch IoT Suite - Device Provisioning
 *
 * The Device Provisioning API is used to provision new devices for the Bosch IoT Suite services. It simplifies provisioning by registering the device in all affected services with a single API call.  Currently the Bosch IoT Hub and Bosch IoT Things service are supported. The following entities can be provisioned: * **Device** (Bosch IoT Hub): The Bosch IoT Hub has to be made aware of the Device which will connect to it and send telemetry data. * **Credentials** (Bosch IoT Hub): Credentials are required to authenticate Devices which want to connect. If Devices connect via a gateway, only credentials for the gateway are required. Hence, Credentials are optional in the Device Provisioning API. * **Thing** (Bosch IoT Things): A Thing is a *Digital Twin*: Among other things, it stores the state of the Device and provides an API for it, which is also called *Device-as-a-Service*. [Vorto models](https://vorto.eclipse.org) can be used to define the capabilities of a Thing. * **Policy** (Bosch IoT Things): Each Thing must have a reference (`policyId`) to a Policy which defines its access  control. You can decide to create a specific Policy for each Thing or to re-use a Policy for multiple Things.  Authentication and authorization of the Device Provisioning API is based on *Suite Auth* access tokens. You have to create a Suite Auth access token which provides full access to the Bosch IoT Hub and Bosch IoT Things service instances you want to address and provide it as *Bearer Authorization*.  **Note**: If you are using the Asset communication package hybrid offering please use the following server: - https://deviceprovisioning.eu-h1.bosch-iot-suite.com/api/1 
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotdevprov
// SecretPasswordBase64 Data structure representing a base64 encoded plaintext secret of a credential.
type SecretPasswordBase64 struct {
	// The point in time from which the secret may be used to authenticate devices. This is optional and can be ignored if the validity period is the same as the period indicated by the client certificate’s corresponding properties. If not null, the value MUST be an ISO 8601 compliant combined date and time representation.
	NotBefore string `json:"notBefore,omitempty"`
	// The point in time until which the secret may be used to authenticate devices. This is optional and can be ignored if the validity period is the same as the period indicated by the client certificate’s corresponding properties. If not null, the value MUST be an ISO 8601 compliant combined date and time representation.
	NotAfter string `json:"notAfter,omitempty"`
	// Plaintext base64 encoded password. Will be salted and hashed by Bosch IoT Hub.
	PasswordBase64 string `json:"passwordBase64"`
}
