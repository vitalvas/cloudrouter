package sysctl

func Get(key string) (string, error) {
	return readFile(pathFromKey(key))
}

func Set(key, value string) error {
	return writeFile(pathFromKey(key), value)
}
