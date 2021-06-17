package goapi

import (
	"context"
	"time"

	"github.com/CarlosTrejo2308/peopleApiResource/abort"
	"github.com/CarlosTrejo2308/peopleApiResource/db"
	"github.com/CarlosTrejo2308/peopleApiResource/people"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// getConnection generates a mongo concection and it returns it with
// a context of 5 seconds with its cancel function.
func getConnection() (*mongo.Collection, context.Context, context.CancelFunc) {
	//uri := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"

	// Create a mongo connection
	uri := db.GeneratePath()
	client := db.Connect(uri)

	// Create a context of 5 sec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Connecto to the database
	collection := client.Database("testing").Collection("personas")

	return collection, ctx, cancel
}

// GetAll connects to a mongo database,
// and returns all its contents with an
// array of Person
func GetAll() []people.Person {
	// Get collection and context
	collection, ctx, cancel := getConnection()
	defer cancel()

	// Find all the documents
	cur, err := collection.Find(ctx, bson.D{})
	abort.AbortOnError(err)

	// Close the cursos later
	defer cur.Close(ctx)

	// Get all the documents to a Person array
	responses := people.People{Persons: []people.Person{}}
	cur.All(ctx, &responses.Persons)

	return responses.Persons
}

// GetById recives an id of a document and
// it returns the content of the given document as an
// array of bson.M
func GetById(id int) []bson.M {
	// Get collection and context
	collection, ctx, cancel := getConnection()
	defer cancel()

	// Find the document with a id filter
	cur, err := collection.Find(ctx, bson.M{"id": id})
	abort.AbortOnError(err)

	// Close the cursor later
	defer cur.Close(ctx)

	// Get the document and stores it to response
	var response []bson.M
	cur.All(ctx, &response)

	return response
}
