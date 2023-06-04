package example_test

import (
	"bytes"
	"testing"

	. "github.com/grape80/cli/example"
)

func TestApp_Run(t *testing.T) {
	tests := map[string]struct {
		exitCodeExpected int
		args             []string
	}{
		"no_args": {EXIT_FAILURE, []string{}},

		// run
		"run":                {EXIT_SUCCESS, []string{"Hello, 世界！"}},
		"run_error":          {EXIT_FAILURE, []string{"Hello"}},
		"run_invalid_option": {EXIT_FAILURE, []string{"--option=invalid"}},

		// help
		"help": {EXIT_SUCCESS, []string{"help"}},

		// version
		"version": {EXIT_SUCCESS, []string{"version"}},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(tt *testing.T) {
			app := New()
			app.Stdin = new(bytes.Buffer)
			app.Stdout = new(bytes.Buffer)
			app.Stderr = new(bytes.Buffer)
			app.Name = "yourapp"
			app.Args = test.args

			exitCode := app.Run()
			if exitCode != test.exitCodeExpected {
				t.Errorf("Run() = %v; want %v", exitCode, test.exitCodeExpected)
			}
		})
	}
}

func BenchmarkApp_Run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := New()
		app.Stdin = new(bytes.Buffer)
		app.Stdout = new(bytes.Buffer)
		app.Stderr = new(bytes.Buffer)
		app.Name = "yourapp"
		app.Args = []string{"Hello, 世界！"}

		app.Run()
	}
}
