package main

import "log"

func panicErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}