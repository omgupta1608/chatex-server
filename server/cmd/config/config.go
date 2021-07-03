package config

import "flag"

func GetApiVersion() string {
	return "v1"
}

// get build flags from command line args
type BuildFlags struct {
	IsProduction bool
	Port         string
}

func GetBuildFlags() BuildFlags {
	buildFlags := BuildFlags{}

	// default value for flags
	const (
		isProdDefault = false
		portDefault   = "8080"
	)

	// register flags
	flag.BoolVar(&buildFlags.IsProduction, "p", isProdDefault, "Compile for production")
	flag.BoolVar(&buildFlags.IsProduction, "prod", isProdDefault, "Compile for production")

	flag.StringVar(&buildFlags.Port, "P", portDefault, "Set port")
	flag.StringVar(&buildFlags.Port, "port", portDefault, "Set port")

	flag.Parse()

	return buildFlags
}
