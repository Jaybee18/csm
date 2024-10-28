package config

import "runtime"

type Alias struct {
	Name        string
	Command     string
	Description string
}

type Config struct {
	Shell   []string
	Aliases []*Alias
}

/*
	Creates a config struct pre-filled with defaults
*/
func GetDefaultConfig() *Config {
	config := &Config{}

	if runtime.GOOS == "windows" {
		config.Shell = []string{"cmd", "/c"}
	} else {
		config.Shell = []string{"/bin/sh", "-c"}
	}

	return config
}

/*
	Gets the default config, then gets the user-defined
	config file and overwrites the default config where
	applicable.
*/
func GetConfig() *Config {
	defaultConfig := GetDefaultConfig()

	// --- TODO ---

	return defaultConfig
}
