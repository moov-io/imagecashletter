package imagecashletter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCashLetterPanics(t *testing.T) {
	var cl *CashLetter

	require.Nil(t, cl.GetBundles())
	require.Nil(t, cl.GetRoutingNumberSummary())
	require.Nil(t, cl.GetCreditItems())
}

// TestCashLetterNoBundle validates no Bundle when CashLetterHeader.RecordTypeIndicator = "N"
func TestCashLetterNoBundle(t *testing.T) {
	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.GetHeader().RecordTypeIndicator = "N"
	cl.AddBundle(bundle)
	err := cl.Create()
	var e *CashLetterError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordTypeIndicator", e.FieldName)
}

// TestCashLetterNoRoutingNumberSummary validates no Bundle when CashLetterHeader.CollectionTypeIndicator is not
// 00, 01, 02
func TestCashLetterRoutingNumberSummary(t *testing.T) {
	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.GetHeader().CollectionTypeIndicator = "03"
	cl.AddBundle(bundle)
	rns := mockRoutingNumberSummary()
	cl.AddRoutingNumberSummary(rns)
	err := cl.Create()
	var e *CashLetterError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CollectionTypeIndicator", e.FieldName)
}

func TestCashLetter_customSequenceNumber(t *testing.T) {
	// Create a forward check bundle
	checkBundleHeader := mockBundleHeader()
	checkBundleHeader.SetBundleSequenceNumber(564)
	checkBundle := NewBundle(checkBundleHeader)
	cd := mockCheckDetail()
	cd.SetEceInstitutionItemSequenceNumber(283)
	cd.AddendumCount = 4
	firstAddendumA := mockCheckDetailAddendumA()
	firstAddendumA.RecordNumber = 1
	cd.AddCheckDetailAddendumA(firstAddendumA)
	secondAddendumA := mockCheckDetailAddendumA()
	secondAddendumA.RecordNumber = 2
	cd.AddCheckDetailAddendumA(secondAddendumA)
	firstAddendumC := mockCheckDetailAddendumC()
	firstAddendumC.RecordNumber = 1
	cd.AddCheckDetailAddendumC(firstAddendumC)
	secondAddendumC := mockCheckDetailAddendumC()
	secondAddendumC.RecordNumber = 2
	cd.AddCheckDetailAddendumC(secondAddendumC)
	checkBundle.AddCheckDetail(cd)

	// Create a return bundle
	returnBundleHeader := mockBundleHeader()
	returnBundleHeader.BundleSequenceNumber = "" // test auto-increment behavior
	returnBundle := NewBundle(returnBundleHeader)
	rd := mockReturnDetail()
	rd.SetEceInstitutionItemSequenceNumber(4923)
	rd.AddendumCount = 2
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	returnBundle.AddReturnDetail(rd)

	clh := mockCashLetterHeader()
	cl := NewCashLetter(clh)
	cl.AddBundle(checkBundle)
	cl.AddBundle(returnBundle)
	require.NoError(t, cl.Create())

	require.Len(t, cl.Bundles, 2)
	require.Equal(t, "0564", cl.Bundles[0].BundleHeader.BundleSequenceNumber)
	require.Equal(t, "0565", cl.Bundles[1].BundleHeader.BundleSequenceNumber)

	require.Len(t, cl.Bundles[0].Checks, 1)
	wantCheckSeq := "000000000000283"
	require.Equal(t, wantCheckSeq, cl.Bundles[0].Checks[0].EceInstitutionItemSequenceNumber)

	// CheckDetailAddendumA
	require.Len(t, cl.Bundles[0].Checks[0].CheckDetailAddendumA, 2)
	require.Equal(t, 1, cl.Bundles[0].Checks[0].CheckDetailAddendumA[0].RecordNumber)
	require.Equal(t, 2, cl.Bundles[0].Checks[0].CheckDetailAddendumA[1].RecordNumber)
	require.Equal(t, wantCheckSeq, cl.Bundles[0].Checks[0].CheckDetailAddendumA[0].BOFDItemSequenceNumber)

	// CheckDetailAddendumC
	require.Len(t, cl.Bundles[0].Checks[0].CheckDetailAddendumC, 2)
	require.Equal(t, 1, cl.Bundles[0].Checks[0].CheckDetailAddendumC[0].RecordNumber)
	require.Equal(t, 2, cl.Bundles[0].Checks[0].CheckDetailAddendumC[1].RecordNumber)
	require.Equal(t, wantCheckSeq, cl.Bundles[0].Checks[0].CheckDetailAddendumC[0].EndorsingBankItemSequenceNumber)

	require.Len(t, cl.Bundles[1].Returns, 1)
	wantReturnSeq := "000000000004923"
	require.Equal(t, wantReturnSeq, cl.Bundles[1].Returns[0].EceInstitutionItemSequenceNumber)
	require.Equal(t, wantReturnSeq, cl.Bundles[1].Returns[0].ReturnDetailAddendumA[0].BOFDItemSequenceNumber)
	require.Equal(t, wantReturnSeq, cl.Bundles[1].Returns[0].ReturnDetailAddendumD[0].EndorsingBankItemSequenceNumber)
}
