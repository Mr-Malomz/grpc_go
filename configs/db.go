package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbHandler interface {
	GetUser(id string) (*User, error)
	CreateUser(user User) (*mongo.InsertOneResult, error)
	UpdateUser(id string, user User) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
	GetAllUsers() ([]*User, error)
}

type DB struct {
	client *mongo.Client
}

func NewDBHandler() dbHandler {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return &DB{client: client}
}

func colHelper(db *DB) *mongo.Collection {
	// return db.client.Database("projectMngt").Collection("user")
	return db.client.Database("rustDB").Collection("User")
}

func (db *DB) CreateUser(user User) (*mongo.InsertOneResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser := User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	res, err := col.InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (db *DB) GetUser(id string) (*User, error) {
	col := colHelper(db)
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := col.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (db *DB) UpdateUser(id string, user User) (*mongo.UpdateResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

	result, err := col.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return nil, err
	}

	return result, err
}

func (db *DB) DeleteUser(id string) (*mongo.DeleteResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := col.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return nil, err
	}

	return result, err
}

func (db *DB) GetAllUsers() ([]*User, error) {
	col := colHelper(db)
	var users []*User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := col.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for results.Next(ctx) {
		var singleUser *User
		if err = results.Decode(&singleUser); err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	return users, err
}
