package exception

import (
	"log"
)

var (
	INTERNAL_SOCKET_ERROR = "INTERNAL_SOCKET_ERROR"
	INTERNAL_API_ERROR    = "INTERNAL_API_ERROR"
	DATABASE_ERROR        = "DATABASE_ERROR"
)

// Logging error messages to the server logs for debugging
func LogError(err error, msg string, t string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Message : " + msg + " \nError : " + err.Error())
}
