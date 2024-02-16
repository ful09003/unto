package internal

import (
	"strings"
)

func GatherEnvVarVals(inEnv []string, k []string) map[string]string {
	ret := make(map[string]string)
	for _, desiredKey := range k {
		k, v := gatherVarVal(inEnv, desiredKey)
		ret[k] = v
	}

	return ret
}

func gatherVarVal(inEnv []string, k string) (string, string) {
	for _, envVarRaw := range inEnv {
		before, after, found := strings.Cut(envVarRaw, "=")
		if found && before == k {
			return k, after
		}
	}
	return k, ""
}
