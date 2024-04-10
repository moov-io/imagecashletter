package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/moov-io/imagecashletter"
)

var (
	fPath      = flag.String("fPath", "BNK20181015-A.icl", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flagJson = flag.Bool("json", false, "Output ICL File in JSON to stdout")
)

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	path := *fPath

	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(path)

	if err != nil {
		log.Printf("ERROR: Can not open file: %s: \n", err)
		os.Exit(1)
	}

	r := imagecashletter.NewReader(f)
	ICLFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if ICLFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}

	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if ICLFile.Create(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	// Output file contents
	if *flagJson {
		if err := json.NewEncoder(os.Stdout).Encode(ICLFile); err != nil {
			fmt.Printf("ERROR: problem writing ICL File to stdout: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Total record count: %v \n", ICLFile.Control.TotalRecordCount)
		fmt.Printf("Cash Letter count: %v \n", ICLFile.Control.CashLetterCount)
		fmt.Printf("File total Item count: %v \n", ICLFile.Control.TotalItemCount)
		fmt.Printf("File total amount: %v \n", ICLFile.Control.FileTotalAmount)
	}
}
