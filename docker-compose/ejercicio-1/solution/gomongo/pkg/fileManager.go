package gomongo

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/CarlosTrejo2308/peopleApiResource/abort"
	"github.com/CarlosTrejo2308/peopleApiResource/people"
)

// readFile recives a a path of an xml file where
// persons data are stored, and it returns a
// People struct containing the persons information.
func ReadFile(path string) people.People {
	// Open file
	xmlFile, err := os.Open(path)
	abort.AbortOnError(err)

	// Close file at the end
	defer xmlFile.Close()

	// Read the file
	byteValue, err := ioutil.ReadAll(xmlFile)
	abort.AbortOnError(err)

	// Store it to a People struct
	var people people.People
	xml.Unmarshal(byteValue, &people)

	return people
}
