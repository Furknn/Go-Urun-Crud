package helpers

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
