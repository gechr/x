# emulator

```go
import "github.com/gechr/x/terminal/emulator"
```

Package emulator identifies the terminal emulator hosting the process.

## Index

- [Constants](<#constants>)
- [func Detect\(\) string](<#Detect>)
- [func IsKnown\(name string\) bool](<#IsKnown>)
- [func Known\(\) \[\]string](<#Known>)

## Constants

<a name="EnvTerm"></a>Environment variables consulted by Detect.

```go
const (
    EnvTerm             = "TERM"
    EnvTermProgram      = "TERM_PROGRAM"
    EnvTerminalEmulator = "TERMINAL_EMULATOR"
)
```

<a name="Alacritty"></a>Recognized terminal emulator names, as returned by Detect.

```go
const (
    Alacritty       = "alacritty"
    AppleTerminal   = "apple-terminal"
    ConEmu          = "conemu"
    Contour         = "contour"
    Foot            = "foot"
    Ghostty         = "ghostty"
    GNOMETerminal   = "gnome-terminal"
    Hyper           = "hyper"
    ITerm2          = "iterm2"
    JetBrains       = "jetbrains"
    Kitty           = "kitty"
    Konsole         = "konsole"
    Mintty          = "mintty"
    Rio             = "rio"
    ST              = "st"
    Tabby           = "tabby"
    Terminator      = "terminator"
    Termux          = "termux"
    Tilix           = "tilix"
    URxvt           = "urxvt"
    VSCode          = "vscode"
    Warp            = "warp"
    WezTerm         = "wezterm"
    WindowsTerminal = "windows-terminal"
    Zed             = "zed"
)
```

<a name="Detect"></a>

## func [Detect](<https://github.com/gechr/x/blob/main/terminal/emulator/detect.go#L79>)

```go
func Detect() string
```

Detect returns the terminal emulator hosting the process, or empty if it cannot be determined. Detection is best\-effort, based on environment variables inherited from the emulator. Priority: TERM\_PROGRAM, TERMINAL\_EMULATOR, emulator\-specific variables, TERM.

<a name="IsKnown"></a>

## func [IsKnown](<https://github.com/gechr/x/blob/main/terminal/emulator/known.go#L96>)

```go
func IsKnown(name string) bool
```

IsKnown reports whether name matches a known terminal emulator.

<a name="Known"></a>

## func [Known](<https://github.com/gechr/x/blob/main/terminal/emulator/known.go#L91>)

```go
func Known() []string
```

Known returns the set of recognized terminal emulator names.
