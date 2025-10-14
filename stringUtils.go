package go_rtf_builder

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf16"
)

func convertNonASCIIToUTF16(text string) string {
	var res strings.Builder
	for _, r := range text {
		switch {
		case unicode.Is(unicode.Cyrillic, r) || r == '№' || r == '°':
			res.WriteString(fmt.Sprintf("\\u%d\\'3f", utf16.Encode([]rune{r})[0]))
		default:
			res.WriteString(string(r))
		}

	}
	return res.String()
}
