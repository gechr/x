package palette

// TrueColorDark returns the first 50 colors of Colorcet's glasbey_light
// palette (https://colorcet.holoviz.org/user_guide/Categorical.html), the
// standard Glasbey palette for dark backgrounds. Colors are ordered from most
// vivid and bright to most muted and dark. Each call returns a fresh palette
// the caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#ff28fd", "#97ff00", "#72b8ff", "#f9ff00", "#ff49b1",
		"#00fdcf", "#ffa52f", "#ff3464", "#ff6200", "#05acc6",
		"#ffcd03", "#afa5ff", "#e48eff", "#00c846", "#829026",
		"#fdf490", "#ff8ec8", "#a0e491", "#ff9070", "#018700",
		"#d3008c", "#b8ba01", "#b500ff", "#c15603", "#90318e",
		"#f2cdff", "#ae083f", "#d3c37c", "#9ee2ff", "#d60000",
		"#009e7c", "#c86e66", "#bceddb", "#77c6ba", "#ac567c",
		"#f4bfb1", "#bc9157", "#9a6900", "#a877ac", "#953f1f",
		"#93ac83", "#c6a5c1", "#bac1df", "#56642a", "#8c9ab1",
		"#5d8c90", "#6b8567", "#916e56", "#366962", "#79525e",
	)
}

// TrueColorLight returns the first 50 colors of Colorcet's glasbey_dark, the
// standard Glasbey palette for light backgrounds, ordered like
// [TrueColorDark].
func TrueColorLight() Palette {
	return hexColors(
		"#00e400", "#ff00cd", "#e252ff", "#8c3bff", "#ff7752",
		"#f2007b", "#bf03b8", "#0000dd", "#15e18c", "#d60000",
		"#ff7ed1", "#8eba00", "#4d33ff", "#e49cff", "#e6a500",
		"#8287ff", "#2db152", "#018700", "#5901a3", "#0774d8",
		"#cc7407", "#00bda3", "#7ecaff", "#c85748", "#bcb6ff",
		"#7e2a8e", "#00acc6", "#c4668e", "#bac389", "#8e7b01",
		"#a03a52", "#9e4b00", "#a57bb8", "#6b004f", "#790000",
		"#e2afaf", "#008069", "#a1c8c8", "#004b00", "#729a7c",
		"#a17569", "#919eb6", "#2f5282", "#546744", "#573b00",
		"#005659", "#5e7b87", "#645472", "#60383b", "#380000",
	)
}
