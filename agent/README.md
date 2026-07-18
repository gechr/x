# agent

```go
import "github.com/gechr/x/agent"
```

Package `agent` detects AI coding agents hosting the process.

## Index

- [Constants](<#constants>)
- [func Detect() string](<#Detect>)
- [func Is() bool](<#Is>)
- [func IsKnown(name string) bool](<#IsKnown>)
- [func Known() \[\]string](<#Known>)

## Constants

<a name="EnvAgent"></a>Environment variables consulted by [Is](<#Is>) and [Detect](<#Detect>).

```go
const (
    // EnvAgent is the proposed cross-tool convention
    // (https://github.com/agentsmd/agents.md/issues/136), set either to the
    // agent's name (e.g. `goose`, `amp`) or a truthy token like `1`.
    EnvAgent = "AGENT"
    // EnvAIAgent is the more specific variant used by Vercel's detection
    // tooling; consulted with the same semantics as [EnvAgent].
    EnvAIAgent = "AI_AGENT"
)
```

<a name="Amp"></a>Recognized AI coding agent names, as returned by [Detect](<#Detect>).

```go
const (
    Amp      = "amp"
    Claude   = "claude"
    Cline    = "cline"
    Codex    = "codex"
    Cursor   = "cursor"
    Gemini   = "gemini"
    Goose    = "goose"
    OpenCode = "opencode"
    Replit   = "replit"
)
```

<a name="Detect"></a>

## func [Detect](<https://github.com/gechr/x/blob/main/agent/detect.go#L70>)

```go
func Detect() string
```

**Detect** returns the name of the AI coding agent hosting the process, or empty if there is none or it cannot be identified. A named `AGENT` / `AI_AGENT` value is returned verbatim (lowercased), so agents unknown to this package still identify themselves; a bare truthy token like `AGENT=1` signals presence without identity, so Detect falls through to the marker variables and may return empty even though [Is](<#Is>) reports true.

<a name="Is"></a>

## func [Is](<https://github.com/gechr/x/blob/main/agent/detect.go#L50>)

```go
func Is() bool
```

**Is** reports whether the process appears to be running under an AI coding agent. Detection is best-effort, based on environment variables set by the hosting agent: the cross-tool `AGENT` / `AI_AGENT` convention first, then agent-specific markers (`CLAUDECODE`, `CODEX_SANDBOX`, `CURSOR_AGENT`, ...). A falsy `AGENT` / `AI_AGENT` value (`0`, `false`, `no`, `off`) is an explicit opt-out and reports false even when a marker variable is present, mirroring the `CI=false` escape hatch.

<a name="IsKnown"></a>

## func [IsKnown](<https://github.com/gechr/x/blob/main/agent/known.go#L42>)

```go
func IsKnown(name string) bool
```

**IsKnown** reports whether `name` matches a known AI coding agent.

<details><summary><b>Example</b></summary>

```go
fmt.Println(agent.IsKnown(agent.Claude))
fmt.Println(agent.IsKnown("some-future-agent"))
```

Output:

```text
true
false
```

</details>

<a name="Known"></a>

## func [Known](<https://github.com/gechr/x/blob/main/agent/known.go#L37>)

```go
func Known() []string
```

**Known** returns the set of recognized AI coding agent names.

<details><summary><b>Example</b></summary>

```go
fmt.Println(agent.Known())
```

Output:

```text
[amp claude cline codex cursor gemini goose opencode replit]
```

</details>
