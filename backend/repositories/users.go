package repositories

import (
	"serverless/interfaces"
	"serverless/models"
)

type UserRepository struct {
	databaseInstance *interfaces.IDbInstance
}

func NewUserRepository(dbInstance *interfaces.IDbInstance) *UserRepository {
	return &UserRepository{
		databaseInstance: dbInstance,
	}
}

func (u UserRepository) CreateUser(email string, username string, password string) (models.Users, error) {
	var user models.Users
	user.Email = email
	user.Username = username
	user.Password = password
	if res := u.databaseInstance.DataStore.Create(&user); res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u UserRepository) UpdateUser() {
}

func (u UserRepository) DeleteUser() {
}

func (u UserRepository) BulkDelete() {
}

func (u UserRepository) BulkUpdate() {
}

func (u UserRepository) FindOne() {
}

func (u UserRepository) FindMany() {
}
func (u UserRepository) FindAndCount() {
}
