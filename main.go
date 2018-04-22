package main

import (
	"github.com/jwmatthews/svcbndl-cli/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("main() invoked")
	cmd.Execute()
}
