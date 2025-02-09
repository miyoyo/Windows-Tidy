package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/go-toast/toast"
)

func CleanDotfiles() {
	os.Chdir(`C:\Users\Hidden`)
	entries, err := os.ReadDir(`C:\Users\Hidden\`)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, entry := range entries {

		if entry.Name() == "Searches" {
			absPath, err := filepath.Abs(entry.Name())
			if err != nil {
				panic(err)
			}
			os.RemoveAll(absPath)
			continue
		}

		if entry.Name()[0] != '.' {
			continue
		}

		i, _ := entry.Info()
		q := i.Sys().(*syscall.Win32FileAttributeData)

		if q.FileAttributes&0x00000400 == 0 {
			count++
			absolutePath, err := filepath.Abs(entry.Name())
			if err != nil {
				panic(err)
			}

			moveCommand := exec.Command("cmd", "/C", "move", absolutePath, `C:\Users\Hidden\AppData\DotFiles\`+entry.Name())
			err = moveCommand.Run()
			if err != nil {
				panic(err)
			}

			if q.FileAttributes&0x00000010 == 0 {
				mklinkCommand := exec.Command("cmd", "/C", "mklink", absolutePath, `C:\Users\Hidden\AppData\DotFiles\`+entry.Name())
				err = mklinkCommand.Run()
				if err != nil {
					panic(err)
				}
			} else {
				mklinkCommand := exec.Command("cmd", "/C", "mklink", "/J", absolutePath, `C:\Users\Hidden\AppData\DotFiles\`+entry.Name())
				err = mklinkCommand.Run()
				if err != nil {
					panic(err)
				}
			}

			attribCommand := exec.Command("cmd", "/C", "attrib", "/L", "+a", "+h", "+s", absolutePath)
			err = attribCommand.Run()
			if err != nil {
				panic(err)
			}
		}
	}

	if count == 0 {
		return
	}

	notification := toast.Notification{
		AppID:    "Windows Tidy",
		Title:    "DotFiles Cleaner",
		Message:  fmt.Sprintf("%d DotFiles have been cleaned up!", count),
		Icon:     `C:\Users\Hidden\AppData\Scripts\windows-tidy.png`,
		Duration: toast.Short,
	}

	notification.Push()
}
