package secretstore

// Left these functions un-implemented since the main purpose of this
// exercise was to practice building CLI + web server and establish
// client-server communication. Would return to this some day though.

func Encrypt(plaintext, encodingKey string) (string, error) {
	return plaintext, nil
}

func Decrypt(cipher, encodingKey string) (string, error) {
	return cipher, nil
}
