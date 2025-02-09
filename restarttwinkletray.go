package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// TwinkleTray has great UI and all, but it keeps failing to reconnect to displays after exiting sleep
// When it doot, reboot
func RestartTwinkleTray() {
	fmt.Println("Restarting TwinkleTray")
	cmd := exec.Command("powershell.exe", "-command", `Stop-Process -Name "Twinkle Tray" -Force`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	_ = cmd.Wait()

	cmd = exec.Command("powershell.exe", "-command", `"$((Get-AppxPackage *Twinkle*)[0].InstallLocation)\app\Twinkle Tray.exe"`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	result, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(strings.TrimSpace(string(result)))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Process.Release()
	if err != nil {
		panic(err)
	}
}
