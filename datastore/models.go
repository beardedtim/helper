package datastore

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"primaryKey,type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"unique"`
	Password  string
}

type PublicUser struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updaated_at"`
	Email     string    `json:"email"`
}

// Before we Create the value in the database, we need to
// hash the password
func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return
}

// Before we Update the value in the database, we need to
// hash the password
func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// but only if the password has changed since
	// if it hasn't changed, we don't want to hash
	// the hash
	if tx.Statement.Changed("Password") {
		bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

		if err != nil {
			return err
		}

		u.Password = string(bytes)
	}

	return
}

func (u *UserModel) PasswordsMatch(email string, plainTextPassword string) (PublicUser, error) {
	DatastoreInstance.Database.First(&u, "email = ?", email)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainTextPassword))

	if err != nil {
		return PublicUser{}, err
	}

	return PublicUser{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (u *UserModel) TableName() string {
	return "users"
}

type Models struct {
	Users UserModel
}
