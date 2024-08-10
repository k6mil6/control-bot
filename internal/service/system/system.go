package system

import "os/exec"

func TurnOff() error {
	err := exec.Command("shutdown", "/s", "/t", "15", "/c", "15 секунд и я отключусь").Run()
	return err
}

func Restart() error {
	err := exec.Command("shutdown", "/r", "/t", "15", "/c", "15 секунд и я перезагружусь").Run()
	return err
}
