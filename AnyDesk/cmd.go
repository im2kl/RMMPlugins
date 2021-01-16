package AnyDesk

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

var CmdArgs = []string{"--plain"}

//var cmd = exec.Command(anyDeskPath)
var returnedoutput = ""

func callCmd() (err error) {

	err = checkFile(anyDeskPath)
	if err != nil {
		return err
	}

	cmd := exec.Command(anyDeskPath)
	cmd.Args = CmdArgs
	//if runtime.GOOS == "windows" {
	//	cmd = exec.Command("tasklist")
	//}

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		return err
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		//"cmd.Run() failed
		return err
	}

	// Must find better way to return error or return value.
	if errStdout != nil || errStderr != nil {
		//"failed to capture stdout or stderr
		return err
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	returnedoutput = outStr
	return nil
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
