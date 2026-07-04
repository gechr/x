# shell

```go
import "github.com/gechr/x/shell"
```

Package shell provides shell detection, path expansion, argument quoting, and splitting.

## Index

- [Constants](<#constants>)
- [func CacheDir\(\) \(string, error\)](<#CacheDir>)
- [func CompletionFile\(command, sh string\) \(string, error\)](<#CompletionFile>)
- [func ConfigDir\(\) \(string, error\)](<#ConfigDir>)
- [func ConfigDirs\(\) \[\]string](<#ConfigDirs>)
- [func DataDir\(\) \(string, error\)](<#DataDir>)
- [func DataDirs\(\) \[\]string](<#DataDirs>)
- [func Detect\(\) string](<#Detect>)
- [func DetectFromEnv\(env string\) string](<#DetectFromEnv>)
- [func DetectFromProcess\(\) string](<#DetectFromProcess>)
- [func IsKnown\(name string\) bool](<#IsKnown>)
- [func Known\(\) \[\]string](<#Known>)
- [func Quote\(s string\) string](<#Quote>)
- [func Split\(s string\) \(\[\]string, error\)](<#Split>)
- [func StateDir\(\) \(string, error\)](<#StateDir>)

## Constants

<a name="Ash"></a>Recognized shell names, as returned by Known.

```go
const (
    Ash    = "ash"
    Bash   = "bash"
    Dash   = "dash"
    Elvish = "elvish"
    Fish   = "fish"
    Ksh    = "ksh"
    Nu     = "nu"
    Pwsh   = "pwsh"
    Sh     = "sh"
    Tcsh   = "tcsh"
    Zsh    = "zsh"
)
```

<a name="EnvShell"></a>EnvShell is the environment variable consulted by DetectFromEnv.

```go
const EnvShell = "SHELL"
```

<a name="CacheDir"></a>

## func [CacheDir](<https://github.com/gechr/x/blob/main/shell/dir.go#L38>)

```go
func CacheDir() (string, error)
```

CacheDir returns the user cache directory: $XDG\_CACHE\_HOME when set to an absolute path, otherwise an OS\-specific default.

<a name="CompletionFile"></a>

## func [CompletionFile](<https://github.com/gechr/x/blob/main/shell/completion.go#L10>)

```go
func CompletionFile(command, sh string) (string, error)
```

CompletionFile returns the standard completion file path for the given command and shell.

<a name="ConfigDir"></a>

## func [ConfigDir](<https://github.com/gechr/x/blob/main/shell/dir.go#L44>)

```go
func ConfigDir() (string, error)
```

ConfigDir returns the user config directory: $XDG\_CONFIG\_HOME when set to an absolute path, otherwise an OS\-specific default.

<a name="ConfigDirs"></a>

## func [ConfigDirs](<https://github.com/gechr/x/blob/main/shell/dir.go#L64>)

```go
func ConfigDirs() []string
```

ConfigDirs returns the ordered, read\-only config search directories: $XDG\_CONFIG\_DIRS when it has absolute entries, otherwise OS\-specific defaults. These are searched after ConfigDir, so a user's config overrides the system defaults.

<a name="DataDir"></a>

## func [DataDir](<https://github.com/gechr/x/blob/main/shell/dir.go#L50>)

```go
func DataDir() (string, error)
```

DataDir returns the user data directory: $XDG\_DATA\_HOME when set to an absolute path, otherwise an OS\-specific default.

<a name="DataDirs"></a>

## func [DataDirs](<https://github.com/gechr/x/blob/main/shell/dir.go#L72>)

```go
func DataDirs() []string
```

DataDirs returns the ordered, read\-only data search directories: $XDG\_DATA\_DIRS when it has absolute entries, otherwise OS\-specific defaults. These are searched after DataDir, so a user's data overrides the system defaults.

<a name="Detect"></a>

## func [Detect](<https://github.com/gechr/x/blob/main/shell/shell.go#L37>)

```go
func Detect() string
```

Detect returns the shell to use for completions. Priority: COMPLETE\_SHELL env var, parent process name, SHELL env var.

<a name="DetectFromEnv"></a>

## func [DetectFromEnv](<https://github.com/gechr/x/blob/main/shell/shell.go#L13>)

```go
func DetectFromEnv(env string) string
```

DetectFromEnv returns the base name of env if it names a recognized shell.

<a name="DetectFromProcess"></a>

## func [DetectFromProcess](<https://github.com/gechr/x/blob/main/shell/shell.go#L28>)

```go
func DetectFromProcess() string
```

DetectFromProcess returns the parent process name if it is a known shell, or empty if unavailable or not recognized.

<a name="IsKnown"></a>

## func [IsKnown](<https://github.com/gechr/x/blob/main/shell/known.go#L54>)

```go
func IsKnown(name string) bool
```

IsKnown reports whether name matches a known shell.

<a name="Known"></a>

## func [Known](<https://github.com/gechr/x/blob/main/shell/known.go#L49>)

```go
func Known() []string
```

Known returns the set of recognized shell names.

<a name="Quote"></a>

## func [Quote](<https://github.com/gechr/x/blob/main/shell/quote.go#L7>)

```go
func Quote(s string) string
```

Quote returns a shell\-escaped version of s. The returned value can safely be used as one token in a POSIX shell command line.

<a name="Split"></a>

## func [Split](<https://github.com/gechr/x/blob/main/shell/split.go#L27>)

```go
func Split(s string) ([]string, error)
```

Split partitions s into shell\-style words. Whitespace separates words, quotes preserve whitespace, backslashes escape the following rune, a backslash\-newline pair is removed as a line continuation, and a "\#" starts a comment when it appears where a new word could start. Inside double quotes, a backslash is special only before '$', '\`', '"', '\\', or a newline; before any other rune it is kept literally, following POSIX.

<a name="StateDir"></a>

## func [StateDir](<https://github.com/gechr/x/blob/main/shell/dir.go#L56>)

```go
func StateDir() (string, error)
```

StateDir returns the user state directory: $XDG\_STATE\_HOME when set to an absolute path, otherwise an OS\-specific default.
