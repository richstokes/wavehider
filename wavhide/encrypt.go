package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

const outputFile = "secret.tmp"

func encryptFile(secretFile string, pw string) {

	log.Println("Starting encryption..")
	text, err := ioutil.ReadFile(secretFile)
	key := []byte(pw)

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		log.Fatalln(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		log.Fatalln(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	// fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile(outputFile, gcm.Seal(nonce, nonce, text, nil), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("File encrypted OK!")
}
