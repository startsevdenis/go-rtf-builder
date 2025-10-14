package go_rtf_builder

import (
	"fmt"
	"hash/fnv"
	"image/color"
)

type colorTable struct {
	colors map[uint64]struct {
		position int
		item     colorItem
	}
	nextPosition int
}

func (ct colorTable) AddColor(v colorItem) int {
	item, found := ct.colors[v.hash()]
	if !found {
		ct.colors[v.hash()] = struct {
			position int
			item     colorItem
		}{position: ct.nextPosition, item: v}

		ct.nextPosition += 1
	}

	return item.position
}

type colorItem struct {
	rgbColor color.RGBA
	name     string
}

func (ci colorItem) hash() uint64 {
	h := fnv.New64()
	r, g, b, a := ci.rgbColor.RGBA()
	h.Write([]byte(fmt.Sprintf("rgba(%d, %d, %d, %d)", r, g, b, a) + ":" + ci.name))
	return h.Sum64()
}
