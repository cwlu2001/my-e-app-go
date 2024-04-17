package main

import (
	"os"
	"path/filepath"

	"github.com/cwlu2001/my-e-app-go/entry"
	"github.com/cwlu2001/my-e-app-go/myip"
	"github.com/cwlu2001/my-e-app-go/systemInfo"
	"github.com/cwlu2001/my-e-app-go/systemdDetectVirt"
)

func selectExe(args []string) {
	excName := filepath.Base(args[0])
	switch excName {
	case "init":
		entry.Entry()

	case "hostnamectl":
		systemInfo.Hostnamectl()

	case "lsb_release":
		systemInfo.Lsb_release()

	case "myip":
		myip.Myip()

	case "systemd-detect-virt":
		systemdDetectVirt.DetectVirt()

	case "sh":
		if len(args) > 2 && args[1] == "-c" {
			selectExe(args[2:])
		}

	case "systemctl":
		//null function
	default:
		msg := "bash: " + excName + ": command not found\n"
		print(msg)
	}
}

func main() {
	selectExe(os.Args)
}
