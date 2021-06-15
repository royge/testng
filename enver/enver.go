package enver

import (
	"os"
	"testing"
)

// CheckEnv checks if list environment variables envs are defined and fail the
// test if not.
func CheckEnv(t *testing.T, envs ...string) {
	t.Helper()

	for _, env := range envs {
		env := env // shadow copy

		if v := os.Getenv(env); v == "" {
			t.Fatalf("environment variable `%s` is not defined", env)
		}
	}
}
