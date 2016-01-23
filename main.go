package main

import (
	logging "log"
	"os"
	"os/exec"
)

func main() {
	var err error
	var stdout []byte
	var ps *exec.Cmd = exec.Command("ps", "aux")
	var log *logging.Logger = logging.New(os.Stderr, "",
		logging.LstdFlags|logging.Lshortfile)

	stdout, err = ps.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(stdout))
}
