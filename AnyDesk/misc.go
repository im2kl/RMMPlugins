package AnyDesk

import (
	"io"
	"log"
	"net/http"
	"os"
)

func checkFile(filepath string) (err error) {
	if _, err := os.Stat(filepath); err == nil {
		// path/to/whatever exists

	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		err := downloadFile(filepath, anyDeskDownLink)
		if err != nil {
			log.Fatal(err)
			return err
		}

	} else {
		// file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func CheckService() (running bool, err error) {

	cmd.Args = []string{anyDeskPath, "--get-status"}
	err = initAnyDesk()
	if err != nil {
		return false, nil
	}

	if &returnedoutput == nil {
		return false, nil
	}

	//fmt.Printf("%s\n", returnedoutput)
	return true, nil
}
