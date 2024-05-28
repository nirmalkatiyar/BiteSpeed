package config

import "os"

// GetPort ..... server PORT
func GetPort() string {
	if os.Getenv("PORT") == "" {
		SetPort("8080")
	}
	return os.Getenv("PORT")
}

func SetPort(port string) {
	// default port set to 8080
	if port == "" {
		os.Setenv("PORT", "8080")
	}
	os.Setenv("PORT", port)
}
