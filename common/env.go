package common

import "syscall"

func GetEnvString(key, fallback string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return fallback
}
