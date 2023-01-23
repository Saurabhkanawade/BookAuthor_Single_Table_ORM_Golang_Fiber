package helper

import "log"

func HandleNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
