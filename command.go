package gobrick

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

/// ExecCommand is a struct used to hold the command values passed to exec.Command()
type ExecCommand struct {
	Name string
	Args []string
}

type Commander interface {
	CombinedOutput(string, ...string) ([]byte, error)
}

type TestCommander struct{}

func (c TestCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	cs := []string{"-test.run=TestOutput", "--"}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_TEST_OUTPUT=1"}
	out, err := cmd.CombinedOutput()
	return out, err
}

func HelperCommandContext(t *testing.T, ctx context.Context, s ...string) (cmd *exec.Cmd) {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, s...)
	if ctx != nil {
		cmd = exec.CommandContext(ctx, os.Args[0], cs...)
	} else {
		cmd = exec.Command(os.Args[0], cs...)
	}
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	return cmd
}

func HelperCommand(t *testing.T, s ...string) *exec.Cmd {
	return HelperCommandContext(t, nil, s...)
}