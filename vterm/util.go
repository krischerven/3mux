package vterm

import (
	"strconv"
	"strings"

	"github.com/aaronduino/i3-tmux/cursor"
)

// parseSemicolonNumSeq parses a series of numbers separated by semicolons, replacing empty values with the given default value
func parseSemicolonNumSeq(s string, d int) []int {
	s = strings.TrimSpace(s)

	if s == "" {
		return []int{d}
	}

	parts := strings.Split(s, ";")

	out := []int{}
	for _, part := range parts {
		if part == "" {
			out = append(out, d)
		} else {
			num, err := strconv.ParseInt(part, 10, 32)
			if err != nil {
				continue // fixme?
			}

			out = append(out, int(num))
		}
	}
	return out
}

func (v *VTerm) debug(s string) {
	for i, r := range []rune(s) {
		v.out <- Char{
			Rune: r,
			Cursor: cursor.Cursor{
				X: i + 20,
				Y: 3,
			},
		}
	}
}
