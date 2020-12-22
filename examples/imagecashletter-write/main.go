// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/imagecashletter"
)

func main() {
	file := imagecashletter.NewFile()
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
	file.SetHeader(fh)

	// Create CheckDetail
	cd := imagecashletter.NewCheckDetail()
	cd.AuxiliaryOnUs = "123456789"
	cd.ExternalProcessingCode = ""
	cd.PayorBankRoutingNumber = "03130001"
	cd.PayorBankCheckDigit = "2"
	cd.OnUs = "5558881"
	cd.ItemAmount = 100000 // 1000.00
	cd.EceInstitutionItemSequenceNumber = "1              "
	cd.DocumentationTypeIndicator = "G"
	cd.ReturnAcceptanceIndicator = "D"
	cd.MICRValidIndicator = 1
	cd.BOFDIndicator = "Y"
	cd.AddendumCount = 3
	cd.CorrectionIndicator = 0
	cd.ArchiveTypeIndicator = "B"

	// create Check Detail AddendumA
	cdAddendumA := imagecashletter.NewCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 1
	cdAddendumA.ReturnLocationRoutingNumber = "121042882"
	cdAddendumA.BOFDEndorsementDate = time.Now()
	cdAddendumA.BOFDItemSequenceNumber = "1              "
	cdAddendumA.BOFDAccountNumber = "938383"
	cdAddendumA.BOFDBranchCode = "01"
	cdAddendumA.PayeeName = "Test Payee"
	cdAddendumA.TruncationIndicator = "Y"
	cdAddendumA.BOFDConversionIndicator = "1"
	cdAddendumA.BOFDCorrectionIndicator = 0
	cdAddendumA.UserField = ""

	// create Check Detail AddendumB
	cdAddendumB := imagecashletter.NewCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 1
	cdAddendumB.MicrofilmArchiveSequenceNumber = "1A             "
	cdAddendumB.LengthImageReferenceKey = "0034"
	cdAddendumB.ImageReferenceKey = "0"
	cdAddendumB.Description = "CD Addendum B"
	cdAddendumB.UserField = ""

	// create Check Detail AddendumC
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

	// create ImageViewDetail
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

	// create ImageViewData
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

	// create ImageViewAnalysis
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

	cd.AddCheckDetailAddendumA(cdAddendumA)
	cd.AddCheckDetailAddendumB(cdAddendumB)
	cd.AddCheckDetailAddendumC(cdAddendumC)

	cd.AddImageViewDetail(ivDetail)
	cd.AddImageViewData(ivData)
	cd.AddImageViewAnalysis(ivAnalysis)

	// create BundleHeader
	bh := imagecashletter.NewBundleHeader()
	bh.CollectionTypeIndicator = "01"
	bh.DestinationRoutingNumber = "231380104"
	bh.ECEInstitutionRoutingNumber = "121042882"
	bh.BundleBusinessDate = time.Now()
	bh.BundleCreationDate = time.Now()
	bh.BundleID = "9999"
	bh.BundleSequenceNumber = "1"
	bh.CycleNumber = "01"
	bh.UserField = ""

	bundle := imagecashletter.NewBundle(bh)
	bundle.AddCheckDetail(cd)

	// CheckDetail 2
	cdTwo := imagecashletter.NewCheckDetail()
	cdTwo.AuxiliaryOnUs = "123456789"
	cdTwo.ExternalProcessingCode = ""
	cdTwo.PayorBankRoutingNumber = "03130001"
	cdTwo.PayorBankCheckDigit = "2"
	cdTwo.OnUs = "5558881"
	cdTwo.ItemAmount = 100000 // 1000.00
	cdTwo.EceInstitutionItemSequenceNumber = "1              "
	cdTwo.DocumentationTypeIndicator = "G"
	cdTwo.ReturnAcceptanceIndicator = "D"
	cdTwo.MICRValidIndicator = 1
	cdTwo.BOFDIndicator = "Y"
	cdTwo.AddendumCount = 3
	cdTwo.CorrectionIndicator = 0
	cdTwo.ArchiveTypeIndicator = "B"

	cdTwo.AddCheckDetailAddendumA(cdAddendumA)
	cdTwo.AddCheckDetailAddendumB(cdAddendumB)
	cdTwo.AddCheckDetailAddendumC(cdAddendumC)
	cdTwo.AddImageViewDetail(ivDetail)
	cdTwo.AddImageViewData(ivData)
	cdTwo.AddImageViewAnalysis(ivAnalysis)
	bundle.AddCheckDetail(cdTwo)

	// Create ReturnDetail
	rd := imagecashletter.NewReturnDetail()
	rd.PayorBankRoutingNumber = "03130001"
	rd.PayorBankCheckDigit = "2"
	rd.OnUs = "5558881"
	rd.ItemAmount = 100000
	rd.ReturnReason = "A"
	rd.AddendumCount = 4
	rd.DocumentationTypeIndicator = "G"
	rd.ForwardBundleDate = time.Now()
	rd.EceInstitutionItemSequenceNumber = "1              "
	rd.ExternalProcessingCode = ""
	rd.ReturnNotificationIndicator = "2"
	rd.ArchiveTypeIndicator = "B"
	rd.TimesReturned = 0

	rdAddendumA := imagecashletter.NewReturnDetailAddendumA()
	rdAddendumA.RecordNumber = 1
	rdAddendumA.ReturnLocationRoutingNumber = "121042882"
	rdAddendumA.BOFDEndorsementDate = time.Now()
	rdAddendumA.BOFDItemSequenceNumber = "1              "
	rdAddendumA.BOFDAccountNumber = "938383"
	rdAddendumA.BOFDBranchCode = "01"
	rdAddendumA.PayeeName = "Test Payee"
	rdAddendumA.TruncationIndicator = "Y"
	rdAddendumA.BOFDConversionIndicator = "1"
	rdAddendumA.BOFDCorrectionIndicator = 0
	rdAddendumA.UserField = ""

	rdAddendumB := imagecashletter.NewReturnDetailAddendumB()
	rdAddendumB.PayorBankName = "Payor Bank Name"
	rdAddendumB.AuxiliaryOnUs = "123456789"
	rdAddendumB.PayorBankSequenceNumber = "1              "
	rdAddendumB.PayorBankBusinessDate = time.Now()
	rdAddendumB.PayorAccountName = "Payor Account Name"

	rdAddendumC := imagecashletter.NewReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKeyIndicator = 1
	rdAddendumC.MicrofilmArchiveSequenceNumber = "1A"
	rdAddendumC.LengthImageReferenceKey = "0034"
	rdAddendumC.ImageReferenceKey = "0"
	rdAddendumC.Description = "RD Addendum C"
	rdAddendumC.UserField = ""

	rdAddendumD := imagecashletter.NewReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 1
	rdAddendumD.EndorsingBankRoutingNumber = "121042882"
	rdAddendumD.BOFDEndorsementBusinessDate = time.Now()
	rdAddendumD.EndorsingBankItemSequenceNumber = "1              "
	rdAddendumD.TruncationIndicator = "Y"
	rdAddendumD.EndorsingBankConversionIndicator = "1"
	rdAddendumD.EndorsingBankCorrectionIndicator = 0
	rdAddendumD.ReturnReason = "A"
	rdAddendumD.UserField = ""
	rdAddendumD.EndorsingBankIdentifier = 0

	rdivDetail := imagecashletter.NewImageViewDetail()
	rdivDetail.ImageIndicator = 1
	rdivDetail.ImageCreatorRoutingNumber = "031300012"
	rdivDetail.ImageCreatorDate = time.Now()
	rdivDetail.ImageViewFormatIndicator = "00"
	rdivDetail.ImageViewCompressionAlgorithm = "00"
	// use of ivDetail.ImageViewDataSize is not recommended
	rdivDetail.ImageViewDataSize = "0000000"
	rdivDetail.ViewSideIndicator = 0
	rdivDetail.ViewDescriptor = "00"
	rdivDetail.DigitalSignatureIndicator = 0
	rdivDetail.DigitalSignatureMethod = "00"
	rdivDetail.SecurityKeySize = 00000
	rdivDetail.ProtectedDataStart = 0000000
	rdivDetail.ProtectedDataLength = 0000000
	rdivDetail.ImageRecreateIndicator = 0
	rdivDetail.UserField = ""
	rdivDetail.OverrideIndicator = "0"

	rdivData := imagecashletter.NewImageViewData()
	rdivData.EceInstitutionRoutingNumber = "121042882"
	rdivData.BundleBusinessDate = time.Now()
	rdivData.CycleNumber = "1"
	rdivData.EceInstitutionItemSequenceNumber = "1             "
	rdivData.SecurityOriginatorName = "Sec Orig Name"
	rdivData.SecurityAuthenticatorName = "Sec Auth Name"
	rdivData.SecurityKeyName = "SECURE"
	rdivData.ClippingOrigin = 0
	rdivData.ClippingCoordinateH1 = ""
	rdivData.ClippingCoordinateH2 = ""
	rdivData.ClippingCoordinateV1 = ""
	rdivData.ClippingCoordinateV2 = ""
	rdivData.LengthImageReferenceKey = "0000"
	rdivData.ImageReferenceKey = ""
	rdivData.LengthDigitalSignature = "0    "
	rdivData.DigitalSignature = []byte("")
	rdivData.LengthImageData = "0000001"
	rdivData.ImageData = []byte("")

	rdivAnalysis := imagecashletter.NewImageViewAnalysis()
	rdivAnalysis.GlobalImageQuality = 2
	rdivAnalysis.GlobalImageUsability = 2
	rdivAnalysis.ImagingBankSpecificTest = 0
	rdivAnalysis.PartialImage = 2
	rdivAnalysis.ExcessiveImageSkew = 2
	rdivAnalysis.PiggybackImage = 2
	rdivAnalysis.TooLightOrTooDark = 2
	rdivAnalysis.StreaksAndOrBands = 2
	rdivAnalysis.BelowMinimumImageSize = 2
	rdivAnalysis.ExceedsMaximumImageSize = 2
	rdivAnalysis.ImageEnabledPOD = 1
	rdivAnalysis.SourceDocumentBad = 0
	rdivAnalysis.DateUsability = 2
	rdivAnalysis.PayeeUsability = 2
	rdivAnalysis.ConvenienceAmountUsability = 2
	rdivAnalysis.AmountInWordsUsability = 2
	rdivAnalysis.SignatureUsability = 2
	rdivAnalysis.PayorNameAddressUsability = 2
	rdivAnalysis.MICRLineUsability = 2
	rdivAnalysis.MemoLineUsability = 2
	rdivAnalysis.PayorBankNameAddressUsability = 2
	rdivAnalysis.PayeeEndorsementUsability = 2
	rdivAnalysis.BOFDEndorsementUsability = 2
	rdivAnalysis.TransitEndorsementUsability = 2

	rd.AddReturnDetailAddendumA(rdAddendumA)
	rd.AddReturnDetailAddendumB(rdAddendumB)
	rd.AddReturnDetailAddendumC(rdAddendumC)
	rd.AddReturnDetailAddendumD(rdAddendumD)
	rd.AddImageViewDetail(rdivDetail)
	rd.AddImageViewData(rdivData)
	rd.AddImageViewAnalysis(rdivAnalysis)

	rdbh := imagecashletter.NewBundleHeader()
	rdbh.CollectionTypeIndicator = "01"
	rdbh.DestinationRoutingNumber = "231380104"
	rdbh.ECEInstitutionRoutingNumber = "121042882"
	rdbh.BundleBusinessDate = time.Now()
	rdbh.BundleCreationDate = time.Now()
	rdbh.BundleID = "9999"
	rdbh.BundleSequenceNumber = "1"
	rdbh.CycleNumber = "01"
	rdbh.UserField = ""

	returnBundle := imagecashletter.NewBundle(rdbh)
	returnBundle.BundleHeader.BundleSequenceNumber = "2"
	returnBundle.AddReturnDetail(rd)

	// Create ReturnDetail
	rdTwo := imagecashletter.NewReturnDetail()
	rdTwo.PayorBankRoutingNumber = "03130001"
	rdTwo.PayorBankCheckDigit = "2"
	rdTwo.OnUs = "5558881"
	rdTwo.ItemAmount = 100000
	rdTwo.ReturnReason = "A"
	rdTwo.AddendumCount = 4
	rdTwo.DocumentationTypeIndicator = "G"
	rdTwo.ForwardBundleDate = time.Now()
	rdTwo.EceInstitutionItemSequenceNumber = "1              "
	rdTwo.ExternalProcessingCode = ""
	rdTwo.ReturnNotificationIndicator = "2"
	rdTwo.ArchiveTypeIndicator = "B"
	rdTwo.TimesReturned = 0

	rdTwo.AddReturnDetailAddendumA(rdAddendumA)
	rdTwo.AddReturnDetailAddendumB(rdAddendumB)
	rdTwo.AddReturnDetailAddendumC(rdAddendumC)
	rdTwo.AddReturnDetailAddendumD(rdAddendumD)
	rdTwo.AddImageViewDetail(rdivDetail)
	rdTwo.AddImageViewData(rdivData)
	rdTwo.AddImageViewAnalysis(rdivAnalysis)
	returnBundle.AddReturnDetail(rdTwo)

	// Create CashLetter

	// create CashLetterHeader
	clh := imagecashletter.NewCashLetterHeader()
	clh.CollectionTypeIndicator = "01"
	clh.DestinationRoutingNumber = "231380104"
	clh.ECEInstitutionRoutingNumber = "121042882"
	clh.CashLetterBusinessDate = time.Now()
	clh.CashLetterCreationDate = time.Now()
	clh.CashLetterCreationTime = time.Now()
	clh.RecordTypeIndicator = "I"
	clh.DocumentationTypeIndicator = "G"
	clh.CashLetterID = "A1"
	clh.OriginatorContactName = "Contact Name"
	clh.OriginatorContactPhoneNumber = "5558675552"
	clh.FedWorkType = ""
	clh.ReturnsIndicator = ""
	clh.UserField = ""
	cl := imagecashletter.NewCashLetter(clh)
	cl.AddBundle(bundle)
	cl.AddBundle(returnBundle)
	cl.Create()
	file.AddCashLetter(cl)

	clTwo := imagecashletter.NewCashLetter(imagecashletter.NewCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)
	clTwo.AddBundle(returnBundle)
	clTwo.Create()
	file.AddCashLetter(clTwo)

	if err := file.Create(); err != nil {
		log.Fatalf("Could not create File: %s\n", err)
	}
	if err := file.Validate(); err != nil {
		log.Fatalf("Could not validate File: %s\n", err)
	}

	opts := []imagecashletter.WriterOption{
		imagecashletter.WriteVariableLineLengthOption(),
	}

	w := imagecashletter.NewWriter(os.Stdout, opts...)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

}
