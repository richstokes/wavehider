package main

import (
	"crypto/sha512"
	"flag"
	"log"
	"os"

	"github.com/howeyc/gopass"
)

func main() {
	var outputFile, audioFile, gap, pw = startup() // Initialize flags

	// sha512 hash of password from user input, make it 32 bytes
	sha512hash := sha512.New()
	sha512hash.Write([]byte(pw))
	// log.Println(string(sha512hash.Sum(nil)))

	// Reveal data from song file
	revealFile(audioFile, gap, outputFile)

	// Then decrypt
	decryptFile(outputFile, string(sha512hash.Sum(nil)[0:32]))
}

func startup() (string, string, int, string) { // Processes flags and other options
	var outputFile string
	var audioFile string
	var gap int
	flag.StringVar(&outputFile, "outputFile", "SECRET_DOC.docx", "File name for the revealed output")
	flag.StringVar(&audioFile, "audioFile", "song-hidden.wav", "Audio file containing the hidden data you wish to reveal")
	flag.IntVar(&gap, "gap", 64, "Gap length to use - see README.md for more info")
	flag.Parse()

	// Check if files exist:
	_, err := os.Stat(audioFile)
	if os.IsNotExist(err) {
		flag.PrintDefaults()
		log.Fatalln(err)
	}
	
	_, err = os.Stat(outputFile)
	if !os.IsNotExist(err) {
		log.Println("Output file already exists: " + outputFile)
		log.Fatalln("You probably want to move/rename this file.")
	}

	// Welcome & Prompt for password
	log.Println("Hello, friend.")
	log.Print("Enter encryption password: ")

	pw, err := gopass.GetPasswdMasked()
	if err != nil {
		log.Fatalln(err)
	}

	return outputFile, audioFile, gap, string(pw)
}
