package palette

// TrueColorDark returns the 24-bit palette tuned for dark backgrounds. Every
// color clears 4.5:1 contrast against #1e1e1e, and colors are ordered so that
// the first N are the most separable N. Each call returns a fresh palette the
// caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#fe05fa", "#c7c7c7", "#00a005", "#ff3217", "#2d81fe",
		"#66ff14", "#b973b0", "#c29203", "#fcd304", "#1daeb5",
		"#09fed9", "#f89cfb", "#7bca56", "#f98a7e", "#fe0c99",
		"#83acff", "#fff8ee", "#dafcb1", "#48d8f5", "#c2715f",
		"#6d9f63", "#e5b46b", "#f368d3", "#e2ff25", "#7884c6",
		"#c087f5", "#0cd8af", "#aee86c", "#04df08", "#e6577d",
		"#ec7208", "#ffacc8",
	)
}

// TrueColorLight returns the 24-bit palette tuned for light backgrounds,
// ordered like [TrueColorDark] but measured against #fafafa.
func TrueColorLight() Palette {
	return hexColors(
		"#0404fe", "#4d4d4d", "#c805d2", "#de2500", "#0d8609",
		"#6a019f", "#556dc6", "#9a1553", "#01268d", "#946611",
		"#820eff", "#680e08", "#da0d7e", "#055901", "#9a5a8c",
		"#512155", "#00559a", "#a10197", "#853c1d", "#4534c5",
		"#075cf9", "#684788", "#9050cf", "#027355", "#b20303",
		"#b3514d", "#410186", "#15789c", "#0e08c4", "#5d6704",
		"#7e0672", "#ba3ea5",
	)
}
