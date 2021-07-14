package exception

import (
	"log"
)

const (
	INTERNAL_SOCKET_ERROR string = "INTERNAL_SOCKET_ERROR"
	INTERNAL_API_ERROR    string = "INTERNAL_API_ERROR"
	DATABASE_ERROR        string = "DATABASE_ERROR"
)

// Logging error messages to the server logs for debugging
func LogError(err error, msg string, t string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Type : " + t + "\nMessage : " + msg + " \nError : " + err.Error())
}
