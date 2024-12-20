package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var mongoTimeout = 15 * time.Second

type Models struct {
	User User
}

type User struct {
	ID       int    `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		User: User{},
	}
}

func (u *User) CreateUser(user User) error {
	collection := client.Database("users").Collection("users")

	_, err := collection.InsertOne(context.TODO(), User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Println("Error while inserting user:", err)
		return err
	}
	return nil
}

func (u *User) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	collection := client.Database("users").Collection("users")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error while converting id to object id:", err)
		return nil, err
	}

	var user User
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&user)
	if err != nil {
		log.Println("Error while fetching user:", err)
		return nil, err
	}

	return &user, nil
}
