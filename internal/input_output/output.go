package inputoutput

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("Available Commands:")
	PrintScan()
	PrintAutoScan()
	fmt.Println("Help:")
	fmt.Println("- help")
	fmt.Println("     Displays this help message")
	PrintHistory()
	PrintIgnore()
	PrintReset()
	PrintServer()
	fmt.Println("Quit:")
	fmt.Println("- quit or exit")
	fmt.Println("     Exits the program")
}

func PrintHistory() {
	fmt.Println("History usage:")
	fmt.Println("- history")
	fmt.Println("     List all history entries")
	fmt.Println("- history ./folder/example.txt")
	fmt.Println("     List the history for a specific file")
}

func PrintScan() {
	fmt.Println("Scan usage:")
	fmt.Println("- scan file ./folder/example.txt")
	fmt.Println("     Scans a single file to calculate it's hash and stores it in the database")
	fmt.Println("- scan directory ./folder/")
	fmt.Println("     Scans the target directory storing all file hashes into the database")
}

func PrintAutoScan() {
	fmt.Println("Autoscan usage:")
	fmt.Println("- autoscan start 10 ./folder/")
	fmt.Println("     Scans a directory every 10(customizable) seconds to calculate it's hash and stores it in the database")
	fmt.Println("- autoscan stop ./folder/")
	fmt.Println("     Stops the scan on the target directory")
}

func PrintIgnore() {
	fmt.Println("Ignore usage: ")
	fmt.Println("- ignore add ./folder/example.txt")
	fmt.Println("     Adds the specified file to the ignore list")
	fmt.Println("- ignore remove ./folder/example.txt")
	fmt.Println("     Removed the specified file from the ignore list")
	fmt.Println("- ignore list")
	fmt.Println("     Shows all items from the ignore list")
}

func PrintReset() {
	fmt.Println("Reset usage:")
	fmt.Println("- reset")
	fmt.Println("     Resets the database")
}

func PrintServer() {
	fmt.Println("Server usage:")
	fmt.Println("- server start")
	fmt.Println("     Launches a local webserver which host a page at http://localhost:8080/report")
	fmt.Println("     The report displays the history in a table format")
	fmt.Println("- server stop")
	fmt.Println("     Stops the local server")
}
