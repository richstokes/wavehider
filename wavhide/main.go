package main

import (
	"crypto/sha512"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

func main() {
	var secretFile, audioFile, gap, pw = startup() // Initialize flags

	// sha512 hash of password from user input, make it 32 bytes
	sha512hash := sha512.New()
	sha512hash.Write([]byte(pw))
	// log.Println(string(sha512hash.Sum(nil)))

	// Encrypt file with given hashed password - must be 32 bytes
	encryptFile(secretFile, string(sha512hash.Sum(nil)[0:32]))

	// hide encrypted file into the audio file
	hideFile("secret.tmp", audioFile, gap) // input file set to the output of the encryption process

	// Delete .tmp file
	var errDel = os.Remove("secret.tmp")
	if errDel != nil {
		log.Println("Error deleting secret.tmp file generated during encryption process")
		log.Fatalln(errDel)
	}
}

func startup() (string, string, int, string) { // Processes flags and other options
	var secretFile string
	var audioFile string
	var gap int
	flag.StringVar(&secretFile, "secretFile", "SECRET_DOC.docx", "File you wish to hide")
	flag.StringVar(&audioFile, "audioFile", "song.wav", "Audio file you wish to hide a file within")
	flag.IntVar(&gap, "gap", 64, "Gap length to use - see README.md for more info")
	flag.Parse()

	// Check if files exist:
	_, err := os.Stat(secretFile)
	if os.IsNotExist(err) {
		flag.PrintDefaults()
		log.Fatalln(err)
	}

	_, err = os.Stat(audioFile)
	if os.IsNotExist(err) {
		flag.PrintDefaults()
		log.Fatalln(err)
	}

	// Check if would-be output file already exists
	fileExtension := strings.Split(audioFile, ".")[1]
	songOutFile := strings.Split(audioFile, ".")[0] + "-wavhidden." + fileExtension

	_, err = os.Stat(songOutFile)
	if !os.IsNotExist(err) {
		log.Println("Output file already exists: " + songOutFile)
		log.Fatalln("You probably want to move/rename this file.")
	}

	// Welcome & Prompt for password
	log.Println("Hello, friend.")
	log.Print("Enter a new encryption password for this file: ")

	pw, err := gopass.GetPasswdMasked()
	if err != nil {
		log.Fatalln(err)
	}

	return secretFile, audioFile, gap, string(pw)
}
