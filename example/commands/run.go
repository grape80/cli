package commands

import (
	"errors"
	"flag"
	"io"
	"strings"

	"github.com/grape80/cli/example/msgs"
	"github.com/grape80/cli/ui"
)

type Run struct {
	*ui.IO
	// TODO: implement
	Message string
}

func (c *Run) Execute() error {
	// TODO: implement
	switch {
	case strings.HasSuffix(c.Message, "世界！"):
		_, err := c.Println("こんにちは, World!")
		return err
	default:
		return errors.New(msgs.ERR_NO_WORLD)
	}
}

func (c *Run) ParseArgs(args []string) error {
	// TODO: implement
	f := flag.NewFlagSet("run", flag.ContinueOnError)
	f.SetOutput(io.Discard)

	if err := f.Parse(args); err != nil {
		m := strings.Replace(err.Error(), " -", " ", 1) // -option -> option
		return errors.New(m)
	}

	msgs := f.Args()
	c.Message = strings.Join(msgs, " ")

	return nil
}
