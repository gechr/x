package palette

// TrueColorDark returns the 24-bit palette tuned for dark backgrounds. Every
// color clears 4.5:1 contrast against #1e1e1e, and colors are ordered so that
// the first N are the most separable N. Each call returns a fresh palette the
// caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#ff00ff", "#ff3217", "#2d81fe", "#66ff66", "#ffcc00",
		"#0caeb6", "#3dffec", "#ff6699", "#4ad103", "#f76e6e",
		"#83acff", "#d4ffa1", "#00ccff", "#e5b46b", "#ff75ff",
		"#ffff00", "#c087f5", "#a6ea5a", "#ff3366", "#ffacc8",
		"#ff8938", "#aaaf03", "#ffa500",
	)
}

// TrueColorLight returns the 24-bit palette tuned for light backgrounds,
// ordered like [TrueColorDark] but measured against #fafafa.
func TrueColorLight() Palette {
	return hexColors(
		"#0404fe", "#de2500", "#0d8609", "#6a019f", "#556dc6",
		"#9f0053", "#966501", "#820eff", "#6d0000",
		"#db047e", "#055901", "#5f0169", "#a10197", "#027355",
		"#b20303", "#e00131", "#15789c", "#5d6704", "#ce02b6",
		"#0249c9", "#5c3db6", "#350078", "#1f00b9",
	)
}
