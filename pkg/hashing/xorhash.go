package hashing

// XORHash applies XOR to the given bytearray and then hashes it.
func XORHash(s []byte) uint64 {
	k := []byte{'h', 'a', 'd', 'e', 's'} // randomize?
	for i := 0; i < len(s); i++ {
		s[i] ^= k[i%len(k)]
	}
	return djb2(s)
}

// simple djb2 hashing algorithm implementation.
func djb2(s []byte) uint64 {
	var hash uint64 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + uint64(c)
	}
	return hash
}
