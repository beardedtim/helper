package http

import (
	"github.com/juju/errors"

	"mckp/helper/datastore"
	"mckp/helper/repositories"

	"github.com/gin-gonic/gin"
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

	user, _ := ctx.Get("User")

	log.WithField("claims", user).Info("Here you go")

	return userRepo.GetById(params.ID)
}

type GetUserTokenParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	token, err := userRepo.CreateToken(user.ID.String())

	return GetUserTokenResult{
		Token: token,
	}, err
}

type CreateGroupParams struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (r *Routes) CreateGroup(ctx *gin.Context, params *CreateGroupParams) (datastore.PublicGroup, error) {
	groupRepo := repositories.GroupRepository{}

	return groupRepo.Create(params.Name, params.Description)
}

type GetGroupByIdParams struct {
	ID string `path:"id"`
}

func (r *Routes) GetGroupById(ctx *gin.Context, params *GetGroupByIdParams) (datastore.PublicGroup, error) {
	groupRepo := repositories.GroupRepository{}

	requestingUser := ctx.MustGet("User").(repositories.TokenClaims)
	userInGroup, err := groupRepo.IsUserInGroup(requestingUser.ID, params.ID)

	if err != nil {
		return datastore.PublicGroup{}, err
	}

	if !userInGroup {
		return datastore.PublicGroup{}, errors.Unauthorizedf("user not in group")
	}

	return groupRepo.GetById(params.ID)
}

type AddUserToGroupParams struct {
	UserID  string `json:"user_id" binding:"required"`
	GroupID string `path:"id"`
}

func (r *Routes) AddUserToGroup(ctx *gin.Context, params *AddUserToGroupParams) (datastore.PublicGroup, error) {
	groupRepo := repositories.GroupRepository{}
	requestingUser := ctx.MustGet("User").(repositories.TokenClaims)

	usersIsAdmin, err := groupRepo.IsUserGroupAdmin(requestingUser.ID, params.GroupID)

	if err != nil {
		return datastore.PublicGroup{}, err
	}

	if !usersIsAdmin {
		return datastore.PublicGroup{}, errors.Unauthorizedf("user is not admin")
	}

	return groupRepo.AddUserToGroup(params.UserID, params.GroupID)
}

func (r *Routes) Test(ctx *gin.Context) (bool, error) {
	return true, nil
}
