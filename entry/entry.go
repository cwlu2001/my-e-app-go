package entry

import "fmt"

var (
	uidPath string
	appPath string
)

func Entry() {
	printBanner()

	userID := getUserID()
	writeStringToFile(userID, uidPath)

	// Register device with the userID
	exeCommand(appPath, "start")

	fmt.Println("✓ App is running")
	exeCommand(appPath, "run")
}
