package AnyDesk

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) Install() error {

	logger.Info("Service is being installed.")
	err := p.Install()
	if err != nil {
		println(err.Error())
	}
	return nil
}

func (p *program) run() error {
	p.Install()
	logger.Infof("I'm running %v.", service.Platform())
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			logger.Infof("Still running at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			return nil
		}
	}
}
func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

// Service setup.
//   Define service config.
//   Create the service.
//   Setup the logger.
//   Handle service controls (optional).
//   Run the service.
func InstallService() {
	//svcFlag := flag.String("service", "", "Control the system service.")
	//flag.Parse()
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path += "\\AnyDesk.exe"
	//fmt.Println(path)

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "PNXBL Remote Control Service",
		DisplayName: "PNXBL Remote Control Service",
		Description: "PNXBL Remote control agent service for desktop access.",
		Executable:  path,
		Option:      options,
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	//if len(*svcFlag) != 0 {
	//	err := service.Control(s, *svcFlag)
	//	if err != nil {
	//		log.Printf("Valid actions: %q\n", service.ControlAction)
	//		log.Fatal(err)
	//	}
	//	return
	//}

	err = s.Uninstall()
	if err != nil {
		println(err.Error())
	}

	//err = s.Run()
	//if err != nil {
	//	logger.Error(err)
	//}
}
