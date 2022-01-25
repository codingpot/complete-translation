package document

import (
	"fmt"
	"log"
	"testing"
)

func TestGetDocument(t *testing.T) {
	var credentials string
	var doc_id string

	doc, err := Get(credentials, doc_id)
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}
	fmt.Printf("The title of the doc is: %s\n", doc.Title)
}
