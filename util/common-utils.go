package util

import (
	"crypto/md5"
	"encoding/hex"
)

// ComputeHashFromFile returns a hash string for a given *[]byte
func ComputeHashFromFile(file *[]byte) string {
	md5 := md5.Sum(*file)
	hash := hex.EncodeToString(md5[:])
	return hash
}

// Map applies the function f to every element of ts and returns the new array
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
