// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidSize(t *testing.T) {
	require.True(t, validSize(10))
	require.True(t, validSize(1e7))

	require.False(t, validSize(1e8+1))
	require.False(t, validSize(1e9))
	require.False(t, validSize(math.MaxInt))

	t.Run("converters", func(t *testing.T) {
		c := &converters{}

		// Do nothing if the request is too large
		require.Equal(t, "a", c.alphaField("a", 1e9))
		require.Equal(t, "7", c.numericField(7, 1e9))
		require.Equal(t, "b", c.nbsmField("b", 1e9))
		require.Equal(t, "c", c.stringField("c", 1e9))
	})

	t.Run("don't grow", func(t *testing.T) {
		cdAddendumB := &CheckDetailAddendumB{}
		cdAddendumB.LengthImageReferenceKey = fmt.Sprintf("%0.0f", 1e9)
		require.Equal(t, "0               1000                        ", cdAddendumB.String())

		ivData := &ImageViewData{}
		ivData.LengthImageReferenceKey = fmt.Sprintf("%0.0f", 1e9)
		ivData.LengthDigitalSignature = fmt.Sprintf("%0.0f", 1e9)
		ivData.LengthImageData = fmt.Sprintf("%0.0f", 1e9)
		expected := "00000000000010101                                                                 0                1000100001000000"
		require.Equal(t, expected, ivData.String())

		rdAddendumC := &ReturnDetailAddendumC{}
		rdAddendumC.LengthImageReferenceKey = fmt.Sprintf("%0.0f", 1e9)
		expected = "0               1000                        "
		require.Equal(t, expected, rdAddendumC.String())

		ug := &UserGeneral{}
		ug.LengthUserData = fmt.Sprintf("%0.0f", 1e9)
		expected = "0                                   1000000"
		require.Equal(t, expected, ug.String())
	})
}
