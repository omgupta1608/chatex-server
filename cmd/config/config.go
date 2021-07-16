package config

import (
	"errors"
	"flag"
	"os"
)

func GetApiVersion() string {
	return "v1"
}

func GetJwtSecret() []byte {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		panic(errors.New("JWT_SECRET must be set as environment var"))
	}

	return []byte(secretKey)
}

// get build flags from command line args
type BuildFlags struct {
	IsProduction bool
	Port         string
	RedisPath    string
	RedisAuth    string
}

var Flags BuildFlags

func init() {
	Flags = BuildFlags{}

	// default value for flags
	const (
		isProdDefault = false
		portDefault   = "8080"
	)

	// register flags
	flag.BoolVar(&Flags.IsProduction, "p", isProdDefault, "Compile for production")
	flag.BoolVar(&Flags.IsProduction, "prod", isProdDefault, "Compile for production")

	flag.StringVar(&Flags.Port, "P", portDefault, "Set port")
	flag.StringVar(&Flags.Port, "port", portDefault, "Set port")

	// redis flags
	flag.StringVar(&Flags.RedisPath, "redispath", "127.0.0.1:6379", "Path to redis cache")
	flag.StringVar(&Flags.RedisAuth, "redisauth", "", "auth to redis")

	flag.Parse()
}
