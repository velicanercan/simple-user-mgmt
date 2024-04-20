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

// MongoDBDatabase struct holds the MongoDB database connection
type MongoDBDatabase struct {
	Client *mongo.Client
}

// NewMongoDBDatabase initializes and returns a MongoDB database connection
func NewMongoDBDatabase(ctx context.Context) *MongoDBDatabase {
	URI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		logger.Log("Failed to connect to database!", "panic", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Log("Failed to ping database!", "panic", err.Error())
	}
	return &MongoDBDatabase{Client: client}
}

// Close closes the MongoDB database connection
func (db *MongoDBDatabase) Close(ctx context.Context) {
	err := db.Client.Disconnect(ctx)
	if err != nil {
		logger.Log("error", "Failed to close the database connection!", err.Error())
	}
}

// TODO: Check all methods below
// InsertUser inserts a user into the MongoDB database
func (db *MongoDBDatabase) InsertUser(ctx context.Context, user *domain.User) error {
	collection := db.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

// GetAllUsers retrieves all users from the MongoDB database
func (db *MongoDBDatabase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	collection := db.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	return users, err
}

// GetUser retrieves a user from the MongoDB database by ID
func (db *MongoDBDatabase) GetUser(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	collection := db.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	return &user, err
}

// UpdateUser updates a user in the MongoDB database
func (db *MongoDBDatabase) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	collection := db.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	filter := bson.M{"id": id}
	update := bson.D{primitive.E{Key: "$set", Value: user}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no result found with id %d", id)
	}
	return nil
}

// DeleteUser deletes a user from the MongoDB database by ID
func (db *MongoDBDatabase) DeleteUser(ctx context.Context, id int) error {
	collection := db.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
