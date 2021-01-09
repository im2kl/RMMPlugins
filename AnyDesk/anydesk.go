package AnyDesk

// https://support.anydesk.com/Automatic_Deployment

const anyDeskDownLink = "https://download.anydesk.com/AnyDesk-CM.exe"
const anyDeskPath = ".\\AnyDesk.exe"

// Version Get version of AnyDesk binary
func Version() (rtn string, err error) {
	//CmdArgs = []string{"--version"}
	CmdArgs = []string{"--version"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetStatus() (rtn string, err error) {
	CmdArgs = []string{"--get-status"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetID() (rtn string, err error) {
	CmdArgs = []string{"--get-id"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func GetAlias() (rtn string, err error) {
	CmdArgs = []string{"--get-alias"}
	err = initAnyDesk()
	if err != nil {
		return "", err
	}
	return returnedoutput, nil
}

func Uninstall() (err error) {
	CmdArgs = []string{"--remove"}
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
		// Better handle the service install
		CmdArgs = []string{"--install", ".", "--start-with-win"}
		err = initAnyDesk()
		if err != nil {
			return err
		}
		//InstallService()
	}
	return nil
}
