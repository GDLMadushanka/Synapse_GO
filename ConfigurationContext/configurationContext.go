package configurationcontext

import (
	"synapse/artifacts"
)

// ConfigurationContext struct which holds the deployed artifact details
type ConfigurationContext struct {
	ApiMap      map[string]artifacts.API
	EndpointMap map[string]artifacts.Endpoint
}

// AddAPI method to add an API to the configuration context
func (c *ConfigurationContext) AddAPI(api artifacts.API) {
	if c.ApiMap == nil {
		c.ApiMap = make(map[string]artifacts.API)
	}
	c.ApiMap[api.Name] = api
}

// AddEndpoint method to add an Endpoint to the configuration context
func (c *ConfigurationContext) AddEndpoint(endpoint artifacts.Endpoint) {
	if c.EndpointMap == nil {
		c.EndpointMap = make(map[string]artifacts.Endpoint)
	}
	c.EndpointMap[endpoint.Name] = endpoint
}
