package daemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var daemon = flag.Bool("d", false, "Run Quran API as a daemon with -d=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *daemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("Quran API running on [PID]:", cmd.Process.Pid)
		os.Exit(0)
	}
}
