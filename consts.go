package go_rtf_builder

const (
	PaperOrientationPortrait  = "orientation_portrait"
	PaperOrientationLandscape = "orientation_landscape"
)

const (
	PaperSizeA5 = "paper_A5"
	PaperSizeA4 = "paper_A4"
	PaperSizeA3 = "paper_A3"
	PaperSizeA2 = "paper_A2"
)

const (
	UnitTwips float32 = 1
	UnitCm    float32 = 567
	UnitInch  float32 = 1440
	UnitPoint float32 = 14.988078
)

const (
	AlignCenter     = "c"
	AlignLeft       = "l"
	AlignRight      = "r"
	AlignJustify    = "j"
	AlignDistribute = "d"
)

const (
	SpaseInLines  = 240
	SpaceInPoints = 20
)

const (
	ColorBlack   = "color_black"
	ColorBlue    = "color_blue"
	ColorAqua    = "color_aqua"
	ColorLime    = "color_lime"
	ColorGreen   = "color_green"
	ColorMagenta = "color_magenta"
	ColorRed     = "color_red"
	ColorYellow  = "color_yellow"
	ColorWhite   = "color_white"
	ColorNavy    = "color_navy"
	ColorTeal    = "color_teal"
	ColorPurple  = "color_purple"
	ColorMaroon  = "color_maroon"
	ColorOlive   = "color_olive"
	ColorGray    = "color_gray"
	ColorSilver  = "color_silver"
)

// Common fonts
const (
	FontTimesNewRoman = "font_times_new_roman"
	FontSymbol        = "font_symbol"
	FontArial         = "font_arial"
	FontComicSansMS   = "font_comic_sans_ms"
	FontCourierNew    = "font_courier_new"
)

const (
	FamilyNil    = "nil"    // Unknown or default fonts (the default)
	FamilyRoman  = "roman"  // Roman, proportionally spaced serif fonts
	FamilySwiss  = "swiss"  // Swiss, proportionally spaced sans serif fonts
	FamilyModern = "modern" // Fixed-pitch serif and sans serif fonts
	FamilyScript = "script" // Script fonts
	FamilyDecor  = "decor"  // Decorative fonts
	FamilyTech   = "tech"   // Technical, symbol, and mathematical fonts
	FamilyBidi   = "bidi"   // Arabic, Hebrew, or other bidirectional font
)

const (
	spaceInPoints = 20
)
