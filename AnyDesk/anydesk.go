package anydesk

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// https://support.anydesk.com/Automatic_Deployment

const anyDeskDownLink = "https://download.anydesk.com/AnyDesk-CM.exe"
const anyDeskPath = ".\\Vendor\\AnyDesk.exe"

var cmd = &exec.Cmd{
	Path: anyDeskPath,
	Args: []string{"--plain", "--plain"},
}

func Start() {

	err := checkFile(anyDeskPath)
	if err != nil {
		log.Fatal(err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}

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
		// Schrodinger: file may or may not exist. See err for details.

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
