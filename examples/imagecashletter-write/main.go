// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/moov-io/imagecashletter"
)

const (
	citadelRoutingNumber    = "231380104"
	wellsFargoRoutingNumber = "121042882"
)

func main() {
	now := time.Now().UTC()
	imageBytes, err := os.ReadFile(filepath.Join("examples", "imagecashletter-write", "check_image.tiff"))
	if err != nil {
		log.Fatalf("could not open check image: %v\n", err)
	}
	imageDataLength := fmt.Sprintf("%d", len(imageBytes))

	header := imagecashletter.NewFileHeader()
	header.StandardLevel = "03"
	header.TestFileIndicator = "T"
	header.ImmediateDestination = citadelRoutingNumber
	header.ImmediateDestinationName = "Citadel"
	header.ImmediateOrigin = wellsFargoRoutingNumber
	header.ImmediateOriginName = "Wells Fargo"
	header.FileCreationDate = now
	header.FileCreationTime = now
	header.ResendIndicator = "N"

	header.FileIDModifier = "1"
	header.CompanionDocumentIndicator = "1"

	letterHeader := imagecashletter.NewCashLetterHeader()
	letterHeader.CollectionTypeIndicator = "01"
	letterHeader.DestinationRoutingNumber = citadelRoutingNumber
	letterHeader.ECEInstitutionRoutingNumber = wellsFargoRoutingNumber
	letterHeader.CashLetterBusinessDate = now
	letterHeader.CashLetterCreationDate = now
	letterHeader.CashLetterCreationTime = now
	letterHeader.RecordTypeIndicator = "I"
	letterHeader.DocumentationTypeIndicator = "G"
	letterHeader.CashLetterID = "74753"
	letterHeader.FedWorkType = "C"

	bundleHeader := imagecashletter.NewBundleHeader()
	bundleHeader.CollectionTypeIndicator = "01"
	bundleHeader.DestinationRoutingNumber = citadelRoutingNumber
	bundleHeader.ECEInstitutionRoutingNumber = wellsFargoRoutingNumber
	bundleHeader.BundleBusinessDate = now
	bundleHeader.BundleCreationDate = now
	bundleHeader.BundleID = "747531"
	bundleHeader.BundleSequenceNumber = "1"

	checkDetail := imagecashletter.NewCheckDetail()
	checkDetail.PayorBankRoutingNumber = "122000661"
	checkDetail.OnUs = "1211-1234-56789"
	checkDetail.ItemAmount = 1000
	checkDetail.EceInstitutionItemSequenceNumber = "29001104"
	checkDetail.DocumentationTypeIndicator = "G"
	checkDetail.ReturnAcceptanceIndicator = "0"
	checkDetail.MICRValidIndicator = 1
	checkDetail.BOFDIndicator = "Y"
	checkDetail.AddendumCount = 1
	checkDetail.CorrectionIndicator = 4
	checkDetail.ArchiveTypeIndicator = "F"

	addendumA := imagecashletter.NewCheckDetailAddendumA()
	addendumA.RecordNumber = 1
	addendumA.ReturnLocationRoutingNumber = wellsFargoRoutingNumber
	addendumA.BOFDEndorsementDate = now
	addendumA.BOFDItemSequenceNumber = "29001104"
	addendumA.BOFDBranchCode = "00029"
	addendumA.TruncationIndicator = "Y"
	addendumA.BOFDConversionIndicator = "2"
	addendumA.BOFDCorrectionIndicator = 0

	ivDetailFront := imagecashletter.NewImageViewDetail()
	ivDetailFront.ImageIndicator = 1
	ivDetailFront.ImageCreatorRoutingNumber = wellsFargoRoutingNumber
	ivDetailFront.ImageCreatorDate = now
	ivDetailFront.ImageViewFormatIndicator = "00"
	ivDetailFront.ImageViewCompressionAlgorithm = "00"
	ivDetailFront.ImageViewDataSize = imageDataLength
	ivDetailFront.ViewSideIndicator = 0
	ivDetailFront.ViewDescriptor = "00"
	ivDetailFront.DigitalSignatureIndicator = 0

	ivDataFront := imagecashletter.NewImageViewData()
	ivDataFront.EceInstitutionRoutingNumber = wellsFargoRoutingNumber
	ivDataFront.BundleBusinessDate = now
	ivDataFront.CycleNumber = "1"
	ivDataFront.EceInstitutionItemSequenceNumber = "29001104"
	ivDataFront.LengthImageData = imageDataLength
	ivDataFront.ImageData = imageBytes

	ivDetailBack := imagecashletter.NewImageViewDetail()
	ivDetailBack.ImageCreatorRoutingNumber = wellsFargoRoutingNumber
	ivDetailBack.ImageCreatorDate = now
	ivDetailBack.ImageViewFormatIndicator = "00"
	ivDetailBack.ImageViewCompressionAlgorithm = "00"
	ivDetailBack.ImageViewDataSize = imageDataLength
	ivDetailBack.ViewSideIndicator = 1
	ivDetailBack.ViewDescriptor = "00"
	ivDetailBack.DigitalSignatureIndicator = 0

	ivDataBack := imagecashletter.NewImageViewData()
	ivDataBack.EceInstitutionRoutingNumber = wellsFargoRoutingNumber
	ivDataBack.BundleBusinessDate = now
	ivDataBack.CycleNumber = "1"
	ivDataBack.EceInstitutionItemSequenceNumber = "29001104"
	ivDataBack.LengthImageData = imageDataLength
	ivDataBack.ImageData = imageBytes

	bundleControl := imagecashletter.NewBundleControl()
	bundleControl.BundleItemsCount = 1
	bundleControl.BundleTotalAmount = 1000
	bundleControl.MICRValidTotalAmount = 1000
	bundleControl.BundleImagesCount = 2

	letterControl := imagecashletter.NewCashLetterControl()
	letterControl.CashLetterBundleCount = 1
	letterControl.CashLetterItemsCount = 1
	letterControl.CashLetterTotalAmount = 1000
	letterControl.ECEInstitutionName = "Wells Fargo"
	letterControl.SettlementDate = now

	fileControl := imagecashletter.NewFileControl()
	fileControl.CashLetterCount = 1
	fileControl.TotalRecordCount = 12
	fileControl.TotalItemCount = 1
	fileControl.FileTotalAmount = 1000

	checkDetail.AddCheckDetailAddendumA(addendumA)
	checkDetail.AddImageViewDetail(ivDetailFront)
	checkDetail.AddImageViewData(ivDataFront)
	checkDetail.AddImageViewDetail(ivDetailBack)
	checkDetail.AddImageViewData(ivDataBack)

	bundle := imagecashletter.NewBundle(bundleHeader)
	bundle.AddCheckDetail(checkDetail)

	cashLetter := imagecashletter.NewCashLetter(letterHeader)
	cashLetter.AddBundle(bundle)
	if err = cashLetter.Create(); err != nil {
		log.Fatalf("could not create cash letter: %v\n", err)
	}

	file := imagecashletter.NewFile()
	file.SetHeader(header)
	file.AddCashLetter(cashLetter)

	if err = file.Create(); err != nil {
		log.Fatalf("Could not create File: %s\n", err)
	}
	if err = file.Validate(); err != nil {
		log.Fatalf("Could not validate File: %s\n", err)
	}

	opts := []imagecashletter.WriterOption{
		imagecashletter.WriteVariableLineLengthOption(),
		imagecashletter.WriteEbcdicEncodingOption(),
	}

	fd, err := os.Create(filepath.Join("examples", "imagecashletter-write", "iclFile.x937"))
	if err != nil {
		log.Fatalf("could not create output file: %v\n", err)
	}

	w := imagecashletter.NewWriter(fd, opts...)
	if err = w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
