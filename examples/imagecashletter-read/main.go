package main

import (
	"fmt"
	"github.com/moov-io/imagecashletter"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open(filepath.Join("examples", "imagecashletter-read", "iclFile.txt"))
	if err != nil {
		log.Fatalf("Could not open ICL File: %s\n", err)

	}
	defer f.Close()
	r := imagecashletter.NewReader(f)

	iclFile, err := r.Read()
	if err != nil {
		log.Fatalf("Could not read ICL File: %s\n", err)
	}
	// ensure we have a validated file structure
	if err = iclFile.Validate(); err != nil {
		log.Fatalf("Could not validate ICL File: %s\n", err)
	}

	fmt.Printf("CashLetterHeader: %v \n", iclFile.CashLetters[0].CashLetterHeader)
	fmt.Printf("CashLetterControl: %v \n", iclFile.CashLetters[0].CashLetterControl)
}
