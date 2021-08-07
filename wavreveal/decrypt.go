package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
	"os"
)

func decryptFile(secretFile string, pw string) {
	log.Println("Decrypting file..")

	key := []byte(pw)
	ciphertext, err := ioutil.ReadFile(secretFile)
	outputFile := secretFile // Overwrite the file with decrypted data when done

	if err != nil {
		log.Fatalln(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatalln(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Fatalln(err)
	}
	
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		os.Remove(outputFile)
		log.Println("Invalid password.")
		log.Fatalln(err)
	}

	// Write the decrypted file
	err = ioutil.WriteFile(outputFile, plaintext, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Decryption complete. File ready: %s", outputFile)
}
