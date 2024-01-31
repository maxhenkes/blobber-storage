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

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
