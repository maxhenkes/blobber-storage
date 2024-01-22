package util

import (
	"crypto/md5"
	"encoding/hex"
)

func ComputeHashFromFile(file *[]byte) string {
	md5 := md5.Sum(*file)
	hash := hex.EncodeToString(md5[:])
	return hash
}
