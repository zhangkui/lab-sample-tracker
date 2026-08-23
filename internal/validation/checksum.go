package validation

import "hash/crc32"

func Checksum(b []byte) uint32           { return crc32.ChecksumIEEE(b) }
func Matches(b []byte, want uint32) bool { return Checksum(b) == want }
