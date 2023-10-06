package repositories

import (
	"mckp/helper/datastore"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juju/errors"
)

type UserRepository struct{}

func (ur *UserRepository) Create(email string, password string) (datastore.PublicUser, error) {
	user := datastore.UsersModel{
		Email:    email,
		Password: password,
	}

	result := datastore.DatastoreInstance.Database.Create(&user)

	if result.Error != nil {
		if result.Error.Error() == "duplicated key not allowed" {
			return datastore.PublicUser{}, errors.AlreadyExistsf("email \"%s\"", email)
		}

		return datastore.PublicUser{}, result.Error
	}

	savedUser := datastore.PublicUser{
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ID:        user.ID,
	}

	return savedUser, nil
}

func (ur *UserRepository) GetById(id string) (datastore.PublicUser, error) {
	user := datastore.UsersModel{}

	result := datastore.DatastoreInstance.Database.First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return datastore.PublicUser{}, errors.NotFoundf("id \"%s\"", id)
		}

		return datastore.PublicUser{}, result.Error
	}

	publicUser := datastore.PublicUser{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return publicUser, nil
}

func (ur *UserRepository) ValidatePassword(email string, password string) (datastore.PublicUser, error) {
	model := datastore.UsersModel{}

	return model.PasswordsMatch(email, password)
}

type TokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func (ur *UserRepository) CreateToken(id string) (string, error) {
	claims := TokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (ur *UserRepository) ParseToken(token string) (TokenClaims, error) {
	claims := &TokenClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return *claims, err
}
