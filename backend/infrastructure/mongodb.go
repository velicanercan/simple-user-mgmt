package infrastructure

import (
	"context"
	"os"

	"github.com/velicanercan/simple-user-mgmt/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBDatabase struct holds the MongoDB database connection
type MongoDBDatabase struct {
	Client *mongo.Client
}

// NewMongoDBDatabase initializes and returns a MongoDB database connection
func NewMongoDBDatabase() *MongoDBDatabase {
	URI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))
	if err != nil {
		panic("Failed to connect to database!")
	}

	return &MongoDBDatabase{Client: client}
}

// Close closes the MongoDB database connection
func (db *MongoDBDatabase) Close() {
	err := db.Client.Disconnect(context.Background())
	if err != nil {
		panic("Failed to close database connection!")
	}
}

// TODO: Check all methods below
// InsertUser inserts a user into the MongoDB database
func (db *MongoDBDatabase) InsertUser(user *models.User) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the MongoDB database
func (db *MongoDBDatabase) GetAllUsers() ([]models.User, error) {
	var users []models.User
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return users, err
	}
	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

// GetUser retrieves a user from the MongoDB database by ID
func (db *MongoDBDatabase) GetUser(id int) (*models.User, error) {
	var user models.User
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	err := collection.FindOne(context.Background(), models.User{ID: id}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// UpdateUser updates a user in the MongoDB database
func (db *MongoDBDatabase) UpdateUser(user *models.User) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	_, err := collection.UpdateOne(context.Background(), models.User{ID: user.ID}, user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the MongoDB database by ID
func (db *MongoDBDatabase) DeleteUser(id int) error {
	collection := db.Client.Database("simple-user-mgmt").Collection("users")
	_, err := collection.DeleteOne(context.Background(), models.User{ID: id})
	if err != nil {
		return err
	}
	return nil
}
