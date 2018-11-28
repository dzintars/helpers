package os

import "os"

// Getenv returns string of argument key or set up new envirionment variable and returns its value.
func Getenv(key, fallback string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		os.Setenv(key, fallback)
		return os.Getenv(key)
	}
	return v
}
