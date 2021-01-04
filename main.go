package main

import (
	"github.com/im2kl/RMMPlugins/AnyDesk"
)

func main() {
	//AnyDesk.Version()

	err := AnyDesk.Install()
	if err != nil {
		println(err.Error())
	}

	//AnyDesk.Version()

}
