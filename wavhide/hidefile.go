package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func hideFile(dataInFile string, songInFile string, gap int) {
	log.Println("Hiding data inside audio file..")
	song, err := ioutil.ReadFile(songInFile) // Load Audio file
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Audio file length: %d bytes", len(song))

	data, _ := ioutil.ReadFile(dataInFile) // Load in document

	startOffset := 1680          // starting place in Audio file, after header info
	currentOffset := startOffset // will increment this by the gap during the main loop

	// First 4 bytes of the output Audio file will contain the data length of the file to be hidden
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(len(data))) // Add document length to byte slice
	dataLength := binary.LittleEndian.Uint32(bs) // Get document length from byte slice
	log.Printf("Data file length: %d bytes", dataLength)

	// See if file is too big to be hidden within given audio file
	biggestFile := (uint32(len(song)) - uint32(startOffset)) / uint32(gap)
	log.Printf("INFO: Largest data file length we could hide in this song, using a gap of %d, would be %d bytes", gap, biggestFile)

	if dataLength > biggestFile {
		os.Remove("secret.tmp")
		log.Fatalln("Audio file not large enough to store given document, aborting. You could try a lower gap value.")
	}

	log.Println("Working..")

	// append bytes representing the data/document length to the start of the data/document file slice
	for i := range bs {
		data = append(data, 0)     // Create space in array
		copy(data[i+1:], data[i:]) // Shift contents over one
		data[i] = bs[i]            // Step 3 - Insert bytes from data/doc length into array
		// log.Println(data[i], bs[i])
	}

	// Take document bytes and insert them at gap intervals into the song

	// Test - overwrite bytes in Audio file - MUCH faster than dping the copy and shift
	// Also more realistic as the resultant Audio file size ends up being the same as the original this way
	for i := range data {
		// log.Printf("Updating offset %d with %X", currentOffset, data[i])
		// song = append(song, 0)                             // Step 1 - Add space to end of song data array
		// copy(song[currentOffset+1:], song[currentOffset:]) // Step 2 - Shift everything over one
		song[currentOffset] = data[i]       // Step 3 - Insert byte from file at position
		currentOffset = currentOffset + gap // Shift to the next gap
	}

	// Write the file to disk
	fileExtension := strings.Split(songInFile, ".")[1]
	songOutFile := strings.Split(songInFile, ".")[0] + "-wavhidden." + fileExtension

	err = ioutil.WriteFile(songOutFile, song, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Done. Data hidden inside %s!", songOutFile)
}
