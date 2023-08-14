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
