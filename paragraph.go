package go_rtf_builder

import (
	"fmt"
	"strings"
)

type Paragraph struct {
	fontCode        string
	fontSize        int
	colorCode       string
	indentFirstLine int
	indentLeft      int
	indentRight     int
	spaceBefore     float32
	spaceAfter      float32
	spaceLines      float32
	align           string
	contentItems    []ContentItemInterface
	fcTables        *fcTables
}

func (p *Paragraph) render(unit float32) string {
	var result strings.Builder

	indentStr := fmt.Sprintf("\\fi%d \\li%d \\ri%d \\sb%.0f \\sa%.0f \\sl%.0f",
		p.indentFirstLine,
		p.indentLeft,
		p.indentRight,
		p.spaceBefore*SpaseInLines,
		p.spaceAfter*SpaseInLines,
		p.spaceLines*SpaseInLines,
	)

	result.WriteString(fmt.Sprintf("\n\\pard \\q%s %s {", p.align, indentStr))

	fmt.Println(p.contentItems)

	for _, c := range p.contentItems {
		result.WriteString(c.render(unit))
	}
	result.WriteString("}")
	result.WriteString("\\par")

	return result.String()
}

func (p *Paragraph) AddNewLine() *Paragraph {
	nl := Text{
		content:  "\\line",
		fcTables: p.fcTables,
	}
	p.contentItems = append(p.contentItems, &nl)
	return p
}

func (p *Paragraph) SetSpaceBetweenLines(space float32) {
	p.spaceLines = space
}

func (p *Paragraph) SetSpaceBefore(spaceBefore float32) {
	p.spaceBefore = spaceBefore
}

func (p *Paragraph) SetSpaceAfter(spaceAfter float32) {
	p.spaceAfter = spaceAfter
}

func (p *Paragraph) AddText(content string, style *FontStyle) *Text {
	text := Text{
		content:       content,
		isBold:        false,
		isItalic:      false,
		isUnderlining: false,
		isScaps:       false,
		isSuper:       false,
		isSub:         false,
		isStrike:      false,
		fontCode:      p.fontCode,
		fontSize:      p.fontSize,
		colorCode:     p.colorCode,
		fcTables:      p.fcTables,
	}

	if style != nil {
		text.fontSize = style.Size
		text.fontCode = style.Name
		text.colorCode = style.Color
	}

	p.contentItems = append(p.contentItems, &text)

	return &text
}
