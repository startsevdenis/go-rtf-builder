package go_rtf_builder

type FontStyle struct {
	size    int
	family  string
	color   string
	bgColor string
}

func NewFontStyle(size int, family string) *FontStyle {
	return &FontStyle{
		size:    size,
		family:  family,
		bgColor: "#ffffff",
		color:   "#000000",
	}
}
