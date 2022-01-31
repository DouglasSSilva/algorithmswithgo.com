package module01

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"alphabet", "tebahpla"},
		{"日本語", "語本日"},
		{"X", "X"},
		{"b\u0301", "b\u0301"},
		{"😎⚽", "⚽😎"},
		{"Les Mise\u0301rables", "selbare\u0301siM seL"},
		{"ab\u0301cde", "edcb\u0301a"},
		{"This `\xc5` is an invalid UTF8 character", "retcarahc 8FTU dilavni na si `�` sihT"},
		{"The quick bròwn 狐 jumped over the lazy 犬", "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", got, tc.want)
			}
		})
	}
}
