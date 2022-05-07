package main

import (
	"strings"
	toplib "toplib/pkg"
)

var w = toplib.NewWorker()

func main() {

	w.AddWork(checker)
	w.StartWorker()
	//Hook into dll
}

func checker() {

	acc := toplib.Account{
		Username: "Fart",
		Password: "Password",
		Capture:  nil,
	}

	if strings.Contains(acc.Username, "Fart") {
		acc.AddCaptureString("Username contains Fart", "True")
		acc.AddHit(true)
	}
}
