package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fname := "/etc/hosts"
	fname_bk := "/etc/hosts.bk"
	fname_wl := "/etc/hosts.whitelist"

	// create /etc/hosts.bk as a copy of /etc/host if it does not exists
	finfo, err := os.Stat( fname_bk )
	if err != nil {
		if errors.Is( err, os.ErrNotExist ) {
			// make the backup
			fmt.Printf("Making [%s] backup ... ", fname)

			fh_src, err := os.Open( fname )
			exit_if_error(err)
			defer fh_src.Close()

			fh_dst, err := os.Create( fname_bk )
			exit_if_error(err)
			defer fh_dst.Close()

			bytesW, err := io.Copy( fh_dst, fh_src )
			exit_if_error(err)
			err = fh_dst.Sync()
			exit_if_error(err)

			fmt.Printf("OK, %d bytes copied.",bytesW)
			fh_dst.Close()
			fh_src.Close()
		} else {
			fmt.Println("ERROR:")
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if finfo.IsDir() {
			fmt.Printf("ERROR: Backup file [%s] exist but is a directory.\n",fname_bk)
			os.Exit(1)
		}
	}

	// eval what is requested from args
	argc := len(os.Args)
	if argc == 1 {
		fmt.Println("======================================================================================")
		fmt.Println("HELP:")
		fmt.Println("")
		fmt.Println("This program allows you to modify the hostname resolution by modifying the hosts file.")
		fmt.Println("")
		fmt.Println("Operation modes:")
		fmt.Println("")
		fmt.Println("hostname-modify restore")
		fmt.Println("    Restore the hostfile to its original, this is a snapshot of the file taken when")
		fmt.Println("    ran this program for first time.")
		fmt.Println("")
		fmt.Println("hostname-modify add $records")
		fmt.Println("    Append $records to hosts file, it is a string with IP name pairs separetd by ;")
		fmt.Println("")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "restore":
		fh_src, err := os.Open( fname_bk )
		exit_if_error(err)
		defer fh_src.Close()

		fh_dst, err := os.OpenFile( fname, os.O_WRONLY | os.O_TRUNC, 0644 )
		exit_if_error(err)
		defer fh_dst.Close()

		bytesW, err := io.Copy( fh_dst, fh_src )
		exit_if_error(err)
		err = fh_dst.Sync()
		exit_if_error(err)
		fmt.Printf( "OK: File %q restored, %d bytes copied.\n", fname, bytesW )

	case "add":
		if argc != 3 {
			fmt.Printf("HELP: You shold provide the IP FQDN pairs separated by ; after the add command.\n")
			os.Exit(1)
		}
		// load whitelist
		allowedH := make( map[string]int )
		fh_wl, err := os.Open( fname_wl )
		exit_if_error(err)
		defer fh_wl.Close()
		scanner := bufio.NewScanner( fh_wl )
		for scanner.Scan() {
			line := scanner.Text()
			allowedH[ line ]++
		}
		scan_err := scanner.Err()
		exit_if_error(scan_err)

		// clear records and check they are in whitelist
		records := os.Args[2]
		recordA := strings.Split( records, ";" )
		recordH := make( map[string]bool )
		for _, record := range recordA {
			record = strings.TrimSpace(record)
			if len(record) == 0 {
				continue
			}
			if allowedH[ record ] > 0 {
				recordH[ record ] = true
			} else {
				fmt.Printf("ERROR: [%s] not in whitelist.\n", record)
				os.Exit(1)
			}
		}
		// fh for append
		fh, err := os.OpenFile( fname, os.O_WRONLY | os.O_APPEND , 0644 )
		exit_if_error(err)
		defer fh.Close()
		// process records
		for k, _ := range recordH {
			fmt.Printf("hostname-modify adding [%s] ", k)
			_, err := fh.WriteString( k + "\n" )
			exit_if_error(err)
			fmt.Printf("OK\n")
		}

	default:
		fmt.Printf("ERROR: %q command not implemented.\n", command)
		os.Exit(1)
	}
}

func exit_if_error(err error) {
	if err != nil {
		fmt.Println("\nUnhandled ERROR:")
		fmt.Println(err)
		os.Exit(255)
	}
}

