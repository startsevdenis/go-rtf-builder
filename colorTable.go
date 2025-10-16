package go_rtf_builder

import (
	"fmt"
	"hash/fnv"
	"image/color"
)

type colorTable struct {
	colors []struct {
		position int
		item     colorItem
	}
	nextPosition int
}

func (ct *colorTable) output() string {
	var result string
	for _, c := range ct.colors {
		result += c.item.output()
	}
	return result
}

func (ft *colorTable) getColorCode(name string) int {
	for _, f := range ft.colors {
		if f.item.name == name {
			return f.position
		}
	}
	return 0
}

func (ct *colorTable) AddColor(v colorItem) int {
	found := false
	position := 0
	for _, i := range ct.colors {
		if i.item.name == v.name {
			found = true
		}
	}

	if !found {
		ct.colors = append(ct.colors, struct {
			position int
			item     colorItem
		}{position: ct.nextPosition, item: v})
		position = ct.nextPosition
		ct.nextPosition += 1
	}

	return position
}

type colorItem struct {
	rgbColor color.RGBA
	name     string
}

func (ci colorItem) output() string {
	r, g, b, _ := ci.rgbColor.RGBA()
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", r/256, g/256, b/256)
}

func (ci colorItem) hash() uint64 {
	h := fnv.New64()
	r, g, b, _ := ci.rgbColor.RGBA()
	h.Write([]byte(fmt.Sprintf("rgba(%d, %d, %d):%s", r, g, b, ci.name)))
	return h.Sum64()
}
