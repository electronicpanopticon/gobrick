package gobrick

import (
	"fmt"
	"os/exec"
	"runtime"
)

/// OpenBrowser opens up a browser in the underlying operating system at the given URL.
func OpenBrowser(url string) {
	execCommand, err := openBrowser(runtime.GOOS, url)
	CheckErr(err)
	err = exec.Command(execCommand.Name, execCommand.Args...).Start()
	CheckErr(err)
}

/// openBrowser separates out the forming of the command from the executing of the command for testing purposes.
func openBrowser(goOs string, url string) (*ExecCommand, error) {
	switch goOs {
	case "linux":
		return &ExecCommand{"xdg-open", []string{url}}, nil
	case "windows":
		return &ExecCommand{"rundll32", []string{"url.dll,FileProtocolHandler", url}}, nil
	case "darwin":
		return &ExecCommand{"open", []string{url}}, nil
	default:
		return &ExecCommand{}, fmt.Errorf("unsupported platform")
	}
}
