package main

import (
	"github.com/CarlosTrejo2308/peopleApiResource/db"
	gomongo "github.com/carlostrejo2308/gomongo/pkg"
)

func main() {
	// The name of the file where the people data is stored
	path := "./resources/people.xml"
	people := gomongo.ReadFile(path)

	// Get the DB path and connect to it
	uripath := db.GeneratePath()
	db := db.Connect(uripath)

	gomongo.InsertToBd(people, db)
}
