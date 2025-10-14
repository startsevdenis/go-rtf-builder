package go_rtf_builder

import (
	"strings"
)

type Text struct {
	content       string
	isBold        bool
	isItalic      bool
	isUnderlining bool
	isScaps       bool
	isSuper       bool
	isSub         bool
	isStrike      bool
	font          FontStyle
}

func (t *Text) render(unit float32) string {
	var result strings.Builder

	if t.isBold {
		result.WriteString("\\b ")
	}

	if t.isItalic {
		result.WriteString("\\i ")
	}

	if t.isUnderlining {
		result.WriteString("\\ul")
	}

	if t.isScaps {
		result.WriteString("\\scaps")
	}

	if t.isSuper {
		result.WriteString("\\super")
	}

	if t.isSub {
		result.WriteString("\\sub")
	}

	if t.isStrike {
		result.WriteString("\\strike")
	}

	convText := convertNonASCIIToUTF16(t.content)
	result.WriteString(convText)
	//result.WriteString(fmt.Sprintf("\n\\sb0 \\sa0 \\fs%d\\f%d \\cf%d{%s}\\f0", t.font.size, t.font.family, text.colorCode, convText))
	if t.isBold {
		result.WriteString("\\b0 ")
	}

	if t.isItalic {
		result.WriteString("\\i0 ")
	}

	if t.isUnderlining {
		result.WriteString("\\ul0 ")
	}

	if t.isScaps {
		result.WriteString("\\scaps0 ")
	}

	if t.isSuper {
		result.WriteString("\\super0 ")
	}

	if t.isSub {
		result.WriteString("\\sub0 ")
	}

	if t.isStrike {
		result.WriteString("\\strike0 ")
	}

	return result.String()
}

func (t *Text) SetBold() *Text {
	t.isBold = true
	return t
}

func (t *Text) SetItalic() *Text {
	t.isItalic = true
	return t
}

func (t *Text) SetUnderlining() *Text {
	t.isUnderlining = true
	return t
}

func (t *Text) SetScaps() *Text {
	t.isScaps = true
	return t
}

func (t *Text) SetSub() *Text {
	t.isSub = true
	return t
}

func (t *Text) SetStrike() *Text {
	t.isStrike = true
	return t
}

func (t *Text) SetFont(f FontStyle) *Text {
	t.font = f
	return t
}
