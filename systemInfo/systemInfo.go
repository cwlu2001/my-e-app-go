package systemInfo

import (
	"fmt"
	"strings"
)

var (
	uidPath string
)

func Hostnamectl() {
	idn := resultHost(uidPath)
	msg := strings.Join(idn, "\n")
	fmt.Println(msg)
	/* Example output
	 Static hostname: debian-11
	       Icon name: computer-desktop
	         Chassis: desktop
	      Machine ID: 7faf76bdf7e550cfe1278a351ec45868
	         Boot ID: c9048c4bdf77a54314322a07d6f479a3
	Operating System: Ubuntu 22.04.3 LTS
	          Kernel: Linux 5.15.130-1-current
	    Architecture: x86-64
	*/
}

func Lsb_release() {
	idn := resultRelease(uidPath)
	msg := strings.Join(idn, "\n")
	fmt.Println(msg)
	/* Example output
	No LSB modules are available.
	Distributor ID: Debian
	Description:    Debian GNU/Linux 11 (bullseye)
	Release:        11
	Codename:       bullseye
	*/
}
