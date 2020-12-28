package AnyDesk

import (
	"fmt"
	"io"
	"os/exec"
)

// https://support.anydesk.com/Automatic_Deployment

const anyDeskDownLink = "https://download.anydesk.com/AnyDesk-CM.exe"
const anyDeskPath = ".\\Vendor\\AnyDesk.exe"

var cmd = &exec.Cmd{
	Path: anyDeskPath,
	Args: []string{"--plain", "--plain"},
}

func initAnyDesk() (err error) {

	err = checkFile(anyDeskPath)
	if err != nil {
		return err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)
	return nil
}

// Version Get version of AnyDesk binary
func Version() (err error) {
	cmd.Args = []string{"--plain", "--version"}
	err = initAnyDesk()

	if err != nil {
		return err
	}
	return nil
}
