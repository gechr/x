# ansi

```go
import "github.com/gechr/x/ansi"
```

Package ansi provides ANSI-aware text wrapping, hyperlinks, and TTY fallback.

## Index

- [Constants](<#constants>)
- [func CursorBackward(n int) string](<#CursorBackward>)
- [func CursorDown(n int) string](<#CursorDown>)
- [func CursorForward(n int) string](<#CursorForward>)
- [func CursorHorizontalAbsolute(col int) string](<#CursorHorizontalAbsolute>)
- [func CursorNextLine(n int) string](<#CursorNextLine>)
- [func CursorPosition(col, row int) string](<#CursorPosition>)
- [func CursorPreviousLine(n int) string](<#CursorPreviousLine>)
- [func CursorUp(n int) string](<#CursorUp>)
- [func DeleteCharacter(n int) string](<#DeleteCharacter>)
- [func DeleteLine(n int) string](<#DeleteLine>)
- [func EraseCharacter(n int) string](<#EraseCharacter>)
- [func EraseDisplay(n int) string](<#EraseDisplay>)
- [func EraseLine(n int) string](<#EraseLine>)
- [func InsertCharacter(n int) string](<#InsertCharacter>)
- [func InsertLine(n int) string](<#InsertLine>)
- [func ScrollDown(n int) string](<#ScrollDown>)
- [func ScrollUp(n int) string](<#ScrollUp>)
- [func SetCursorStyle(style int) string](<#SetCursorStyle>)
- [func SetIconName(s string) string](<#SetIconName>)
- [func SetIconNameWindowTitle(s string) string](<#SetIconNameWindowTitle>)
- [func SetWindowTitle(s string) string](<#SetWindowTitle>)
- [func StringWidth(s string) int](<#StringWidth>)
- [func Strip(s string) string](<#Strip>)
- [func Truncate(s string, length int, tail string) string](<#Truncate>)
- [func WrapHard(s string, width int) string](<#WrapHard>)
- [func WrapSoft(s string, width int) string](<#WrapSoft>)
- [type ANSI](<#ANSI>)
  - [func Auto(files ...\*os.File) \*ANSI](<#Auto>)
  - [func Force() \*ANSI](<#Force>)
  - [func Never() \*ANSI](<#Never>)
  - [func New(opts ...Option) \*ANSI](<#New>)
  - [func (w \*ANSI) Hyperlink(url, text string) string](<#ANSI.Hyperlink>)
  - [func (w \*ANSI) Terminal() bool](<#ANSI.Terminal>)
- [type HyperlinkFallback](<#HyperlinkFallback>)
- [type Method](<#Method>)
- [type Option](<#Option>)
  - [func WithHyperlinkFallback(fallback HyperlinkFallback) Option](<#WithHyperlinkFallback>)
  - [func WithTerminal(v bool) Option](<#WithTerminal>)
- [type WrapOption](<#WrapOption>)
  - [func WithBreakpoints(chars string) WrapOption](<#WithBreakpoints>)
  - [func WithPreserveStyle(preserve bool) WrapOption](<#WithPreserveStyle>)
  - [func WithWidth(width int) WrapOption](<#WithWidth>)
  - [func WithWidthFunc(fn func() int) WrapOption](<#WithWidthFunc>)
  - [func WithWrapHard() WrapOption](<#WithWrapHard>)
  - [func WithWrapSoft() WrapOption](<#WithWrapSoft>)
- [type Wrapper](<#Wrapper>)
  - [func NewWrapper(opts ...WrapOption) \*Wrapper](<#NewWrapper>)
  - [func (w \*Wrapper) Wrap(s string) string](<#Wrapper.Wrap>)

## Constants

<a name="EraseLineRight"></a>Erase constants (delegated).

```go
const (
    EraseLineRight    = xansi.EraseLineRight
    EraseLineLeft     = xansi.EraseLineLeft
    EraseEntireLine   = xansi.EraseEntireLine
    EraseScreenBelow  = xansi.EraseScreenBelow
    EraseScreenAbove  = xansi.EraseScreenAbove
    EraseEntireScreen = xansi.EraseEntireScreen
)
```

<a name="CursorUp1"></a>Cursor movement - single-step constants. Upstream's preferred names are the opaque CUU1/CUD1/CUF1/CUB1 spellings; we keep the readable aliases here and forward to the non-deprecated upstream symbols.

```go
const (
    CursorUp1    = xansi.CUU1
    CursorDown1  = xansi.CUD1
    CursorRight1 = xansi.CUF1
    CursorLeft1  = xansi.CUB1
)
```

<a name="ShowCursor"></a>Cursor visibility (DECTCEM).

```go
const (
    ShowCursor = xansi.ShowCursor
    HideCursor = xansi.HideCursor
)
```

<a name="SaveCursorPosition"></a>Cursor save/restore (SCOSC / SCORC).

```go
const (
    SaveCursorPosition    = xansi.SaveCurrentCursorPosition
    RestoreCursorPosition = xansi.RestoreCurrentCursorPosition
)
```

<a name="Focus"></a>**Focus** reporting events (terminal → app, when focus events are enabled).

```go
const (
    Focus = xansi.Focus
    Blur  = xansi.Blur
)
```

<a name="EnterAltScreen"></a>Alt-screen buffer (DEC 1049: also saves/restores cursor).

```go
const (
    EnterAltScreen = xansi.SetModeAltScreenSaveCursor
    ExitAltScreen  = xansi.ResetModeAltScreenSaveCursor
)
```

<a name="EnableBracketedPaste"></a>Bracketed paste mode (DEC 2004).

```go
const (
    EnableBracketedPaste  = xansi.SetModeBracketedPaste
    DisableBracketedPaste = xansi.ResetModeBracketedPaste
)
```

<a name="EnableFocusEvents"></a>Focus event reporting mode (DEC 1004).

```go
const (
    EnableFocusEvents  = xansi.SetModeFocusEvent
    DisableFocusEvents = xansi.ResetModeFocusEvent
)
```

<a name="EnableSyncOutput"></a>Synchronized output mode (DEC 2026). Wrap a repaint in Enable/Disable so the terminal applies the whole frame atomically; terminals without support ignore both sequences.

```go
const (
    EnableSyncOutput  = xansi.SetModeSynchronizedOutput
    DisableSyncOutput = xansi.ResetModeSynchronizedOutput
)
```

<a name="RequestCursorPosition"></a>Terminal queries - the terminal replies with a corresponding report.

```go
const (
    // RequestCursorPosition asks the terminal for the current cursor
    // position (DSR 6). The reply is CSI <row> ; <col> R.
    RequestCursorPosition = xansi.RequestCursorPositionReport

    // RequestExtendedCursorPosition asks for the unambiguous DEC form
    // (DSR ?6). Preferred over [RequestCursorPosition] when the terminal
    // supports it, because the reply is distinguishable from key input.
    RequestExtendedCursorPosition = xansi.RequestExtendedCursorPositionReport

    // RequestPrimaryDeviceAttributes asks the terminal to identify itself
    // (DA1). Reply format varies by terminal.
    RequestPrimaryDeviceAttributes = xansi.RequestPrimaryDeviceAttributes
)
```

<a name="KeypadApplicationMode"></a>Keypad modes (DECKPAM / DECKPNM). Application mode makes the numeric keypad emit escape sequences rather than raw digits; relevant when the host reads keys in raw mode.

```go
const (
    KeypadApplicationMode = xansi.KeypadApplicationMode
    KeypadNumericMode     = xansi.KeypadNumericMode
)
```

<a name="EnableMouseX10"></a>Mouse tracking modes. Each pair toggles a DEC private mode; the encoding mode (SGR) is usually enabled alongside one of the event modes:

```text
X10         legacy, presses only (no release)
Normal      press + release
ButtonEvent press + release + motion while button held
AnyEvent    press + release + motion at all times
SGR         extended coordinate encoding, required for columns > 223
```

```go
const (
    EnableMouseX10          = xansi.SetModeMouseX10
    DisableMouseX10         = xansi.ResetModeMouseX10
    EnableMouseNormal       = xansi.SetModeMouseNormal
    DisableMouseNormal      = xansi.ResetModeMouseNormal
    EnableMouseButtonEvent  = xansi.SetModeMouseButtonEvent
    DisableMouseButtonEvent = xansi.ResetModeMouseButtonEvent
    EnableMouseAnyEvent     = xansi.SetModeMouseAnyEvent
    DisableMouseAnyEvent    = xansi.ResetModeMouseAnyEvent
    EnableMouseSGR          = xansi.SetModeMouseExtSgr
    DisableMouseSGR         = xansi.ResetModeMouseExtSgr
)
```

<a name="WcWidth"></a>Display width methods.

```go
const (
    WcWidth       = xansi.WcWidth
    GraphemeWidth = xansi.GraphemeWidth
)
```

<a name="CSI"></a>**CSI** is the Control Sequence Introducer prefix (ESC + `[`). Exported for callers that parse or construct their own CSI-family sequences.

```go
const CSI = "\x1b["
```

<a name="ClearLine"></a>**ClearLine** erases the entire current line (EL 2) and returns the cursor to column 0 (CR). This is a convenience composite not present in upstream; every other escape here delegates to [github.com/charmbracelet/x/ansi](<https://pkg.go.dev/github.com/charmbracelet/x/ansi>).

```go
const ClearLine = xansi.EraseEntireLine + "\r"
```

<a name="CursorHomePosition"></a>**CursorHomePosition** moves the cursor to row 1, column 1 (`"\x1b[H"`).

```go
const CursorHomePosition = xansi.CursorHomePosition
```

<a name="RequestNameVersion"></a>**RequestNameVersion** is the XTVERSION query (`"\x1b[>q"`): asks the terminal for its name and version. Complements the DA1 query.

```go
const RequestNameVersion = xansi.RequestNameVersion
```

<a name="ResetStyle"></a>**ResetStyle** is the SGR reset sequence (`"\x1b[m"`): clears all text attributes and colours. Included because writing any styled output typically requires emitting this afterwards.

```go
const ResetStyle = xansi.ResetStyle
```

<a name="CursorBackward"></a>

## func [CursorBackward](<https://github.com/gechr/x/blob/main/ansi/escape.go#L56>)

```go
func CursorBackward(n int) string
```

**CursorBackward** returns the CUB sequence: move cursor left `n` columns.

<a name="CursorDown"></a>

## func [CursorDown](<https://github.com/gechr/x/blob/main/ansi/escape.go#L50>)

```go
func CursorDown(n int) string
```

**CursorDown** returns the CUD sequence: move cursor down `n` lines.

<a name="CursorForward"></a>

## func [CursorForward](<https://github.com/gechr/x/blob/main/ansi/escape.go#L53>)

```go
func CursorForward(n int) string
```

**CursorForward** returns the CUF sequence: move cursor right `n` columns.

<a name="CursorHorizontalAbsolute"></a>

## func [CursorHorizontalAbsolute](<https://github.com/gechr/x/blob/main/ansi/escape.go#L66>)

```go
func CursorHorizontalAbsolute(col int) string
```

**CursorHorizontalAbsolute** returns the CHA sequence: move to column `col` on the current line.

<a name="CursorNextLine"></a>

## func [CursorNextLine](<https://github.com/gechr/x/blob/main/ansi/escape.go#L59>)

```go
func CursorNextLine(n int) string
```

**CursorNextLine** returns the CNL sequence: move down `n` lines and to column 1.

<a name="CursorPosition"></a>

## func [CursorPosition](<https://github.com/gechr/x/blob/main/ansi/escape.go#L70>)

```go
func CursorPosition(col, row int) string
```

**CursorPosition** returns the CUP sequence: move to (`col`, `row`). Coordinates are 1-based.

<a name="CursorPreviousLine"></a>

## func [CursorPreviousLine](<https://github.com/gechr/x/blob/main/ansi/escape.go#L62>)

```go
func CursorPreviousLine(n int) string
```

**CursorPreviousLine** returns the CPL sequence: move up `n` lines and to column 1.

<a name="CursorUp"></a>

## func [CursorUp](<https://github.com/gechr/x/blob/main/ansi/escape.go#L47>)

```go
func CursorUp(n int) string
```

**CursorUp** returns the CUU sequence: move cursor up `n` lines.

<a name="DeleteCharacter"></a>

## func [DeleteCharacter](<https://github.com/gechr/x/blob/main/ansi/escape.go#L161>)

```go
func DeleteCharacter(n int) string
```

**DeleteCharacter** returns the DCH sequence: delete `n` characters at the cursor, pulling subsequent characters left.

<a name="DeleteLine"></a>

## func [DeleteLine](<https://github.com/gechr/x/blob/main/ansi/escape.go#L153>)

```go
func DeleteLine(n int) string
```

**DeleteLine** returns the DL sequence: delete `n` lines starting at the cursor, pulling subsequent lines up.

<a name="EraseCharacter"></a>

## func [EraseCharacter](<https://github.com/gechr/x/blob/main/ansi/escape.go#L34>)

```go
func EraseCharacter(n int) string
```

**EraseCharacter** returns the ECH sequence: erase `n` characters from the cursor position (no cursor movement).

<a name="EraseDisplay"></a>

## func [EraseDisplay](<https://github.com/gechr/x/blob/main/ansi/escape.go#L30>)

```go
func EraseDisplay(n int) string
```

**EraseDisplay** returns the ED sequence. `n` selects the variant: 0 = below cursor, 1 = above cursor, 2 = entire screen.

<a name="EraseLine"></a>

## func [EraseLine](<https://github.com/gechr/x/blob/main/ansi/escape.go#L26>)

```go
func EraseLine(n int) string
```

**EraseLine** returns the EL sequence. `n` selects the variant: 0 = right of cursor, 1 = left of cursor, 2 = entire line.

<a name="InsertCharacter"></a>

## func [InsertCharacter](<https://github.com/gechr/x/blob/main/ansi/escape.go#L157>)

```go
func InsertCharacter(n int) string
```

**InsertCharacter** returns the ICH sequence: insert `n` blank characters at the cursor, shifting existing characters right.

<a name="InsertLine"></a>

## func [InsertLine](<https://github.com/gechr/x/blob/main/ansi/escape.go#L149>)

```go
func InsertLine(n int) string
```

**InsertLine** returns the IL sequence: insert `n` blank lines at the cursor, pushing existing lines down.

<a name="ScrollDown"></a>

## func [ScrollDown](<https://github.com/gechr/x/blob/main/ansi/escape.go#L143>)

```go
func ScrollDown(n int) string
```

**ScrollDown** returns the SD sequence: scroll viewport down `n` lines (content moves down; new blank lines appear at the top).

<a name="ScrollUp"></a>

## func [ScrollUp](<https://github.com/gechr/x/blob/main/ansi/escape.go#L139>)

```go
func ScrollUp(n int) string
```

**ScrollUp** returns the SU sequence: scroll viewport up `n` lines (content moves up; new blank lines appear at the bottom).

<a name="SetCursorStyle"></a>

## func [SetCursorStyle](<https://github.com/gechr/x/blob/main/ansi/escape.go#L173>)

```go
func SetCursorStyle(style int) string
```

**SetCursorStyle** returns the DECSCUSR sequence. Style selects shape and blink state:

```text
0: default         1: blinking block    2: steady block
3: blinking under  4: steady under      5: blinking bar    6: steady bar
```

<a name="SetIconName"></a>

## func [SetIconName](<https://github.com/gechr/x/blob/main/ansi/escape.go#L181>)

```go
func SetIconName(s string) string
```

**SetIconName** sets the icon/tab name only (OSC 1).

<a name="SetIconNameWindowTitle"></a>

## func [SetIconNameWindowTitle](<https://github.com/gechr/x/blob/main/ansi/escape.go#L185>)

```go
func SetIconNameWindowTitle(s string) string
```

**SetIconNameWindowTitle** sets both the icon name and window title in a single sequence (OSC 0).

<a name="SetWindowTitle"></a>

## func [SetWindowTitle](<https://github.com/gechr/x/blob/main/ansi/escape.go#L178>)

```go
func SetWindowTitle(s string) string
```

**SetWindowTitle** sets the terminal window title only (OSC 2).

<a name="StringWidth"></a>

## func [StringWidth](<https://github.com/gechr/x/blob/main/ansi/text.go#L22>)

```go
func StringWidth(s string) int
```

**StringWidth** returns the display width of a string in cells, ignoring ANSI escape codes and accounting for wide characters. Uses grapheme clustering.

<a name="Strip"></a>

## func [Strip](<https://github.com/gechr/x/blob/main/ansi/text.go#L15>)

```go
func Strip(s string) string
```

**Strip** removes ANSI escape codes from a string.

<a name="Truncate"></a>

## func [Truncate](<https://github.com/gechr/x/blob/main/ansi/text.go#L28>)

```go
func Truncate(s string, length int, tail string) string
```

**Truncate** truncates a string to a given cell width, appending `tail` if the string was truncated. ANSI escape codes are preserved.

<a name="WrapHard"></a>

## func [WrapHard](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L133>)

```go
func WrapHard(s string, width int) string
```

**WrapHard** wraps `s` at exactly `width` columns, breaking mid-word if needed. ANSI styles are preserved.

<a name="WrapSoft"></a>

## func [WrapSoft](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L127>)

```go
func WrapSoft(s string, width int) string
```

**WrapSoft** wraps `s` to fit within `width` columns, breaking at space boundaries. Words longer than `width` are hard-wrapped. ANSI styles are preserved.

<a name="ANSI"></a>

## type [ANSI](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L12-L15>)

**ANSI** produces ANSI-aware output, falling back to plain text when the output is not a terminal.

```go
type ANSI struct {
    // contains filtered or unexported fields
}
```

<a name="Auto"></a>

### func [Auto](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L39>)

```go
func Auto(files ...*os.File) *ANSI
```

**Auto** creates an ANSI that auto-detects whether the output is a terminal. All provided `files` must be terminals for ANSI output to be enabled. Defaults to [os.Stdout](<https://pkg.go.dev/os#Stdout>) if no `files` are provided.

<a name="Force"></a>

### func [Force](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L32>)

```go
func Force() *ANSI
```

**Force** creates an ANSI with ANSI output unconditionally enabled.

<a name="Never"></a>

### func [Never](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L27>)

```go
func Never() *ANSI
```

**Never** creates an ANSI with ANSI output unconditionally disabled.

<a name="New"></a>

### func [New](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L18>)

```go
func New(opts ...Option) *ANSI
```

**New** creates an ANSI with the given options.

<a name="ANSI.Hyperlink"></a>

### func (\*ANSI) [Hyperlink](<https://github.com/gechr/x/blob/main/ansi/hyperlink.go#L29>)

```go
func (w *ANSI) Hyperlink(url, text string) string
```

**Hyperlink** creates an OSC 8 terminal hyperlink. When the output is not a terminal, the [HyperlinkFallback](<#HyperlinkFallback>) mode controls how the link is rendered in plain text.

<a name="ANSI.Terminal"></a>

### func (\*ANSI) [Terminal](<https://github.com/gechr/x/blob/main/ansi/ansi.go#L52>)

```go
func (w *ANSI) Terminal() bool
```

**Terminal** reports whether the output target is a terminal.

<a name="HyperlinkFallback"></a>

## type [HyperlinkFallback](<https://github.com/gechr/x/blob/main/ansi/hyperlink.go#L6>)

**HyperlinkFallback** controls how hyperlinks render when the output is not a terminal.

```go
type HyperlinkFallback int
```

<a name="HyperlinkFallbackExpanded"></a>

```go
const (
    // HyperlinkFallbackExpanded renders "text (url)".
    HyperlinkFallbackExpanded HyperlinkFallback = iota
    // HyperlinkFallbackMarkdown renders "[text](url)".
    HyperlinkFallbackMarkdown
    // HyperlinkFallbackText renders only the display text, discarding the URL.
    HyperlinkFallbackText
    // HyperlinkFallbackURL renders only the URL, discarding the display text.
    HyperlinkFallbackURL
)
```

<a name="Method"></a>

## type [Method](<https://github.com/gechr/x/blob/main/ansi/text.go#L6>)

**Method** controls how display width is calculated.

```go
type Method = xansi.Method
```

<a name="Option"></a>

## type [Option](<https://github.com/gechr/x/blob/main/ansi/options.go#L4>)

**Option** configures an ANSI.

```go
type Option func(*ANSI)
```

<a name="WithHyperlinkFallback"></a>

### func [WithHyperlinkFallback](<https://github.com/gechr/x/blob/main/ansi/hyperlink.go#L20>)

```go
func WithHyperlinkFallback(fallback HyperlinkFallback) Option
```

**WithHyperlinkFallback** sets how hyperlinks render when the output is not a terminal.

<a name="WithTerminal"></a>

### func [WithTerminal](<https://github.com/gechr/x/blob/main/ansi/options.go#L7>)

```go
func WithTerminal(v bool) Option
```

**WithTerminal** sets whether the output target is a terminal.

<a name="WrapOption"></a>

## type [WrapOption](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L84>)

**WrapOption** configures a [Wrapper](<#Wrapper>).

```go
type WrapOption func(*Wrapper)
```

<a name="WithBreakpoints"></a>

### func [WithBreakpoints](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L115>)

```go
func WithBreakpoints(chars string) WrapOption
```

**WithBreakpoints** adds characters (beyond spaces) that are treated as word break opportunities during soft wrapping. Has no effect in hard wrap mode.

<a name="WithPreserveStyle"></a>

### func [WithPreserveStyle](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L121>)

```go
func WithPreserveStyle(preserve bool) WrapOption
```

**WithPreserveStyle** controls whether ANSI styles and hyperlinks are reset and reapplied across line breaks. Default: true.

<a name="WithWidth"></a>

### func [WithWidth](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L87>)

```go
func WithWidth(width int) WrapOption
```

**WithWidth** sets a static wrap width. A `width` \< 1 disables wrapping.

<a name="WithWidthFunc"></a>

### func [WithWidthFunc](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L96>)

```go
func WithWidthFunc(fn func() int) WrapOption
```

**WithWidthFunc** sets a dynamic width function, called on each [Wrapper.Wrap](<#Wrapper.Wrap>) invocation. Takes precedence over [WithWidth](<#WithWidth>).

<a name="WithWrapHard"></a>

### func [WithWrapHard](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L108>)

```go
func WithWrapHard() WrapOption
```

**WithWrapHard** selects hard wrapping: break at the exact column width, even mid-word.

<a name="WithWrapSoft"></a>

### func [WithWrapSoft](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L102>)

```go
func WithWrapSoft() WrapOption
```

**WithWrapSoft** selects soft wrapping: break at space boundaries, with hard-wrap fallback for words longer than the width. This is the default.

<a name="Wrapper"></a>

## type [Wrapper](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L26-L33>)

**Wrapper** wraps text to a configured width, preserving ANSI escape sequences.

```go
type Wrapper struct {
    // contains filtered or unexported fields
}
```

<a name="NewWrapper"></a>

### func [NewWrapper](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L37>)

```go
func NewWrapper(opts ...WrapOption) *Wrapper
```

**NewWrapper** creates a [Wrapper](<#Wrapper>) with the given options. Defaults: soft wrap, no additional breakpoints, ANSI style preservation enabled.

<a name="Wrapper.Wrap"></a>

### func (\*Wrapper) [Wrap](<https://github.com/gechr/x/blob/main/ansi/wrap.go#L50>)

```go
func (w *Wrapper) Wrap(s string) string
```

**Wrap** wraps `s` according to the configured mode and width. Returns `s` unchanged if the effective width is \< 1.
