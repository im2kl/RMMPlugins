package AnyDesk

import (
	"os/exec"
	"time"
)

// https://support.anydesk.com/Automatic_Deployment

const anyDeskDownLink = "https://download.anydesk.com/AnyDesk-CM.exe"
const anyDeskPath = ".\\AnyDesk.exe"

var cmd = &exec.Cmd{
	Path: anyDeskPath,
	Args: []string{anyDeskPath, "--plain"},
}

var returnedoutput = ""

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
		//io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	returnedoutput = string(out)
	time.Sleep(2 * time.Second)
	//fmt.Printf("%s\n", out)
	return nil
}

// Version Get version of AnyDesk binary
func Version() (rtn string, err error) {
	cmd.Args = []string{anyDeskPath, "--version"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetStatus() (rtn string, err error) {
	cmd.Args = []string{anyDeskPath, "--get-status"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetID() (rtn string, err error) {
	cmd.Args = []string{anyDeskPath, "--get-id"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetAlias() (rtn string, err error) {
	cmd.Args = []string{anyDeskPath, "--get-alias"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func Uninstall() (err error) {
	cmd.Args = []string{anyDeskPath, "--remove"}
	err = initAnyDesk()
	if err != nil {
		return err
	}
	return nil
}

func Install() (err error) {
	t, _ := CheckService()

	//time.Sleep(2 * time.Second)
	if t != true {
		cmd.Args = []string{anyDeskPath, "--install . --start-with-win"}
		err = initAnyDesk()
		if err != nil {
			return err
		}
		println("intall now ...")

	}
	return nil
}
