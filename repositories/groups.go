package repositories

import (
	"mckp/helper/datastore"

	"github.com/google/uuid"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

type GroupRepository struct{}

func (repo *GroupRepository) Create(name string, description string) (datastore.PublicGroup, error) {
	group := datastore.GroupsModel{
		Name:        name,
		Description: description,
	}

	result := datastore.DatastoreInstance.Database.Create(&group)

	if result.Error != nil {
		return datastore.PublicGroup{}, result.Error
	}

	publicGroup := datastore.PublicGroup{
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
	}

	return publicGroup, nil
}

func (repo *GroupRepository) GetById(id string) (datastore.PublicGroup, error) {
	group := datastore.GroupsModel{}

	result := datastore.DatastoreInstance.Database.Model(&datastore.GroupsModel{}).Preload("Members").First(&group, "id = ?", id)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return datastore.PublicGroup{}, errors.NotFoundf("id \"%s\"", id)
		}

		return datastore.PublicGroup{}, result.Error
	}

	groupUsers := []*datastore.PublicUser{}

	for _, dbUser := range group.Members {
		user := datastore.PublicUser{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			Email:     dbUser.Email,
		}

		groupUsers = append(groupUsers, &user)
	}

	publicGroup := datastore.PublicGroup{
		ID:          group.ID,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		Name:        group.Name,
		Description: group.Description,
		Members:     groupUsers,
	}

	return publicGroup, nil
}

func (repo *GroupRepository) AddUserToGroup(userId string, groupId string) (datastore.PublicGroup, error) {
	userInGroup, err := repo.IsUserInGroup(userId, groupId)

	if userInGroup {
		log.WithFields(log.Fields{
			"userId":  userId,
			"groupId": groupId,
		}).Debug("User already in Group")

		return repo.GetById(groupId)
	}

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"userId":  userId,
			"groupId": groupId,
		}).Warn("Error adding User to Group")

		return datastore.PublicGroup{}, err
	}

	group := datastore.GroupsModel{
		ID: uuid.MustParse(groupId),
	}

	datastore.DatastoreInstance.Database.Model(&group).Association("Members").Append(&datastore.UsersModel{
		ID: uuid.MustParse(userId),
	})

	return repo.GetById(groupId)
}

func (repo *GroupRepository) IsUserInGroup(userId string, groupId string) (bool, error) {
	user := datastore.UsersModel{}

	result := datastore.DatastoreInstance.Database.Preload("Groups").First(&user, "id = ?", userId)

	if result.Error != nil {
		return false, result.Error
	}

	for _, group := range user.Groups {
		if group.ID == uuid.MustParse(groupId) {
			return true, nil
		}
	}

	return false, nil
}

func (repo *GroupRepository) IsUserGroupAdmin(userId string, groupId string) (bool, error) {
	models := datastore.Models{}
	adminRole := models.Roles

	// Step 1: find the role that is labeled ADMIN
	result := datastore.DatastoreInstance.Database.Find(&adminRole, "name = ?", "admin")

	if result.Error != nil {
		return false, result.Error
	}

	groupRole := models.GroupRoles

	// Get the GroupRole for that user, group, and role
	result = datastore.DatastoreInstance.Database.Find(
		&groupRole,
		"user_id = ? AND group_id = ? AND role_id = ?",
		userId,
		groupId,
		adminRole.ID,
	)

	if result.Error != nil {
		return false, result.Error
	}

	// If the IDs do not match the user is not admin
	if groupRole.UserId != uuid.MustParse(userId) {
		return false, nil
	}

	// If the IDs do match, the user is admin
	return true, nil
}
