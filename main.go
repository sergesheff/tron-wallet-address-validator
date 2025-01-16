package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcutil/base58"
	"strings"
)

// byte2hexStr converts a single byte to its two-character hexadecimal representation
func byte2hexStr(b byte) string {
	hexByteMap := "0123456789ABCDEF"
	return string(hexByteMap[b>>4]) + string(hexByteMap[b&0x0f])
}

// isHexChar checks if a byte represents a valid hexadecimal character (0-9, A-F, a-f)
func isHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

// hexChar2byte converts a hexadecimal character to its decimal value
func hexChar2byte(c byte) int {
	var d int

	if c >= 'A' && c <= 'F' {
		d = int(c - 'A' + 10)
	} else if c >= 'a' && c <= 'f' {
		d = int(c - 'a' + 10)
	} else if c >= '0' && c <= '9' {
		d = int(c - '0')
	}

	return d
}

// customSha256 computes the SHA256 hash of a hexadecimal string
// Returns the hash as a hexadecimal string
func customSha256(payload string) (string, error) {
	b, err := hex.DecodeString(payload)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write(b)
	hash := h.Sum(nil)

	return hex.EncodeToString(hash), nil
}

// byteArray2HexStr converts a byte array to its hexadecimal string representation
func byteArray2HexStr(payload []byte) string {
	var str string
	for _, b := range payload {
		str += byte2hexStr(b)
	}
	return str
}

// hexStr2byteArray converts a hexadecimal string to its byte array representation
// Ignores non-hexadecimal characters in the input string
func hexStr2byteArray(str string) []byte {
	byteArray := []byte{}
	d := 0
	j := 0

	for i := 0; i < len(str); i++ {
		c := str[i]
		if isHexChar(c) {
			d <<= 4
			d += hexChar2byte(c)
			j++
			if j%2 == 0 {
				byteArray = append(byteArray, byte(d))
				d = 0
			}
		}
	}
	return byteArray
}

// decodeBase58Address decodes a TRON address from its Base58 representation
// Performs checksum verification to ensure address integrity
// Returns nil if the address is invalid or checksum verification fails
func decodeBase58Address(address string) []byte {
	address = strings.TrimSpace(address)
	if len(address) != 34 {
		return nil
	}

	bAddress := base58.Decode(address)
	if len(bAddress) == 0 {
		return nil
	}
	checkSum := bAddress[len(bAddress)-4:]

	bAddress = bAddress[:len(bAddress)-4]

	hash0, err := customSha256(byteArray2HexStr(bAddress))
	if err != nil {
		return nil
	}

	sha, err := customSha256(hash0)
	if err != nil {
		return nil
	}

	hash1 := hexStr2byteArray(sha)
	checkSum1 := hash1[:4]

	for i := 0; i < 4; i++ {
		if checkSum1[i] != checkSum[i] {
			return nil
		}
	}

	return bAddress
}

// IsValidAddress checks if a given string is a valid TRON address
func IsValidAddress(address string) bool {
	bAddress := decodeBase58Address(address)

	if bAddress == nil || len(bAddress) != 21 {
		return false
	}

	return bAddress[0] == 0x41
}
