package sysctl

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const defaultPath = "/proc/sys/"

func pathFromKey(key string) string {
	return filepath.Join(defaultPath, strings.Replace(key, ".", "/", -1))
}

func readFile(path string) (string, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return "", os.ErrNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func writeFile(path, value string) error {
	return os.WriteFile(path, []byte(value), 0644)
}
