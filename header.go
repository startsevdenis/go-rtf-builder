package go_rtf_builder

import (
	"fmt"
	"strings"
)

type header struct {
	version string
	charSet string
	deff    string
}

func (h header) output() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.version, h.charSet, h.deff))
	return res.String()
}
