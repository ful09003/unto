package internal

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func PersistOutput(p string, output Marshalable) error {
	if err := os.MkdirAll(p, 0700); err != nil {
		return err
	}
	out, err := yaml.Marshal(output)
	if err != nil {
		return err
	}
	fName := fmt.Sprintf("%s_%d", output.Command, output.InvokedAt.UTC().Unix())
	return os.WriteFile(path.Join(p, fName), out, 0600)
}
