package loader

import (
	"os"
	"path/filepath"
)

func (lo *Loader) execute(appPath string) {
	child, err := os.StartProcess(
		appPath, nil,
		&os.ProcAttr{
			Env: append(os.Environ(), lo.generateEnvData()...),
		},
	)
	if err != nil {
		lo.log.Fatal(err)
	}

	name := filepath.Base(appPath)

	lo.apps[name] = child
}
