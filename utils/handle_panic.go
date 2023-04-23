package utils

import (
	"log"
)

// RecoverAfterPanic recovers from a panic and logs the error
func RecoverAfterPanic(i any) {
	if r := recover(); r != nil {
		log.Printf("Recovered in %s", GetFuncName(i))
	}
}
