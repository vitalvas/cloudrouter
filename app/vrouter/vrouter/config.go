package vrouter

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/vitalvas/cloudrouter/app/vrouter/config"
)

const configDir = "/var/lib/cloudrouter/config"

func LoadOrCreateConfig() (*config.Config, error) {
	conf := config.NewConfig()

	currentPath := path.Join(configDir, "current")
	if _, err := os.Stat(currentPath); !errors.Is(err, os.ErrNotExist) {
		data, err := os.ReadFile(currentPath)

		if err == nil {
			configFilePath := path.Join(configDir, strings.TrimSpace(string(data)))
			if _, err := os.Stat(configFilePath); !errors.Is(err, os.ErrNotExist) {
				data, err := os.ReadFile(configFilePath)
				if err != nil {
					return nil, err
				}

				conf, err = config.UnpackFile(data)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return conf, nil
}

func (this *VRouter) saveConfig() error {
	data, fileName, err := this.config.PackFile()
	if err != nil {
		return err
	}

	configFileName := path.Join(configDir, fileName)

	if _, err := os.Stat(configFileName); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(configFileName, data, 0640); err != nil {
			return err
		}

		return os.WriteFile(path.Join(configDir, "current"), []byte(fileName), 0640)
	}

	return nil
}
