package commands

import "github.com/grape80/cli/ui"

type Version struct {
	*ui.IO
	// TODO: implement
	Version string
}

func (c *Version) Execute() error {
	// TODO: implement
	_, err := c.Println(c.Version)
	return err
}
