# emulator

```go
import "github.com/gechr/x/terminal/emulator"
```

Package emulator identifies the terminal emulator hosting the process.

## Index

- [Constants](<#constants>)
- [func Detect() string](<#Detect>)
- [func IsKnown(name string) bool](<#IsKnown>)
- [func Known() \[\]string](<#Known>)
- [func SupportsGraphemes() bool](<#SupportsGraphemes>)

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

## func [Detect](<https://github.com/gechr/x/blob/main/terminal/emulator/detect.go#L83>)

```go
func Detect() string
```

**Detect** returns the terminal emulator hosting the process, or empty if it cannot be determined. Detection is best-effort, based on environment variables inherited from the emulator. Priority: `TERM`, `TERM_PROGRAM`, `TERMINAL_EMULATOR`, emulator-specific variables. `TERM` wins because the innermost emulator always sets it fresh for its own session, whereas `TERM_PROGRAM` and marker variables leak through from an outer terminal when one emulator is launched from another that does not scrub them (e.g. kitty launched from iTerm2 inherits both `TERM_PROGRAM=iTerm.app` and `ITERM_SESSION_ID`).

<a name="IsKnown"></a>

## func [IsKnown](<https://github.com/gechr/x/blob/main/terminal/emulator/known.go#L96>)

```go
func IsKnown(name string) bool
```

**IsKnown** reports whether name matches a known terminal emulator.

<a name="Known"></a>

## func [Known](<https://github.com/gechr/x/blob/main/terminal/emulator/known.go#L91>)

```go
func Known() []string
```

**Known** returns the set of recognized terminal emulator names.

<a name="SupportsGraphemes"></a>

## func [SupportsGraphemes](<https://github.com/gechr/x/blob/main/terminal/emulator/graphemes.go#L21>)

```go
func SupportsGraphemes() bool
```

**SupportsGraphemes** reports whether the detected terminal emulator is known to measure text in grapheme clusters rather than per-codepoint wcwidth. It returns false when the emulator cannot be determined.
