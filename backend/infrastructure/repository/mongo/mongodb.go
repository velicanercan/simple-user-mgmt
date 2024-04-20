package mongo

import (
	"context"
	"fmt"
	"os"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoCTX = context.TODO()

// MongoDBDatabase struct holds the MongoDB database connection
type MongoDBDatabase struct {
	Client *mongo.Client
}

// NewMongoDBDatabase initializes and returns a MongoDB database connection
func NewMongoDBDatabase() *MongoDBDatabase {
	URI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(mongoCTX, options.Client().ApplyURI(URI))
	if err != nil {
		logger.Log("Failed to connect to database!", "panic", err.Error())
	}

	if err := client.Ping(mongoCTX, readpref.Primary()); err != nil {
		logger.Log("Failed to ping database!", "panic", err.Error())
	}
	return &MongoDBDatabase{Client: client}
}

// Close closes the MongoDB database connection
func (db *MongoDBDatabase) Close() {
	err := db.Client.Disconnect(mongoCTX)
	if err != nil {
		logger.Log("error", "Failed to close the database connection!", err.Error())
	}
}

// TODO: Check all methods below
// InsertUser inserts a user into the MongoDB database
func (db *MongoDBDatabase) InsertUser(user *domain.User) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	_, err := collection.InsertOne(mongoCTX, user)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the MongoDB database
func (db *MongoDBDatabase) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	cursor, err := collection.Find(mongoCTX, bson.D{{}})
	if err != nil {
		return users, err
	}
	defer cursor.Close(mongoCTX)
	err = cursor.All(mongoCTX, &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

// GetUser retrieves a user from the MongoDB database by ID
func (db *MongoDBDatabase) GetUser(id int) (*domain.User, error) {
	var user domain.User
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	err := collection.FindOne(mongoCTX, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// UpdateUser updates a user in the MongoDB database
func (db *MongoDBDatabase) UpdateUser(id int, user *domain.User) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	filter := bson.M{"id": id}
	update := bson.D{primitive.E{Key: "$set", Value: user}}
	result, err := collection.UpdateOne(mongoCTX, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no result found with id %d", id)
	}
	return nil
}

// DeleteUser deletes a user from the MongoDB database by ID
func (db *MongoDBDatabase) DeleteUser(id int) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(mongoCTX, filter)
	if err != nil {
		return err
	}
	return nil
}
