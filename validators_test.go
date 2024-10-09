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
	require.True(t, validSizeInt(10))
	require.True(t, validSizeInt(1e7))

	require.False(t, validSizeInt(1e8+1))
	require.False(t, validSizeInt(1e9))
	require.False(t, validSizeInt(math.MaxInt))

	t.Run("converters", func(t *testing.T) {
		c := &converters{}

		// Do nothing if the request is too large
		require.Equal(t, "", c.alphaField("a", 1e9))
		require.Equal(t, "", c.numericField(7, 1e9))
		require.Equal(t, "", c.nbsmField("b", 1e9))
		require.Equal(t, "", c.stringField("c", 1e9))
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

	t.Run("int", func(t *testing.T) {
		require.False(t, validSizeInt(int(1e9)))
	})

	t.Run("uint", func(t *testing.T) {
		a := uint(100)
		b := uint(201)

		require.False(t, validSizeUint(a-b))
		require.True(t, validSizeUint(b-a))
	})
}
