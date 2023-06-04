package example

import (
	"errors"
	"os"
	"strings"

	"github.com/grape80/cli"
	"github.com/grape80/cli/example/commands"
	"github.com/grape80/cli/example/msgs"
	"github.com/grape80/cli/ui"
)

type App struct {
	*ui.IO
	Name    string
	Version string
	Usage   string
	Args    []string
}

func New() *App {
	// TODO: implement
	return &App{
		&ui.IO{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		},
		os.Args[0],
		"",
		usage,
		os.Args[1:],
	}
}

func (a *App) Run() (exitStatus int) {
	// TODO: implement
	if err := a.run(); err != nil {
		a.Eprintln(err.Error())
		return EXIT_FAILURE
	}

	return EXIT_SUCCESS
}

func (a *App) run() error {
	// TODO: implement
	cmd, err := a.newCommand()
	if err != nil {
		return err
	}

	if err := cmd.Execute(); err != nil {
		return err
	}

	return nil
}

func (a *App) newCommand() (cli.Command, error) {
	// TODO: implement
	if err := a.parseArgs(); err != nil {
		return nil, err
	}

	s := strings.Join(a.Args, " ")
	switch {
	case strings.HasPrefix(s, "version"):
		return &commands.Version{
			IO:      a.IO,
			Version: a.Version,
		}, nil
	case strings.HasPrefix(s, "help"):
		return &commands.Help{
			IO:    a.IO,
			Usage: a.Usage,
		}, nil
	default:
		cmd := &commands.Run{
			IO: a.IO,
		}
		err := cmd.ParseArgs(a.Args)
		return cmd, err
	}
}

func (a *App) parseArgs() error {
	// TODO: implement
	if len(a.Args) == 0 {
		return errors.New(msgs.ERR_NO_ARGS)
	}

	return nil
}
