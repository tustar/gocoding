package main

import (
	"exercise/ch4/omdb"
	"os"
	"log"
)

func main() {
	result, err := omdb.SearchMovie(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	omdb.SavePoster(result.Poster)
}

