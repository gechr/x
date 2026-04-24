package shell_test

import (
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: "''"},
		{name: "safe", in: "foo.example.com", want: "foo.example.com"},
		{name: "safe symbols", in: "a_b@%+=:,./-", want: "a_b@%+=:,./-"},
		{name: "space", in: "no quotes", want: "'no quotes'"},
		{name: "double quotes", in: `"double quoted"`, want: `'"double quoted"'`},
		{name: "single quotes", in: `'single quoted'`, want: `''"'"'single quoted'"'"''`},
		{name: "semicolon", in: ";", want: "';'"},
		{name: "backticks", in: "`echo hello`", want: "'`echo hello`'"},
		{name: "expansion", in: ";${}", want: "';${}'"},
		{name: "newline", in: "line\nbreak", want: "'line\nbreak'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, shell.Quote(tt.in))
		})
	}
}
