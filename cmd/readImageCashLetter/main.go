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

	flagJson           = flag.Bool("json", false, "Output ICL File in JSON to stdout")
	flagSkipValidation = flag.Bool("skip-validation", false, "Skip validation checks for non-compliant or archived files")
	flagEbcdic         = flag.Bool("ebcdic", false, "Decode the file as EBCDIC (default is ASCII)")
	flagVariableLength = flag.Bool("variable-length", false, "Read records using Inserted Length Field prefixes")
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

	if err := readFile(*fPath); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func readerOptions() []imagecashletter.ReaderOption {
	var opts []imagecashletter.ReaderOption
	if *flagVariableLength {
		opts = append(opts, imagecashletter.ReadVariableLineLengthOption())
	}
	if *flagEbcdic {
		opts = append(opts, imagecashletter.ReadEbcdicEncodingOption())
	}
	if *flagSkipValidation {
		// Use SkipAll for the broad "skip validation" use case documented on the flag.
		// (SkipCountValidation is a narrower option for specific count checks.)
		opts = append(opts, imagecashletter.ReadValidateOpts(&imagecashletter.ValidateOpts{SkipAll: true}))
	}
	return opts
}

func readFile(path string) error {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("ERROR: Can not open file: %s", err)
	}
	defer f.Close()

	r := imagecashletter.NewReader(f, readerOptions()...)
	ICLFile, err := r.Read()
	if err != nil {
		return fmt.Errorf("Issue reading file: %+v", err)
	}

	if *flagSkipValidation {
		ICLFile.SetValidation(&imagecashletter.ValidateOpts{SkipAll: true})
	}

	// ensure we have a validated file structure
	if err := ICLFile.Validate(); err != nil {
		return fmt.Errorf("Could not validate entire read file: %v", err)
	}

	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := ICLFile.Create(); err != nil {
		return fmt.Errorf("Could not build file with read properties: %v", err)
	}

	// Output file contents
	if *flagJson {
		if err := json.NewEncoder(os.Stdout).Encode(ICLFile); err != nil {
			return fmt.Errorf("ERROR: problem writing ICL File to stdout: %v", err)
		}
		return nil
	}

	fmt.Printf("Total record count: %v \n", ICLFile.Control.TotalRecordCount)
	fmt.Printf("Cash Letter count: %v \n", ICLFile.Control.CashLetterCount)
	fmt.Printf("File total Item count: %v \n", ICLFile.Control.TotalItemCount)
	fmt.Printf("File total amount: %v \n", ICLFile.Control.FileTotalAmount)
	return nil
}
