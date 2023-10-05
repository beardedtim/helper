package http

import (
	"github.com/juju/errors"

	"mckp/helper/datastore"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

var HTTPRoutes = Routes{}

type HealthcheckResponse struct {
	Healthy bool `json:"healthy"`
}

func (r *Routes) Healthcheck() func(*gin.Context) (HealthcheckResponse, error) {
	return func(c *gin.Context) (HealthcheckResponse, error) {
		healthy := datastore.DatastoreInstance.IsHealthy()

		if healthy {
			return HealthcheckResponse{
				Healthy: healthy,
			}, nil
		}

		return HealthcheckResponse{
			Healthy: healthy,
		}, errors.New("not healthy")
	}
}
