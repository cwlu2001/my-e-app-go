package systemInfo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/cwlu2001/my-e-app-go/errorHandler"
)

var hostLists = [][]string{
	{"Ubuntu", "Ubuntu 22.04.3 LTS", "Linux 6.6.23-current", "22.04", "jammy"},
	{"Debian", "Debian GNU/Linux 12 (bookworm)", "Linux 6.1.77-current", "12", "bookworm"},
	{"Linuxmint", "Linux Mint 21.3", "Linux 5.15.0-91-generic", "21.3", "virginia"},
	{"Ubuntu", "Ubuntu 24.04 LTS", "Linux 6.8.0-31-generic", "24.04", "noble"},
}

func readFile(file string) []byte {
	data, err := os.ReadFile(file)
	errorHandler.Handler(err)
	return data
}

func md5sum(data []byte) []byte {
	d := md5.Sum(data)
	return d[:]
}

/* Get a random int based on the md5sum(uid) */
func getIndexFromUID(uid []byte) int {
	digest := md5sum(uid)
	return int(digest[0] & 0b11)
}

func getHostName() (h string) {
	h, _ = os.Hostname()
	return h
}

func getMachineID(uid []byte) string {
	mID := md5sum(uid)
	return hex.EncodeToString(mID)
}

func getBootID(uid []byte) string {
	mID := md5sum(md5sum(uid))
	return hex.EncodeToString(mID)
}

func getDistributor(idx int) string {
	return hostLists[idx][0]
}
func getOS(idx int) string {
	return hostLists[idx][1]
}
func getKernel(idx int) string {
	return hostLists[idx][2]
}
func getReleaseNum(idx int) string {
	return hostLists[idx][3]
}
func getCodename(idx int) string {
	return hostLists[idx][4]
}

func getArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86-64"
	default:
		return runtime.GOARCH
	}
}

func resultHost(uidPath string) []string {
	uid := readFile(uidPath)
	idx := getIndexFromUID(uid)
	return []string{
		fmt.Sprintf("%18s: %s", "Static hostname", getHostName()),
		fmt.Sprintf("%18s: %s", "Icon name", "computer-desktop"),
		fmt.Sprintf("%18s: %s", "Chassis", "desktop"),
		fmt.Sprintf("%18s: %s", "Machine ID", getMachineID(uid)),
		fmt.Sprintf("%18s: %s", "Boot ID", getBootID(uid)),
		fmt.Sprintf("%18s: %s", "Operating System", getOS(idx)),
		fmt.Sprintf("%18s: %s", "Kernel", getKernel(idx)),
		fmt.Sprintf("%18s: %s", "Architecture", getArch()),
	}
}
func resultRelease(uidPath string) []string {
	uid := readFile(uidPath)
	idx := getIndexFromUID(uid)
	return []string{
		"No LSB modules are available.",
		fmt.Sprintf("%-15s %s", "Distributor ID:", getDistributor(idx)),
		fmt.Sprintf("%-15s %s", "Description:", getOS(idx)),
		fmt.Sprintf("%-15s %s", "Release:", getReleaseNum(idx)),
		fmt.Sprintf("%-15s %s", "Codename:", getCodename(idx)),
	}
}
