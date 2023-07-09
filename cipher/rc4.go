package cipher

import "crypto/rc4"

// Rc4 Failed, pending refactoring
func Rc4(text, key string) (string, error) {
	var desKey []byte
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	cipher.XORKeyStream(desKey, []byte(text))
	return string(desKey), nil
}
