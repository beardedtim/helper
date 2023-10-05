package repositories

import (
	"mckp/helper/datastore"

	"github.com/juju/errors"
)

type UserRepository struct{}

func (ur *UserRepository) Create(email string, password string) (datastore.PublicUser, error) {
	user := datastore.UserModel{
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
	user := datastore.UserModel{}

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
	model := datastore.UserModel{}

	return model.PasswordsMatch(email, password)
}
