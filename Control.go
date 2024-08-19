package main

import (
	"fmt"
	"os"
)

func CheckArg(arg *ArgVar) {
	switch arg.Protocol {
	case "wordpress":
		if arg.URL == "" {
			fmt.Println("[+] Usage: URL for target WordPress site is required.")
			os.Exit(0)
		} else if arg.Address != "" {
			fmt.Println("[+] Usage: IP address is not required for WordPress protocol.")
			fmt.Println("[+] Usage: Provide only the URL for the target WordPress site.")
			os.Exit(0)
		}
	default:
		if arg.Address == "" {
			fmt.Println("[+] Usage: IP address for target machine is required.")
			os.Exit(0)
		}
	}

	if arg.User == "" && arg.UserList == "" {
		fmt.Println("[+] Usage: Provide a username or a file list of usernames.")
		os.Exit(0)
	}
	if arg.PassList == "" && arg.Pass == "" {
		fmt.Println("[+] Usage: Provide a password or a file list of passwords.")
		os.Exit(0)
	}
	switch arg.Protocol {
	case "ssh":
		if arg.Port == "" {
			arg.Port = "22"
		}
	case "ftp":
		if arg.Port == "" {
			arg.Port = "21"
		}
	}
}
