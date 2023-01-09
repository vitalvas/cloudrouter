package loader

import (
	"bytes"
	"os"
	"os/exec"
)

func (lo *Loader) execute(appPath string) {
	cmd := exec.Command(appPath)

	cmd.Env = append(os.Environ(), lo.generateEnvData()...)

	var cmdErrors bytes.Buffer
	cmd.Stderr = &cmdErrors

	if err := cmd.Run(); err != nil {
		lo.log.Println(err)
	}

	if cmdErrors.Len() > 0 {
		lo.log.Println(cmdErrors.String())
	}
}
