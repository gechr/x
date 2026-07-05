package emulator_test

import (
	"fmt"

	"github.com/gechr/x/terminal/emulator"
)

func ExampleKnown() {
	fmt.Println(emulator.Known())
	// Output:
	// [alacritty apple-terminal conemu contour foot ghostty gnome-terminal hyper iterm2 jetbrains kitty konsole mintty rio screen st tabby terminator termux tilix tmux urxvt vscode warp wezterm windows-terminal zed]
}

func ExampleIsKnown() {
	fmt.Println(emulator.IsKnown(emulator.Kitty))
	fmt.Println(emulator.IsKnown("xterm-256color"))
	// Output:
	// true
	// false
}
