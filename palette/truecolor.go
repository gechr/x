package palette

// TrueColorDark returns the 24-bit palette tuned for dark backgrounds. Every
// color clears 4.5:1 contrast against #1e1e1e and sits at least 0.08 CIEDE2000
// from every other, and colors are ordered so that consecutive entries stay
// distinct from one another as well. Each call returns a fresh palette the
// caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#00ccff", "#fffe44", "#66ff66", "#ff3366", "#ff75ff",
		"#ffa500", "#2d81fe", "#3dffec", "#c087f5", "#f44527",
		"#d4ffa1", "#62ce34", "#0caeb6", "#ffacc8", "#e554e2",
		"#aaaf03", "#ff8938", "#ff6699", "#a6ea5a", "#ffcc00",
		"#f76e6e", "#e5b46b", "#83acff",
	)
}

// TrueColorLight returns the 24-bit palette tuned for light backgrounds,
// ordered like [TrueColorDark] but measured against #fafafa.
func TrueColorLight() Palette {
	return hexColors(
		"#8231df", "#5d6704", "#d9300c", "#15789c", "#9f0053",
		"#027355", "#2909b0", "#966501", "#b20303", "#556dc6",
		"#db047e", "#0d8609", "#350078", "#2210f4", "#6d0000",
		"#e00131", "#ca1bb2", "#055901", "#5f0169", "#0249c9",
		"#5c3db6", "#a10197", "#6a019f",
	)
}
