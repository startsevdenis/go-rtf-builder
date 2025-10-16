package go_rtf_builder

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

type fontTable struct {
	fonts map[uint64]struct {
		position int
		item     fontItem
	}
	nextPosition int
}

func (ft *fontTable) AddFont(v fontItem) int {
	item, found := ft.fonts[v.hash()]
	if !found {
		ft.fonts[v.hash()] = struct {
			position int
			item     fontItem
		}{position: ft.nextPosition, item: v}

		ft.nextPosition += 1
	}

	return item.position
}

func (ft *fontTable) getFontCode(name string) int {
	for _, f := range ft.fonts {
		if f.item.code == name {
			return f.position
		}
	}
	return 0
}

func (ft *fontTable) output() string {
	var result strings.Builder
	for i := range ft.fonts {
		result.WriteString(fmt.Sprintf("{\\f%d%s}", ft.fonts[i].position, ft.fonts[i].item.output()))
	}
	return result.String()
}

type fontItem struct {
	family  string // see consts comment
	charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	prq     int    // Specifies the pitch of a font in the font table.
	name    string
	code    string
}

func (fi fontItem) output() string {
	return fmt.Sprintf("\\f%s\\fprq%d\\fcharset%d %s;", fi.family, fi.prq, fi.charset, fi.name)
}

func (fi fontItem) hash() uint64 {
	h := fnv.New64()
	h.Write([]byte(fi.family + ":" + fi.name + ":" + strconv.Itoa(fi.prq) + ":" + strconv.Itoa(fi.charset)))
	return h.Sum64()
}
