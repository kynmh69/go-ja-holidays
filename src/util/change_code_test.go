package util

import (
	"testing"
)

func TestShiftJISToUTF8(t *testing.T) {
	// テストケースの設定
	cases := []struct {
		sjis []byte
		utf8 string
	}{
		{
			[]byte{0x82, 0xa0, 0x82, 0xa2, 0x82, 0xa4, 0x82, 0xa6, 0x82, 0xa8}, // "あいうえお" をShift JISで表現
			"あいうえお",
		},
		// 他にもテストケースを追加できます
	}

	for _, c := range cases {
		got := string(ShiftJISToUTF8(c.sjis))
		if got != c.utf8 {
			t.Errorf("ShiftJISToUTF8(%q) == %q, want %q", c.sjis, got, c.utf8)
		}
	}
}
