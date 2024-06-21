package cryptomath

import (
	"crypto/rand"
	"encoding/binary"
)

// Generates a cryptographically secure random int8.
func CryptoRandomInt8() (int8, error) {
	var b [1]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

// Generates a cryptographically secure random int16.
func CryptoRandomInt16() (int16, error) {
	var b [2]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(b[:])), nil
}

// Generates a cryptographically secure random int32.
func CryptoRandomInt32() (int32, error) {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(b[:])), nil
}

// Generates a cryptographically secure random int64.
func CryptoRandomInt64() (int64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

// Generates a cryptographically secure random uint8.
func CryptoRandomUint8() (uint8, error) {
	var b [1]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return b[0], nil
}

// Generates a cryptographically secure random uint16.
func cryptoRandomUint16() (uint16, error) {
	var b [2]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b[:]), nil
}

// Generates a cryptographically secure random uint32.
func cryptoRandomUint32() (uint32, error) {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b[:]), nil
}

// Generates a cryptographically secure random uint64.
func cryptoRandomUint64() (uint64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

// Generates a cryptographically secure random float32 between 0 and 1.
func cryptoRandomFloat32() (float32, error) {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return float32(binary.LittleEndian.Uint32(b[:])) / (1 << 32), nil
}

// Generates a cryptographically secure random float64 between 0 and 1.
func CryptoRandomFloat64() (float64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return float64(binary.LittleEndian.Uint64(b[:])) / (1 << 64), nil
}
