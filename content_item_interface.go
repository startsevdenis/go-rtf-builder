package go_rtf_builder

type ContentItemInterface interface {
	render(unit float32) string
}
