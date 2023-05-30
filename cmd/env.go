package main

import "os"

func GetEnv(key string, fallback string) string {

	env := os.Getenv(key)

	if len(env) == 0 {
		return fallback
	}

	return env
}
