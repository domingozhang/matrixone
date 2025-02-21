// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"bytes"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

var area = make([]byte, 0, 10000)

func TestRangeShuffle(t *testing.T) {
	require.Equal(t, GetRangeShuffleIndexUnsigned(0, 1000000, 299999, 10), uint64(2))
	require.Equal(t, GetRangeShuffleIndexUnsigned(0, 1000000, 888, 10), uint64(0))
	require.Equal(t, GetRangeShuffleIndexUnsigned(0, 1000000, 100000000, 10), uint64(9))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 299999, 10), uint64(2))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 888, 10), uint64(0))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 100000000, 10), uint64(9))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, -2, 10), uint64(0))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 999000, 10), uint64(9))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 99999, 10), uint64(0))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 100000, 10), uint64(1))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 100001, 10), uint64(1))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 199999, 10), uint64(1))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 200000, 10), uint64(2))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 999999, 10), uint64(9))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 899999, 10), uint64(8))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 900000, 10), uint64(9))
	require.Equal(t, GetRangeShuffleIndexSigned(0, 1000000, 1000000, 10), uint64(9))
}

func buildVarlenaFromByteSlice(bs []byte) *types.Varlena {
	var v types.Varlena
	vlen := len(bs)
	if vlen <= types.VarlenaInlineSize {
		// first clear varlena to 0
		p1 := v.UnsafePtr()
		*(*int64)(p1) = 0
		*(*int64)(unsafe.Add(p1, 8)) = 0
		*(*int64)(unsafe.Add(p1, 16)) = 0
		v[0] = byte(vlen)
		copy(v[1:1+vlen], bs)
		return &v
	} else {
		voff := len(area)
		area = append(area, bs...)
		v.SetOffsetLen(uint32(voff), uint32(vlen))
	}
	return &v
}

// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func compareUint64(a, b uint64) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func TestStringToUint64(t *testing.T) {
	s1 := []byte("abc")
	u1 := VarlenaToUint64Inline(buildVarlenaFromByteSlice(s1))
	require.Equal(t, u1, ByteSliceToUint64(s1))
	s2 := []byte("abcde")
	u2 := VarlenaToUint64Inline(buildVarlenaFromByteSlice(s2))
	require.Equal(t, u2, ByteSliceToUint64(s2))
	s3 := []byte("abcdeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	u3 := VarlenaToUint64(buildVarlenaFromByteSlice(s3), area)
	require.Equal(t, u3, ByteSliceToUint64(s3))
	s4 := []byte("a")
	u4 := VarlenaToUint64(buildVarlenaFromByteSlice(s4), area)
	require.Equal(t, u4, ByteSliceToUint64(s4))
	s5 := []byte("")
	u5 := VarlenaToUint64(buildVarlenaFromByteSlice(s5), area)
	require.Equal(t, u5, ByteSliceToUint64(s5))
	s6 := []byte("A")
	u6 := VarlenaToUint64(buildVarlenaFromByteSlice(s6), area)
	require.Equal(t, u6, ByteSliceToUint64(s6))
	s7 := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	u7 := VarlenaToUint64(buildVarlenaFromByteSlice(s7), area)
	require.Equal(t, u7, ByteSliceToUint64(s7))

	require.Equal(t, bytes.Compare(s1, s2), compareUint64(u1, u2))
	require.Equal(t, bytes.Compare(s1, s3), compareUint64(u1, u3))
	require.Equal(t, bytes.Compare(s1, s4), compareUint64(u1, u4))
	require.Equal(t, bytes.Compare(s1, s5), compareUint64(u1, u5))
	require.Equal(t, bytes.Compare(s1, s6), compareUint64(u1, u6))
	require.Equal(t, bytes.Compare(s1, s7), compareUint64(u1, u7))
	require.Equal(t, bytes.Compare(s2, s3), compareUint64(u2, u3))
	require.Equal(t, bytes.Compare(s2, s4), compareUint64(u2, u4))
	require.Equal(t, bytes.Compare(s2, s5), compareUint64(u2, u5))
	require.Equal(t, bytes.Compare(s2, s6), compareUint64(u2, u6))
	require.Equal(t, bytes.Compare(s2, s7), compareUint64(u2, u7))
	require.Equal(t, bytes.Compare(s3, s4), compareUint64(u3, u4))
	require.Equal(t, bytes.Compare(s3, s5), compareUint64(u3, u5))
	require.Equal(t, bytes.Compare(s3, s6), compareUint64(u3, u6))
	require.Equal(t, bytes.Compare(s3, s7), compareUint64(u3, u7))
	require.Equal(t, bytes.Compare(s4, s5), compareUint64(u4, u5))
	require.Equal(t, bytes.Compare(s4, s6), compareUint64(u4, u6))
	require.Equal(t, bytes.Compare(s4, s7), compareUint64(u4, u7))
	require.Equal(t, bytes.Compare(s5, s6), compareUint64(u5, u6))
	require.Equal(t, bytes.Compare(s5, s7), compareUint64(u5, u7))
	require.Equal(t, bytes.Compare(s6, s7), compareUint64(u6, u7))
}
