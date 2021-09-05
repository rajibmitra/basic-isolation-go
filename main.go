package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("bad command")
	}
}
func run() {
	fmt.Printf("running %v\n as PID %d\n", os.Args[2:], os.Getgid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot:                     "",
		Credential:                 &syscall.Credential{},
		Ptrace:                     false,
		Setsid:                     false,
		Setpgid:                    false,
		Setctty:                    false,
		Noctty:                     false,
		Ctty:                       0,
		Foreground:                 false,
		Pgid:                       0,
		Pdeathsig:                  0,
		Cloneflags:                 syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		Unshareflags:               0,
		UidMappings:                []syscall.SysProcIDMap{},
		GidMappings:                []syscall.SysProcIDMap{},
		GidMappingsEnableSetgroups: false,
		AmbientCaps:                []uintptr{},
	}
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
