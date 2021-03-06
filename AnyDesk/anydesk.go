package AnyDesk

// https://support.anydesk.com/Automatic_Deployment

const anyDeskDownLink = "https://download.anydesk.com/AnyDesk-CM.exe"
const anyDeskPath = ".\\AnyDesk.exe"

var CmdArgs = []string{"--plain"}

// Version Get version of AnyDesk binary
func Version() (rtn string, err error) {
	//CmdArgs = []string{"--version"}
	CmdArgs = []string{"--version"}
	err = callCmd()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetStatus() (rtn string, err error) {
	CmdArgs = []string{"--get-status"}
	err = callCmd()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetID() (rtn string, err error) {
	CmdArgs = []string{"--get-id"}
	err = callCmd()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetAlias() (rtn string, err error) {
	CmdArgs = []string{"--get-alias"}
	err = callCmd()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func Uninstall() (err error) {
	CmdArgs = []string{"--remove"}
	err = callCmd()
	if err != nil {
		return err
	}
	return nil
}

func Install() (err error) {
	t, _ := CheckService()

	//time.Sleep(2 * time.Second)
	if t != true {
		// Better handle the service install
		CmdArgs = []string{"--install", anyDeskPath, "--start-with-win"}
		err = callCmd()
		if err != nil {
			return err
		}
		//InstallService()
	}
	return nil
}
