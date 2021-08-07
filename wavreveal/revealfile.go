package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
)

func revealFile(songInFile string, gap int, outputFile string) {
	log.Println("Revealing data hidden in audio file..")
	song, err := ioutil.ReadFile(songInFile)
	if err != nil {
		log.Fatalln(err)
	}
	startOffset := 1680 // starting place in song file, after header info
	currentOffset := startOffset
	count := 1

	bs := make([]byte, 4) // array for storting the data length info

	for i := range bs { // Get the 4-byte length data from the source song file
		bs[i] = song[currentOffset]
		// log.Println(currentOffset, bs[i])
		currentOffset = currentOffset + gap
	}
	// fmt.Println(bs)
	dataLength := binary.LittleEndian.Uint32(bs)

	var decodedBytes []byte // array to store bytes to make up output file

	log.Printf("Data length detected as %d, decoding file %s now..", dataLength, songInFile)

	for count <= int(dataLength) {
		decoded := song[currentOffset : currentOffset+1]
		// log.Printf("Byte: %d \t String: %s \t Offset: %d", decoded, string(decoded), currentOffset)
		currentOffset = currentOffset + gap // Shift to the next gap
		count++
		decodedBytes = append(decodedBytes, decoded[0])
	}

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile(outputFile, decodedBytes, 0644)
	// handle this error
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Finished extracting encrypted data from song")
}
