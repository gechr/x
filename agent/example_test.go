package agent_test

import (
	"fmt"

	"github.com/gechr/x/agent"
)

func ExampleKnown() {
	fmt.Println(agent.Known())
	// Output:
	// [amp claude cline codex cursor gemini goose opencode replit]
}

func ExampleIsKnown() {
	fmt.Println(agent.IsKnown(agent.Claude))
	fmt.Println(agent.IsKnown("some-future-agent"))
	// Output:
	// true
	// false
}
