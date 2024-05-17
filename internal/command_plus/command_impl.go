package command_plus

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func (c *CommandEntry) IsCliAvailable() bool {
	_, err := exec.LookPath(c.cliName)
	return err == nil
}

func (c *CommandEntry) Exec() error {

	c.mutex.Lock()

	if !c.IsCliAvailable() {
		return fmt.Errorf("cli not found: %s", c.cliName)
	}

	c.stdCode = 0
	c.stdOut = ""
	c.stdErr = ""

	cmd := exec.Command(c.cliName)
	if c.runPath != "" {
		if !pathExistsFast(c.runPath) {
			return fmt.Errorf("run exec path not exists: %s", c.runPath)
		}
		cmd.Dir = c.runPath
	}
	if len(c.commandArgs) > 0 {
		cmd.Args = append(cmd.Args, c.commandArgs...)
	}

	var outStdOut bytes.Buffer
	cmd.Stdout = &outStdOut
	var outStdErr bytes.Buffer
	cmd.Stderr = &outStdErr

	errRun := cmd.Run()
	exitCode := 0
	var exitError *exec.ExitError
	if errors.As(errRun, &exitError) {
		// 仅当命令执行失败时，才会有ExitError
		exitCode = exitError.ExitCode()
	}

	c.stdCode = exitCode
	c.stdErr = outStdErr.String()
	c.stdOut = outStdOut.String()
	c.mutex.Unlock()
	return nil
}

func (c *CommandEntry) ExecCode() int {
	return c.stdCode
}

func (c *CommandEntry) ExecStdOut() string {
	return c.stdOut
}

func (c *CommandEntry) ExecStdErr() string {
	return c.stdErr
}

// pathExistsFast
//
//	path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// pathExists
//
//	path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
