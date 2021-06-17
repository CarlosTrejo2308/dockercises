package gomongo

import (
	"context"
	"fmt"
	"time"

	"github.com/CarlosTrejo2308/peopleApiResource/abort"
	"github.com/CarlosTrejo2308/peopleApiResource/people"
	"go.mongodb.org/mongo-driver/mongo"
)

// insertToBd recives a People struct and a mongo Client,
// and it inserts the People content to the mongo client in the
// personas collection of the testing database
func InsertToBd(people people.People, client *mongo.Client) {
	// Define a collection to use
	collection := client.Database("testing").Collection("personas")

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// For each person in the Persons array, it inserts it
	// to the mongo database
	for i := 0; i < len(people.Persons); i++ {
		// Inser to DB
		res, err := collection.InsertOne(ctx, people.Persons[i])
		abort.AbortOnError(err)

		// Print the id generated while it was inserted
		id := res.InsertedID
		fmt.Println(id)
	}

}
