package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"learn/read_write_json/utils"
	"os"
)

type Encripter struct {
	key string
}

func NewEncrypter() *Encripter {
	inKey := os.Getenv("KEY")

	return &Encripter{
		key: inKey,
	}
}

// !На данном этапе это лишь заглушка.
// Программу можно расширить, позволив пользователю ручной ввод ключа
func (e *Encripter) HasKey() bool {
	return e.key != ""
}

func (e *Encripter) DoEncrypt(bytesOfPlainText []byte) ([]byte, bool) {
	// generate a new aes cipher using our 32 byte long key
	aes, err := aes.NewCipher([]byte(e.key))
	if utils.HasError(err, "Encrypter/DoEncrypt/NewCipher") {
		return nil, false
	}

	// gcm (Galois/Counter Mode) - is a mode of operation
	// for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(aes)
	if utils.HasError(err, "Encrypter/DoEncrypt/NewGCM") {
		return nil, false
	}

	// We need a 12-byte nonce for GCM (modifiable
	// if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure random sequence
	_, errReadFull := io.ReadFull(rand.Reader, nonce)
	if utils.HasError(errReadFull, "Encrypter/DoEncrypt/ReadFull") {
		return nil, false
	}

	// Seal encrypts and authenticates bytesOfPlainText,
	// authenticates the additional data and appends the result to dst,
	// returning the updated slice. The nonce must be NonceSize() bytes long
	// and unique for all time, for a given key.
	return gcm.Seal(nonce, nonce, bytesOfPlainText, nil), true
}

func (e *Encripter) DoDecript(cipherText []byte) []byte {
	// generate a new aes cipher using our 32 byte long key
	aes, err := aes.NewCipher([]byte(e.key))
	if err != nil {
		panic(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// Since we know the decipherText is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, decipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, decipherText, nil)
	if err != nil {
		panic(err)
	}

	return plainText
}
