package repository

import (
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct{
	DB *mongo.Collection
}

func NewUserRepository(DB *mongo.Collection) *userRepository{
	return &userRepository{DB}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User,  error){
	

	return interface{}, nil
}