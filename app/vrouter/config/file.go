package config

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

var ErrorConfigNotVerified = errors.New("config not verified")

func (this *Config) PackFile() ([]byte, string, error) {
	data, err := Marshal(this)
	if err != nil {
		return nil, "", err
	}

	hash := md5.Sum(data)
	dataHash := hex.EncodeToString(hash[:])

	content := pem.EncodeToMemory(&pem.Block{
		Type: "CONFIG",
		Headers: map[string]string{
			"xid": fmt.Sprintf("%d:%s", this.ID, dataHash),
		},
		Bytes: data,
	})

	return content, fmt.Sprintf("%d_%s.dat", this.ID, dataHash[:12]), nil
}

func UnpackFile(data []byte) (*Config, error) {
	block, _ := pem.Decode(data)

	xid, ok := block.Headers["xid"]
	if !ok {
		return nil, ErrorConfigNotVerified
	}

	xidSplit := strings.Split(xid, ":")
	if len(xidSplit) != 2 {
		return nil, ErrorConfigNotVerified
	}

	hash := md5.Sum(block.Bytes)

	if hex.EncodeToString(hash[:]) != xidSplit[1] {
		return nil, ErrorConfigNotVerified
	}

	conf, err := Unmarshal(block.Bytes)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
