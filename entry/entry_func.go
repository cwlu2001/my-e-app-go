package entry

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cwlu2001/my-e-app-go/errorHandler"
)

var USER_ID string

func printBanner() {
	banner := `
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⠛⠛⠲⢦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣄⠀⠀⠀⠈⠳⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⠴⠒⠒⠒⠒⠒⠻⢦⡀⠀⠀⠀⠈⠻⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢧⣀⠀⠀⠀⠀⠀⠀⣀⣙⣷⣤⣀⣀⣀⣈⣳⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⣳⣶⡶⠟⠛⠋⠉⠉⠉⠉⠉⠉⠉⠛⠻⠷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⣦⡀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣦⡀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣄⠀⠀⠀
	⠀⠀⠀⠀⣀⠀⠀⢠⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⡆⠀⠀
	⢀⡀⠀⣾⣿⣆⢠⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣶⣦⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣴⣿⡄⠀
	⣾⣿⣶⣽⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣷⣦⣤⣀⣠⣤⣶⣾⣿⣿⣿⣿⣿⣧⠀
	⠿⠿⠿⣿⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡽⠉⠛⠛⢿⣿⣿⣿⣿⣿⡿⠿⠟⠻⡟⠉⠉⣿⠀
	⠀⠀⣾⣿⡿⢿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⣿⣿⠂⣹⡵⣿⣷⠀⠀⠀⣽⠀⠀⣿⡇
	⠀⠀⠈⠋⠀⢸⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢳⡀⠀⠀⢀⣠⠞⠋⠛⠿⣄⡀⢀⡼⠃⠀⠀⣿⠁
	⠀⠀⠀⠀⠀⠸⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢙⡶⠋⠀⠀⠀⠀⠀⠈⠙⢧⡀⠀⠀⢠⡿⠀
	⠀⠀⠀⠀⠀⠀⢻⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣯⣀⡀⠀⠀⠀⠀⠀⠀⠀⠈⢳⡄⠀⣼⠃⠀
	⠀⠀⠀⠀⠀⠀⠈⢿⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢯⡉⠙⠓⠒⠒⠒⠶⢶⡶⠶⠿⣾⠏⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠈⠻⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣄⠀⠀⠀⣠⠴⠋⠀⢠⣾⠋⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢳⡤⠞⠁⠀⣀⣴⠟⠁⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠷⣦⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⠾⠋⠁⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⠷⢶⣤⣤⣤⣤⣤⣤⣤⡴⠶⠟⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	`
	fmt.Println(banner)
}

// Return USER ID from the environment variables
func getUserID() (uid string) {
	uid = os.Getenv(USER_ID)
	uid = strings.ReplaceAll(uid, "\"", "")
	uid = strings.ReplaceAll(uid, "'", "")
	uid = strings.TrimSuffix(uid, "\n")
	if uid == "" {
		err := fmt.Sprintf("❗ %s is not set or is an empty string", USER_ID)
		errorHandler.Handler(err)
	}
	return uid
}

// Write the string to the provided path, set the file permission to 666
func writeStringToFile(s string, path string) {
	err := os.WriteFile(path, []byte(s), 0666)
	errorHandler.Handler(err)
}

func exeCommand(name string, arg ...string) (cmd *exec.Cmd) {
	cmd = exec.Command(name, arg...)

	// Proxy the subcommand's stdin, stdout & stderr to the parent.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	errorHandler.Handler(err)
	return cmd
}
