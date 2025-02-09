package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-toast/toast"
	"golang.org/x/sys/windows"
)

func WingetCheck() {
	ownHandle := windows.CurrentProcess()
	err := windows.SetPriorityClass(ownHandle, 0x40)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("winget", "upgrade")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	result, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	if strings.Contains(string(result), "No installed package found matching input criteria.") {
		fmt.Println("No updates found.")
		time.Sleep(3 * time.Second)
	}

	matcher, err := regexp.Compile(`(\d) upgrades available\.`)
	if err != nil {
		panic(err)
	}
	found := matcher.FindSubmatch(result)
	count, err := strconv.Atoi(string(found[1]))
	if err != nil {
		panic(err)
	}

	notification := toast.Notification{
		AppID:               "Windows Tidy",
		Title:               "Winget Upgrades",
		Message:             fmt.Sprintf("%d upgrades are available.", count),
		ActivationType:      "Protocol",
		ActivationArguments: "windowstidy:upgrade",
		Icon:                `C:\Users\Hidden\AppData\Scripts\windows-tidy.png`,
	}

	notification.Push()
}
