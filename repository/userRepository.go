package repository

import (
	"context"
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByEmail(ctx context.Context,email string) (models.User, error)
}

type userRepository struct{
	DB *mongo.Collection
}

func NewUserRepository(DB *mongo.Collection) *userRepository{
	return &userRepository{DB}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (models.User,  error){

	var user models.User

	err := r.DB.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil{
		return user, err
	}

	return user, nil
}