package go_rtf_builder

import "hash/fnv"

type fontTable struct {
	fonts map[uint64]struct {
		position int
		item     fontItem
	}
	nextPosition int
}

func (ft fontTable) AddFont(v fontItem) int {
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

type fontItem struct {
	family string
	name   string
}

func (fi fontItem) hash() uint64 {
	h := fnv.New64()
	h.Write([]byte(fi.family + ":" + fi.name))
	return h.Sum64()
}
