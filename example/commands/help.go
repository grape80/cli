package commands

import "github.com/grape80/cli/ui"

type Help struct {
	*ui.IO
	// TODO: implement
	Usage string
}

func (c *Help) Execute() error {
	// TODO: implement
	_, err := c.Print(c.Usage)
	return err
}
