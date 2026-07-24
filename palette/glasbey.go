package palette

// TrueColorDark returns the first 50 colors of Colorcet's glasbey_light
// palette (https://colorcet.holoviz.org/user_guide/Categorical.html), the
// standard Glasbey palette for dark backgrounds. Glasbey orders colors so each
// addition is maximally distinct from the preceding colors. Each call returns
// a fresh palette the caller owns.
func TrueColorDark() Palette {
	return hexColors(
		"#d60000", "#018700", "#b500ff", "#05acc6", "#97ff00",
		"#ffa52f", "#ff8ec8", "#79525e", "#00fdcf", "#afa5ff",
		"#93ac83", "#9a6900", "#366962", "#d3008c", "#fdf490",
		"#c86e66", "#9ee2ff", "#00c846", "#a877ac", "#b8ba01",
		"#f4bfb1", "#ff28fd", "#f2cdff", "#009e7c", "#ff6200",
		"#56642a", "#953f1f", "#90318e", "#ff3464", "#a0e491",
		"#8c9ab1", "#829026", "#ae083f", "#77c6ba", "#bc9157",
		"#e48eff", "#72b8ff", "#c6a5c1", "#ff9070", "#d3c37c",
		"#bceddb", "#6b8567", "#916e56", "#f9ff00", "#bac1df",
		"#ac567c", "#ffcd03", "#ff49b1", "#c15603", "#5d8c90",
	)
}

// TrueColorLight returns the first 50 colors of Colorcet's glasbey_dark, the
// standard Glasbey palette for light backgrounds. See [TrueColorDark].
func TrueColorLight() Palette {
	return hexColors(
		"#d60000", "#8c3bff", "#018700", "#00acc6", "#e6a500",
		"#ff7ed1", "#6b004f", "#573b00", "#005659", "#15e18c",
		"#0000dd", "#a17569", "#bcb6ff", "#bf03b8", "#645472",
		"#790000", "#0774d8", "#729a7c", "#ff7752", "#004b00",
		"#8e7b01", "#f2007b", "#8eba00", "#a57bb8", "#5901a3",
		"#e2afaf", "#a03a52", "#a1c8c8", "#9e4b00", "#546744",
		"#bac389", "#5e7b87", "#60383b", "#8287ff", "#380000",
		"#e252ff", "#2f5282", "#7ecaff", "#c4668e", "#008069",
		"#919eb6", "#cc7407", "#7e2a8e", "#00bda3", "#2db152",
		"#4d33ff", "#00e400", "#ff00cd", "#c85748", "#e49cff",
	)
}
