package command_plus

import (
	"context"
	"sync"
)

type CommandEntry struct {
	CommandEntryFunc `json:"-"`

	ctx   context.Context
	mutex *sync.RWMutex

	runPath string

	cliName     string
	commandArgs []string
	stdErr      string
	stdOut      string
	stdCode     int
}

type CommandEntryFunc interface {
	IsCliAvailable() bool

	Exec() error

	ExecCode() int

	ExecStdOut() string

	ExecStdErr() string
}

var (
	defaultOptionCommandEntry = setDefaultOptionCommandEntry()
)

func setDefaultOptionCommandEntry() *CommandEntry {
	return &CommandEntry{
		ctx:     context.Background(),
		mutex:   &sync.RWMutex{},
		runPath: "",
	}
}

// NewCmd
//
//	use as
//
//	changeCommandEntry := NewCmd(
//	    WithRunPath(""),
//	)
func NewCmd(cliName string, opts ...CommandEntryOption) (opt *CommandEntry) {
	opt = defaultOptionCommandEntry
	for _, o := range opts {
		o(opt)
	}
	defaultOptionCommandEntry = setDefaultOptionCommandEntry()
	opt.cliName = cliName
	return
}

func WithRunPath(runPath string) CommandEntryOption {
	return func(o *CommandEntry) {
		o.runPath = runPath
	}
}

func WithCommandArgs(commandArgs []string) CommandEntryOption {
	return func(o *CommandEntry) {
		o.commandArgs = commandArgs
	}
}

type CommandEntryOption func(*CommandEntry)
