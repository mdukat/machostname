package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"
)

func main() {
	// Try local directory
	readConfigFile, err := os.Open("machostname.conf")
	if err != nil {
		fmt.Println("Configuration not found in current directory", err)
	} else {
		fmt.Println("Found configuration in current directory")
		goto POSTREADCONFIG
	}

	// Try /etc/machostname.conf
	readConfigFile, err = os.Open("/etc/machostname.conf")
	if err != nil {
		fmt.Println("Configuration file error", err)
		os.Exit(1)
	}
	fmt.Println("Using /etc/machostname.conf")

POSTREADCONFIG:

	// Get all interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Interfaces error", err)
		os.Exit(1)
	}

	var ifaceSlice []string
	for _, iface := range interfaces {
		ifaceMACString := iface.HardwareAddr.String()

		// Ignore interfaces with empty MAC address
		if ifaceMACString != "" {
			fmt.Println("Found interface", ifaceMACString)
			ifaceSlice = append(ifaceSlice, ifaceMACString)
		}
	}

	configScanner := bufio.NewScanner(readConfigFile)
	configScanner.Split(bufio.ScanLines)

	for configScanner.Scan() {
		currentLine := configScanner.Text()

		// Ignore empty lines
		if currentLine == "" {
			continue
		}

		// Split line into fields
		currentFields := strings.Fields(currentLine)

		// Ignore comments
		if currentFields[0][0:1] == "#" {
			continue
		}

		// Ignore anything else than 2 fields (MAC + hostname)
		if len(currentFields) != 2 {
			continue
		}

		// Parse MAC and hostname
		for _, currentMAC := range ifaceSlice {
			// Ignore if current MAC in list is not what we're looking for
			if currentMAC != currentFields[0] {
				continue
			}

			// Else, use the hostname
			fmt.Println("Found the hostname", currentFields[1], currentFields[0])
			err := syscall.Sethostname([]byte(currentFields[1]))
			if err != nil {
				fmt.Println("Sethostname syscall error", err)
				os.Exit(1)
			}
			fmt.Println("Hostname changed to", currentFields[1])
			os.Exit(0)
		}
	}

	readConfigFile.Close()
	fmt.Println("Any interface MAC address not found in configuration")
}
