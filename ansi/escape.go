package ansi

import xansi "github.com/charmbracelet/x/ansi"

// CSI is the Control Sequence Introducer prefix (ESC + '['). Exported for
// callers that parse or construct their own CSI-family sequences.
const CSI = "\x1b["

// ClearLine erases the entire current line (EL 2) and returns the cursor to
// column 0 (CR). This is a convenience composite not present in upstream;
// every other escape here delegates to [github.com/charmbracelet/x/ansi].
const ClearLine = xansi.EraseEntireLine + "\r"

// Erase constants (delegated).
const (
	EraseLineRight    = xansi.EraseLineRight
	EraseLineLeft     = xansi.EraseLineLeft
	EraseEntireLine   = xansi.EraseEntireLine
	EraseScreenBelow  = xansi.EraseScreenBelow
	EraseScreenAbove  = xansi.EraseScreenAbove
	EraseEntireScreen = xansi.EraseEntireScreen
)

// EraseLine returns the EL sequence. n selects the variant:
// 0 = right of cursor, 1 = left of cursor, 2 = entire line.
func EraseLine(n int) string { return xansi.EraseLine(n) }

// EraseDisplay returns the ED sequence. n selects the variant:
// 0 = below cursor, 1 = above cursor, 2 = entire screen.
func EraseDisplay(n int) string { return xansi.EraseDisplay(n) }

// EraseCharacter returns the ECH sequence: erase n characters from the
// cursor position (no cursor movement).
func EraseCharacter(n int) string { return xansi.EraseCharacter(n) }

// Cursor movement — single-step constants. Upstream's preferred names are
// the opaque CUU1/CUD1/CUF1/CUB1 spellings; we keep the readable aliases
// here and forward to the non-deprecated upstream symbols.
const (
	CursorUp1    = xansi.CUU1
	CursorDown1  = xansi.CUD1
	CursorRight1 = xansi.CUF1
	CursorLeft1  = xansi.CUB1
)

// CursorUp returns the CUU sequence: move cursor up n lines.
func CursorUp(n int) string { return xansi.CursorUp(n) }

// CursorDown returns the CUD sequence: move cursor down n lines.
func CursorDown(n int) string { return xansi.CursorDown(n) }

// CursorForward returns the CUF sequence: move cursor right n columns.
func CursorForward(n int) string { return xansi.CursorForward(n) }

// CursorBackward returns the CUB sequence: move cursor left n columns.
func CursorBackward(n int) string { return xansi.CursorBackward(n) }

// CursorNextLine returns the CNL sequence: move down n lines and to column 1.
func CursorNextLine(n int) string { return xansi.CursorNextLine(n) }

// CursorPreviousLine returns the CPL sequence: move up n lines and to column 1.
func CursorPreviousLine(n int) string { return xansi.CursorPreviousLine(n) }

// CursorHorizontalAbsolute returns the CHA sequence: move to column col
// on the current line.
func CursorHorizontalAbsolute(col int) string { return xansi.CursorHorizontalAbsolute(col) }

// CursorPosition returns the CUP sequence: move to (col, row). Coordinates
// are 1-based.
func CursorPosition(col, row int) string { return xansi.CursorPosition(col, row) }

// CursorHomePosition moves the cursor to row 1, column 1 ("\x1b[H").
const CursorHomePosition = xansi.CursorHomePosition

// Cursor visibility (DECTCEM).
const (
	ShowCursor = xansi.ShowCursor
	HideCursor = xansi.HideCursor
)

// Cursor save/restore (SCOSC / SCORC).
const (
	SaveCursorPosition    = xansi.SaveCurrentCursorPosition
	RestoreCursorPosition = xansi.RestoreCurrentCursorPosition
)

// Focus reporting events (terminal → app, when focus events are enabled).
const (
	Focus = xansi.Focus
	Blur  = xansi.Blur
)

// Alt-screen buffer (DEC 1049: also saves/restores cursor).
const (
	EnterAltScreen = xansi.SetModeAltScreenSaveCursor
	ExitAltScreen  = xansi.ResetModeAltScreenSaveCursor
)

// Bracketed paste mode (DEC 2004).
const (
	EnableBracketedPaste  = xansi.SetModeBracketedPaste
	DisableBracketedPaste = xansi.ResetModeBracketedPaste
)

// Focus event reporting mode (DEC 1004).
const (
	EnableFocusEvents  = xansi.SetModeFocusEvent
	DisableFocusEvents = xansi.ResetModeFocusEvent
)

// Terminal queries — the terminal replies with a corresponding report.
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

// Scrolling.

// ScrollUp returns the SU sequence: scroll viewport up n lines
// (content moves up; new blank lines appear at the bottom).
func ScrollUp(n int) string { return xansi.ScrollUp(n) }

// ScrollDown returns the SD sequence: scroll viewport down n lines
// (content moves down; new blank lines appear at the top).
func ScrollDown(n int) string { return xansi.ScrollDown(n) }

// Line and character insert/delete at the cursor.

// InsertLine returns the IL sequence: insert n blank lines at the cursor,
// pushing existing lines down.
func InsertLine(n int) string { return xansi.InsertLine(n) }

// DeleteLine returns the DL sequence: delete n lines starting at the
// cursor, pulling subsequent lines up.
func DeleteLine(n int) string { return xansi.DeleteLine(n) }

// InsertCharacter returns the ICH sequence: insert n blank characters
// at the cursor, shifting existing characters right.
func InsertCharacter(n int) string { return xansi.InsertCharacter(n) }

// DeleteCharacter returns the DCH sequence: delete n characters at the
// cursor, pulling subsequent characters left.
func DeleteCharacter(n int) string { return xansi.DeleteCharacter(n) }

// ResetStyle is the SGR reset sequence ("\x1b[m"): clears all text
// attributes and colours. Included because writing any styled output
// typically requires emitting this afterwards.
const ResetStyle = xansi.ResetStyle

// SetCursorStyle returns the DECSCUSR sequence. Style selects shape and
// blink state:
//
//	0: default         1: blinking block    2: steady block
//	3: blinking under  4: steady under      5: blinking bar    6: steady bar
func SetCursorStyle(style int) string { return xansi.SetCursorStyle(style) }

// Terminal title (OSC 0/1/2). These emit OSC sequences terminated by BEL.

// SetWindowTitle sets the terminal window title only (OSC 2).
func SetWindowTitle(s string) string { return xansi.SetWindowTitle(s) }

// SetIconName sets the icon/tab name only (OSC 1).
func SetIconName(s string) string { return xansi.SetIconName(s) }

// SetIconNameWindowTitle sets both the icon name and window title in a
// single sequence (OSC 0).
func SetIconNameWindowTitle(s string) string { return xansi.SetIconNameWindowTitle(s) }

// Keypad modes (DECKPAM / DECKPNM). Application mode makes the numeric
// keypad emit escape sequences rather than raw digits; relevant when the
// host reads keys in raw mode.
const (
	KeypadApplicationMode = xansi.KeypadApplicationMode
	KeypadNumericMode     = xansi.KeypadNumericMode
)

// RequestNameVersion is the XTVERSION query ("\x1b[>q"): asks the terminal
// for its name and version. Complements the DA1 query.
const RequestNameVersion = xansi.RequestNameVersion

// Mouse tracking modes. Each pair toggles a DEC private mode; the encoding
// mode (SGR) is usually enabled alongside one of the event modes:
//
//	X10         legacy, presses only (no release)
//	Normal      press + release
//	ButtonEvent press + release + motion while button held
//	AnyEvent    press + release + motion at all times
//	SGR         extended coordinate encoding, required for columns > 223
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
