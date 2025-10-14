package go_rtf_builder

import (
	"fmt"
	"strings"
)

type docMargins struct {
	left   int
	right  int
	top    int
	bottom int
}

func (m docMargins) output(unit float32) string {
	return fmt.Sprintf(
		"\n\\margl%.0f\\margr%.0f\\margt%.0f\\margb%.0f",
		float32(m.left)*unit,
		float32(m.right)*unit,
		float32(m.top)*unit,
		float32(m.bottom)*unit,
	)
}

type RtfDocument struct {
	paperSize        string
	paperOrientation string
	margins          docMargins
	header           header
	documentHeader   DocumentHeader
	measureUnit      float32
	contentItems     []ContentItemInterface
}

func NewRtfDocument(paperSize string, paperOrientation string) *RtfDocument {
	return &RtfDocument{
		paperSize:        paperSize,
		paperOrientation: paperOrientation,
		margins: docMargins{
			left:   0,
			right:  0,
			top:    0,
			bottom: 0,
		},
		measureUnit: UnitCm,
		header: header{
			version: "1",
			charSet: "ansi",
			deff:    "0",
		}}
}

func (doc *RtfDocument) AddParagraph(font FontStyle) *Paragraph {
	var paragraph = Paragraph{
		font:            font,
		indentFirstLine: 10,
		indentLeft:      0,
		indentRight:     0,
		align:           AlignLeft,
		contentItems:    nil,
	}
	doc.contentItems = append(doc.contentItems, &paragraph)
	return &paragraph
}

// SetMarginUnit set margin unit
func (doc *RtfDocument) SetMarginUnit(value float32) *RtfDocument {
	doc.measureUnit = value
	return doc
}

// SetMarginLeft set margin left
func (doc *RtfDocument) SetMarginLeft(value int) *RtfDocument {
	doc.margins.left = value
	return doc
}

// SetMarginRight set margin right
func (doc *RtfDocument) SetMarginRight(value int) *RtfDocument {
	doc.margins.right = value
	return doc
}

// SetMarginTop set margin top
func (doc *RtfDocument) SetMarginTop(value int) *RtfDocument {
	doc.margins.top = value
	return doc
}

// SetMarginBottom set margin bottom
func (doc *RtfDocument) SetMarginBottom(value int) *RtfDocument {
	doc.margins.bottom = value
	return doc
}

// SetMargins set all margins (top, right, bottom, left)
func (doc *RtfDocument) SetMargins(top, right, bottom, left int) *RtfDocument {
	doc.margins.top = top
	doc.margins.bottom = bottom
	doc.margins.left = left
	doc.margins.right = right
	return doc
}

func (doc *RtfDocument) Output() []byte {
	var result strings.Builder

	result.WriteString("{")
	result.WriteString(doc.header.output())
	if doc.paperOrientation == PaperOrientationLandscape {
		result.WriteString(fmt.Sprintf("\n\\landscape"))
	}

	for _, c := range doc.contentItems {
		result.WriteString(fmt.Sprintf("\n%s", c.render(doc.measureUnit)))
	}

	result.WriteString(doc.margins.output(doc.measureUnit))

	result.WriteString("\n}")

	return []byte(result.String())
}
