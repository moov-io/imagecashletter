package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/imagecashletter"
)

func main() {
	f, err := os.Open(filepath.Join("examples", "imagecashletter-read", "iclFile.x937"))
	if err != nil {
		log.Fatalf("Could not open ICL File: %s\n", err)

	}
	defer f.Close()
	r := imagecashletter.NewReader(f, imagecashletter.ReadVariableLineLengthOption(), imagecashletter.ReadEbcdicEncodingOption())

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
