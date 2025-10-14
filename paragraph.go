package go_rtf_builder

import (
	"fmt"
	"strings"
)

type Paragraph struct {
	font            FontStyle
	indentFirstLine int
	indentLeft      int
	indentRight     int
	align           string
	contentItems    []ContentItemInterface
}

func (p *Paragraph) render(unit float32) string {
	var result strings.Builder

	indentStr := fmt.Sprintf("\\fi%d \\li%d \\ri%d",
		p.indentFirstLine,
		p.indentLeft,
		p.indentRight)

	result.WriteString(fmt.Sprintf("\n\\pard \\q%s %s {", p.align, indentStr))

	for _, c := range p.contentItems {
		result.WriteString(c.render(unit))
	}
	result.WriteString("}")
	result.WriteString("\\par")

	return result.String()
}

func (p *Paragraph) AddText(content string) *Text {
	text := Text{
		content:       content,
		isBold:        false,
		isItalic:      false,
		isUnderlining: false,
		isScaps:       false,
		isSuper:       false,
		isSub:         false,
		isStrike:      false,
		font:          p.font,
	}

	p.contentItems = append(p.contentItems, &text)

	return &text
}
