/*
 * Bosch IoT Suite - Device Provisioning
 *
 * The Device Provisioning API is used to provision new devices for the Bosch IoT Suite services. It simplifies provisioning by registering the device in all affected services with a single API call.  Currently the Bosch IoT Hub and Bosch IoT Things service are supported. The following entities can be provisioned: * **Device** (Bosch IoT Hub): The Bosch IoT Hub has to be made aware of the Device which will connect to it and send telemetry data. * **Credentials** (Bosch IoT Hub): Credentials are required to authenticate Devices which want to connect. If Devices connect via a gateway, only credentials for the gateway are required. Hence, Credentials are optional in the Device Provisioning API. * **Thing** (Bosch IoT Things): A Thing is a *Digital Twin*: Among other things, it stores the state of the Device and provides an API for it, which is also called *Device-as-a-Service*. [Vorto models](https://vorto.eclipse.org) can be used to define the capabilities of a Thing. * **Policy** (Bosch IoT Things): Each Thing must have a reference (`policyId`) to a Policy which defines its access  control. You can decide to create a specific Policy for each Thing or to re-use a Policy for multiple Things.  Authentication and authorization of the Device Provisioning API is based on *Suite Auth* access tokens. You have to create a Suite Auth access token which provides full access to the Bosch IoT Hub and Bosch IoT Things service instances you want to address and provide it as *Bearer Authorization*.  **Note**: If you are using the Asset communication package hybrid offering please use the following server: - https://deviceprovisioning.eu-h1.bosch-iot-suite.com/api/1 
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotdevprov
// HashedPasswordAllOf struct for HashedPasswordAllOf
type HashedPasswordAllOf struct {
	Type string `json:"type"`
	// A list of secrets scoped to a particular time period. This array must contain at least one element - an empty array is handled as an error.
	Secrets []OneOfsecretHashsecretPasswordsecretPasswordBase64 `json:"secrets"`
}
