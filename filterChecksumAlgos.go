package main

import (
	"crypto"
	"strings"
)

// Checksum algos
import (
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
)

// Man, why don't people allow their table to be exported...
var checksumLookupTable = map[ChecksumType]crypto.Hash{
	MD5:    crypto.MD5,
	SHA1:   crypto.SHA1,
	SHA256: crypto.SHA256,
	SHA512: crypto.SHA512,
}

var allChecksumAlgos []ChecksumType

func init() {
	allChecksumAlgos = make([]ChecksumType, len(checksumLookupTable))
	i := 0
	for k := range checksumLookupTable {
		allChecksumAlgos[i] = k
		i++
	}
}

func filterChecksumAlgos() {
	i := strings.Split(config.Checksums, ",")
	var j = map[ChecksumType]crypto.Hash{}
	for _, checksum := range i {
		csumtype := StringToChecksumType(checksum)
		if checksumLookupTable[csumtype].Available() == false {
			Error.Fatalf("Unsupported checksum algorithm: %v\n", checksum)
		}
		j[csumtype] = checksumLookupTable[csumtype]
	}
	checksumLookupTable = j
}

/*****************************************************************************/

type ChecksumType uint8

const (
	Unknown ChecksumType = iota // 0
	MD5                         // 1
	SHA1                        // 2
	SHA256                      // 3
	SHA512                      // 4
)

func StringToChecksumType(algo string) ChecksumType {
	switch strings.ToLower(algo) {
	case "md5", "md5sum":
		return MD5
	case "sha1", "sha1sum":
		return SHA1
	case "sha256", "sha256sum":
		return SHA256
	case "sha512", "sha512sum":
		return SHA512
	default:
		return Unknown
	}
}

func (checksumType ChecksumType) String() string {
	switch checksumType {
	case MD5:
		return "md5"
	case SHA1:
		return "sha1"
	case SHA256:
		return "sha256"
	case SHA512:
		return "sha512"
	default:
		return "unknown"
	}
}
