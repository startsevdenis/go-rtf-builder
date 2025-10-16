package go_rtf_builder

type FontStyle struct {
	Size  int
	Name  string
	Color string
}

func NewFontStyle(size int, name string, color string) *FontStyle {
	return &FontStyle{
		Size:  size,
		Name:  name,
		Color: color,
	}
}
