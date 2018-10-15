package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/moov-io/x9"
	"log"
	"os"
	"runtime/pprof"
)

var (
	fPath      = flag.String("fPath", "BNK20181015-A.x9", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flagJson = flag.Bool("json", false, "Output X9 File in JSON to stdout")
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
		log.Panicf("Can not open file: %s: \n", err)
	}

	r := x9.NewReader(f)
	x9File, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if x9File.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}

	// If you trust the file but it's formating is off building will probably resolve the malformed file.
	if x9File.Create(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	// Output file contents
	if *flagJson {
		if err := json.NewEncoder(os.Stdout).Encode(x9File); err != nil {
			fmt.Printf("ERROR: problem writing X9 File to stdout: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Total record count: %v \n", x9File.Control.TotalRecordCount)
		fmt.Printf("Cash Letter count: %v \n", x9File.Control.CashLetterCount)
		fmt.Printf("File total Item count: %v \n", x9File.Control.TotalItemCount)
		fmt.Printf("File total amount: %v \n", x9File.Control.FileTotalAmount)
	}
}
