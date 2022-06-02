package config

import (
	"bytes"
	"compress/flate"
	"encoding/gob"
)

func Marshal(conf Config) ([]byte, error) {
	var data bytes.Buffer

	if err := gob.NewEncoder(&data).Encode(conf); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	zw, err := flate.NewWriter(&buf, flate.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := zw.Write(data.Bytes()); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
