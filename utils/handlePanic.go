package utils

import "log"

func HandlePanic() {
	if err := recover(); err != nil {
		log.Println("| Panic occured!")
		log.Println("|", err)
	}
}
