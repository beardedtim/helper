package http

import (
	"os"

	"github.com/juju/errors"

	"mckp/helper/datastore"
	"mckp/helper/repositories"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
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

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Routes) CreateUser(ctx *gin.Context, newUser *CreateUserParams) (datastore.PublicUser, error) {
	userRepo := repositories.UserRepository{}

	return userRepo.Create(newUser.Email, newUser.Password)
}

type GetUserByIdParams struct {
	ID string `path:"id"`
}

func (r *Routes) GetUserById(ctx *gin.Context, params *GetUserByIdParams) (datastore.PublicUser, error) {
	userRepo := repositories.UserRepository{}

	return userRepo.GetById(params.ID)
}

type GetUserTokenParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserTokenResult struct {
	Token string `json:"token"`
}

func (r *Routes) GetUserToken(ctx *gin.Context, params *GetUserTokenParams) (GetUserTokenResult, error) {
	userRepo := repositories.UserRepository{}

	user, err := userRepo.ValidatePassword(params.Email, params.Password)

	if err != nil {
		return GetUserTokenResult{}, errors.BadRequestf("passwords do not match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	log.WithError(err).Warn("Error trting to sign string")

	return GetUserTokenResult{
		Token: tokenString,
	}, err
}
