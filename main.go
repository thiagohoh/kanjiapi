package main

import (
	"fmt"

	"github.com/kyary/kanjiapi/database"
	"github.com/kyary/kanjiapi/models"
	"github.com/kyary/kanjiapi/routes"
)

func main() {
	fmt.Println("OP")

	k := models.Kanji{
		Kanji:   "kanji",
		Meaning: "meaning",
		Reading: "reading",
	}

	fmt.Println(k)

	database.Connection()
	routes.HandleRequests()

}
