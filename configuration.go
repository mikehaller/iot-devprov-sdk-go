/*
 * Bosch IoT Suite - Device Provisioning
 *
 * The Device Provisioning API is used to provision new devices for the Bosch IoT Suite services. It simplifies provisioning by registering the device in all affected services with a single API call.  Currently the Bosch IoT Hub and Bosch IoT Things service are supported. The following entities can be provisioned: * **Device** (Bosch IoT Hub): The Bosch IoT Hub has to be made aware of the Device which will connect to it and send telemetry data. * **Credentials** (Bosch IoT Hub): Credentials are required to authenticate Devices which want to connect. If Devices connect via a gateway, only credentials for the gateway are required. Hence, Credentials are optional in the Device Provisioning API. * **Thing** (Bosch IoT Things): A Thing is a *Digital Twin*: Among other things, it stores the state of the Device and provides an API for it, which is also called *Device-as-a-Service*. [Vorto models](https://vorto.eclipse.org) can be used to define the capabilities of a Thing. * **Policy** (Bosch IoT Things): Each Thing must have a reference (`policyId`) to a Policy which defines its access  control. You can decide to create a specific Policy for each Thing or to re-use a Policy for multiple Things.  Authentication and authorization of the Device Provisioning API is based on *Suite Auth* access tokens. You have to create a Suite Auth access token which provides full access to the Bosch IoT Hub and Bosch IoT Things service instances you want to address and provide it as *Bearer Authorization*.  **Note**: If you are using the Asset communication package hybrid offering please use the following server: - https://deviceprovisioning.eu-h1.bosch-iot-suite.com/api/1 
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotdevprov

import (
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")

)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}


// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	Url string
	Description string
	Variables map[string]ServerVariable
}

// Configuration stores the configuration of the API client
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	Servers       []ServerConfiguration
	HTTPClient    *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://deviceprovisioning.eu-1.bosch-iot-suite.com/api/1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers:       []ServerConfiguration{
			{
				Url: "https://deviceprovisioning.eu-1.bosch-iot-suite.com/api/1",
				Description: "No description provided",
			},
			{
				Url: "https://deviceprovisioning.eu-h1.bosch-iot-suite.com/api/1",
				Description: "No description provided",
			},
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// ServerUrl returns URL based on server settings
func (c *Configuration) ServerUrl(index int, variables map[string]string) (string, error) {
	if index < 0 || len(c.Servers) <= index {
		return "", fmt.Errorf("Index %v out of range %v", index, len(c.Servers) - 1)
	}
	server := c.Servers[index]
	url := server.Url

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("The variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}
