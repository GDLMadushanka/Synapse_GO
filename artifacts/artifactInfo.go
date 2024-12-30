package artifacts

import "sync"

// ConfigurationContext struct which holds the deployed artifact details
type ArtifactInfo struct {
	ApiMap      map[string]API
	EndpointMap map[string]Endpoint
}

// AddAPI method to add an API to the configuration context
func (c *ArtifactInfo) AddAPI(api API) {
	c.ApiMap[api.Name] = api
}

// AddEndpoint method to add an Endpoint to the configuration context
func (c *ArtifactInfo) AddEndpoint(endpoint Endpoint) {
	c.EndpointMap[endpoint.Name] = endpoint
}

var instance *ArtifactInfo

var once sync.Once

func GetArtifactInfoInstance() *ArtifactInfo {
	once.Do(func() {
		instance = &ArtifactInfo{
			ApiMap:      make(map[string]API),
			EndpointMap: make(map[string]Endpoint),
		}
	})
	return instance
}
