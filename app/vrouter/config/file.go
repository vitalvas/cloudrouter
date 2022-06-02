package config

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
)

func PackFile(id int64, data []byte) {
	hash := md5.Sum(data)

	content := pem.EncodeToMemory(&pem.Block{
		Type: "CONFIG",
		Headers: map[string]string{
			"xid": fmt.Sprintf("%d:%s", id, base64.RawURLEncoding.EncodeToString(hash[:])),
		},
		Bytes: data,
	})

	print(strings.TrimSpace(string(content)))
}
