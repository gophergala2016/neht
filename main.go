package main

import (
	"log"
	"os/exec"
)

func main() {
	var err error
	var stdout []byte
	var ps *exec.Cmd

	ps = exec.Command("ps")
	err = ps.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = ps.Wait()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err = ps.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(stdout)
}
