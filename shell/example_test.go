package shell_test

import (
	"fmt"

	"github.com/gechr/x/shell"
)

func ExampleQuote() {
	fmt.Println(shell.Quote("safe-token_1.txt"))
	fmt.Println(shell.Quote("has spaces"))
	fmt.Println(shell.Quote("$HOME"))
	fmt.Println(shell.Quote(""))
	// Output:
	// safe-token_1.txt
	// 'has spaces'
	// '$HOME'
	// ''
}

// Single quotes inside the input are escaped so the result stays one token.
func ExampleQuote_singleQuotes() {
	fmt.Println(shell.Quote("it's fine"))
	// Output:
	// 'it'"'"'s fine'
}

func ExampleSplit() {
	words, err := shell.Split(`cp "my file.txt" backup/ # keep a copy`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", words)
	// Output:
	// ["cp" "my file.txt" "backup/"]
}

// A backslash-newline pair is removed as a line continuation.
func ExampleSplit_lineContinuation() {
	words, err := shell.Split("echo one \\\ntwo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", words)
	// Output:
	// ["echo" "one" "two"]
}

func ExampleSplit_unclosedQuote() {
	_, err := shell.Split(`echo "unterminated`)
	fmt.Println(err)
	// Output:
	// EOF found when expecting closing quote
}

func ExampleKnown() {
	fmt.Println(shell.Known())
	// Output:
	// [ash bash dash elvish fish ksh nu pwsh sh tcsh zsh]
}

func ExampleIsKnown() {
	fmt.Println(shell.IsKnown("zsh"))
	fmt.Println(shell.IsKnown("cmd.exe"))
	// Output:
	// true
	// false
}
