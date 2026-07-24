package palette

// TrueColorDark returns a vivid 36-color Glasbey palette curated for dark
// backgrounds. It was generated with glasbey using lightness bounds of 50-90
// and chroma bounds of 40-100, then optimized as a complete set for perceptual
// separation. Each color has a contrast ratio of at least 4.5:1 against black.
// Each call returns a fresh palette the caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#ff3d00", "#1cae00", "#ce51ff", "#1ca2ca", "#ffc200",
		"#00ffba", "#ffa2ca", "#bab6ff", "#a6ff00", "#ff00b2",
		"#be8239", "#92e7ff", "#39a692", "#aace79", "#6d8eff",
		"#ffaa75", "#f39aff", "#c275aa", "#9a9a00", "#d76d69",
		"#04e351", "#e7f37d", "#a286ce", "#ff7d00", "#75beff",
		"#ff3161", "#00dbc6", "#6da25d", "#ff51ef", "#c2ce00",
		"#caae55", "#b6ffb6", "#96ffef", "#fb699e", "#10c682",
		"#00cae7",
	)
}

// TrueColorLight returns a vivid 36-color Glasbey palette curated for light
// backgrounds. It was generated with glasbey using lightness bounds of 10-35
// and chroma bounds of 40-100, then optimized as a complete set for perceptual
// separation. Each color has a contrast ratio of at least 4.5:1 against white.
// See [TrueColorDark].
func TrueColorLight() Palette {
	return hexColors(
		"#c60000", "#8624ff", "#148200", "#65004d", "#04758a",
		"#080096", "#714900", "#003500", "#b204b2", "#450c00",
		"#a24965", "#0000fb", "#550096", "#6d0014", "#6561a2",
		"#007d5d", "#004971", "#414d00", "#865192", "#a64d20",
		"#be0075", "#390449", "#5900df", "#752800", "#756d00",
		"#920075", "#45001c", "#28185d", "#414186", "#b2103d",
		"#790092", "#a218df", "#005d2d", "#5551ff", "#0071c6",
		"#792d45",
	)
}
