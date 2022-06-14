package lib

import "os"

func IsRailwayInstance() bool {
	return os.Getenv("RAILWAY_ENVIRONMENT") != ""
}

func GetRailwayEnvironment() string {
	return os.Getenv("RAILWAY_ENVIRONMENT")
}
