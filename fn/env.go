package fn

import "os"

func EnvOrDefault(env, def string) string {
	value, found := os.LookupEnv(env)
	if found {
		return value
	}
	return def
}
