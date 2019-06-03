package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/moov-io/imagecashletter"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"time"
)

var (
	fPath      = flag.String("fPath", "", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	// output formats
	flagJson = flag.Bool("json", false, "Output file in json")
)

// main creates an X9 File with 2 CashLetters
// Each CashLetter contains 2 Bundles, with 100 CheckDetails
func main() {
	flag.Parse()

	filename := time.Now().UTC().Format("200601021504")
	if *flagJson {
		filename += ".json"
	} else {
		filename += ".x9"
	}

	path := filepath.Join(*fPath, filename)
	write(path)
}

func write(path string) {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// To create a file
	fh := imagecashletter.NewFileHeader()
	fh.StandardLevel = "35"
	fh.TestFileIndicator = "T"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now()
	fh.FileCreationTime = time.Now()
	fh.ResendIndicator = "N"
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	fh.FileIDModifier = ""
	fh.CountryCode = "US"
	fh.UserField = ""
	fh.CompanionDocumentIndicator = ""

	file := imagecashletter.NewFile()
	file.SetHeader(fh)

	// Create 4 CashLetters
	for i := 0; i < 4; i++ {

		count := strconv.Itoa(i + 1)
		// cashLetterHeader
		clh := imagecashletter.NewCashLetterHeader()
		clh.CollectionTypeIndicator = "01"
		clh.DestinationRoutingNumber = "231380104"
		clh.ECEInstitutionRoutingNumber = "121042882"
		clh.CashLetterBusinessDate = time.Now()
		clh.CashLetterCreationDate = time.Now()
		clh.CashLetterCreationTime = time.Now()
		clh.RecordTypeIndicator = "I"
		clh.DocumentationTypeIndicator = "G"
		clh.CashLetterID = "A" + count
		clh.OriginatorContactName = "Contact Name"
		clh.OriginatorContactPhoneNumber = "5558675552"
		clh.FedWorkType = ""
		clh.ReturnsIndicator = ""
		clh.UserField = ""
		cl := imagecashletter.NewCashLetter(clh)

		for y := 0; y < 2; y++ {
			{
				// Create Bundle
				bcount := strconv.Itoa(i + y)
				bh := imagecashletter.NewBundleHeader()
				bh.CollectionTypeIndicator = "01"
				bh.DestinationRoutingNumber = "231380104"
				bh.ECEInstitutionRoutingNumber = "121042882"
				bh.BundleBusinessDate = time.Now()
				bh.BundleCreationDate = time.Now()
				bh.BundleID = "B" + bcount
				bh.BundleSequenceNumber = bcount
				bh.CycleNumber = "01"
				bh.UserField = ""
				bundle := imagecashletter.NewBundle(bh)

				for z := 0; z < 100; z++ {
					cdCount := strconv.Itoa(z + 1)

					// Create Check Detail
					cd := imagecashletter.NewCheckDetail()
					cd.AuxiliaryOnUs = "123456789"
					cd.ExternalProcessingCode = ""
					cd.PayorBankRoutingNumber = "03130001"
					cd.PayorBankCheckDigit = "2"
					cd.OnUs = "5558881"
					cd.ItemAmount = 100000 // 1000.00
					cd.EceInstitutionItemSequenceNumber = cdCount
					cd.DocumentationTypeIndicator = "G"
					cd.ReturnAcceptanceIndicator = "D"
					cd.MICRValidIndicator = 1
					cd.BOFDIndicator = "Y"
					cd.AddendumCount = 3
					cd.CorrectionIndicator = 0
					cd.ArchiveTypeIndicator = "B"

					cdAddendumA := imagecashletter.NewCheckDetailAddendumA()
					cdAddendumA.RecordNumber = 1
					cdAddendumA.ReturnLocationRoutingNumber = "121042882"
					cdAddendumA.BOFDEndorsementDate = time.Now()
					cdAddendumA.BOFDItemSequenceNumber = cdCount
					cdAddendumA.BOFDAccountNumber = "938383"
					cdAddendumA.BOFDBranchCode = "01"
					cdAddendumA.PayeeName = "Test Payee"
					cdAddendumA.TruncationIndicator = "Y"
					cdAddendumA.BOFDConversionIndicator = "1"
					cdAddendumA.BOFDCorrectionIndicator = 0
					cdAddendumA.UserField = ""

					cdAddendumB := imagecashletter.NewCheckDetailAddendumB()
					cdAddendumB.ImageReferenceKeyIndicator = 1
					cdAddendumB.MicrofilmArchiveSequenceNumber = "1A             "
					cdAddendumB.LengthImageReferenceKey = "0034"
					cdAddendumB.ImageReferenceKey = "0"
					cdAddendumB.Description = "CD Addendum B"
					cdAddendumB.UserField = ""

					cdAddendumC := imagecashletter.NewCheckDetailAddendumC()
					cdAddendumC.RecordNumber = 1
					cdAddendumC.EndorsingBankRoutingNumber = "121042882"
					cdAddendumC.BOFDEndorsementBusinessDate = time.Now()
					cdAddendumC.EndorsingBankItemSequenceNumber = "1              "
					cdAddendumC.TruncationIndicator = "Y"
					cdAddendumC.EndorsingBankConversionIndicator = "1"
					cdAddendumC.EndorsingBankCorrectionIndicator = 0
					cdAddendumC.ReturnReason = "A"
					cdAddendumC.UserField = ""
					cdAddendumC.EndorsingBankIdentifier = 0

					ivDetail := imagecashletter.NewImageViewDetail()
					ivDetail.ImageIndicator = 1
					ivDetail.ImageCreatorRoutingNumber = "031300012"
					ivDetail.ImageCreatorDate = time.Now()
					ivDetail.ImageViewFormatIndicator = "00"
					ivDetail.ImageViewCompressionAlgorithm = "00"
					// use of ivDetail.ImageViewDataSize is not recommended
					ivDetail.ImageViewDataSize = "0000000"
					ivDetail.ViewSideIndicator = 0
					ivDetail.ViewDescriptor = "00"
					ivDetail.DigitalSignatureIndicator = 0
					ivDetail.DigitalSignatureMethod = "00"
					ivDetail.SecurityKeySize = 00000
					ivDetail.ProtectedDataStart = 0000000
					ivDetail.ProtectedDataLength = 0000000
					ivDetail.ImageRecreateIndicator = 0
					ivDetail.UserField = ""
					ivDetail.OverrideIndicator = "0"

					ivData := imagecashletter.NewImageViewData()
					ivData.EceInstitutionRoutingNumber = "121042882"
					ivData.BundleBusinessDate = time.Now()
					ivData.CycleNumber = "1"
					ivData.EceInstitutionItemSequenceNumber = "1             "
					ivData.SecurityOriginatorName = "Sec Orig Name"
					ivData.SecurityAuthenticatorName = "Sec Auth Name"
					ivData.SecurityKeyName = "SECURE"
					ivData.ClippingOrigin = 0
					ivData.ClippingCoordinateH1 = ""
					ivData.ClippingCoordinateH2 = ""
					ivData.ClippingCoordinateV1 = ""
					ivData.ClippingCoordinateV2 = ""
					ivData.LengthImageReferenceKey = "0000"
					ivData.ImageReferenceKey = ""
					ivData.LengthDigitalSignature = "0    "
					ivData.DigitalSignature = []byte("")
					ivData.LengthImageData = "0000001"
					ivData.ImageData = []byte("")

					ivAnalysis := imagecashletter.NewImageViewAnalysis()
					ivAnalysis.GlobalImageQuality = 2
					ivAnalysis.GlobalImageUsability = 2
					ivAnalysis.ImagingBankSpecificTest = 0
					ivAnalysis.PartialImage = 2
					ivAnalysis.ExcessiveImageSkew = 2
					ivAnalysis.PiggybackImage = 2
					ivAnalysis.TooLightOrTooDark = 2
					ivAnalysis.StreaksAndOrBands = 2
					ivAnalysis.BelowMinimumImageSize = 2
					ivAnalysis.ExceedsMaximumImageSize = 2
					ivAnalysis.ImageEnabledPOD = 1
					ivAnalysis.SourceDocumentBad = 0
					ivAnalysis.DateUsability = 2
					ivAnalysis.PayeeUsability = 2
					ivAnalysis.ConvenienceAmountUsability = 2
					ivAnalysis.AmountInWordsUsability = 2
					ivAnalysis.SignatureUsability = 2
					ivAnalysis.PayorNameAddressUsability = 2
					ivAnalysis.MICRLineUsability = 2
					ivAnalysis.MemoLineUsability = 2
					ivAnalysis.PayorBankNameAddressUsability = 2
					ivAnalysis.PayeeEndorsementUsability = 2
					ivAnalysis.BOFDEndorsementUsability = 2
					ivAnalysis.TransitEndorsementUsability = 2

					// Add CheckDetailAddendum* to CheckDetail
					cd.AddCheckDetailAddendumA(cdAddendumA)
					cd.AddCheckDetailAddendumB(cdAddendumB)
					cd.AddCheckDetailAddendumC(cdAddendumC)

					// Add ImageView* to CheckDetail
					cd.AddImageViewDetail(ivDetail)
					cd.AddImageViewData(ivData)
					cd.AddImageViewAnalysis(ivAnalysis)

					// Add CheckDetail to Bundle
					bundle.AddCheckDetail(cd)
				}
				cl.AddBundle(bundle)
			}
		}
		cl.Create()
		file.AddCashLetter(cl)
	}

	// ensure we have a validated file structure
	if file.Validate(); err != nil {
		fmt.Printf("Could not validate entire file: %v", err)
	}

	// Create the file
	if err := file.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// Write to a file
	if *flagJson {
		// Write in JSON format
		if err := json.NewEncoder(f).Encode(file); err != nil {
			fmt.Printf("%T: %s", err, err)
		}
	} else {
		// Write in X9 plain text format
		w := imagecashletter.NewWriter(f)
		if err := w.Write(file); err != nil {
			fmt.Printf("%T: %s", err, err)
		}
		w.Flush()
	}

	if err := f.Close(); err != nil {
		fmt.Println(err.Error())
	}

	/*	// We want to write the file to an io.Writer
		w := x9.NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/

	fmt.Printf("Wrote %s\n", path)

}
