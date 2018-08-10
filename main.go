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
        case "child":
            child()
        default:
            panic("Error!")
    }
}

func run() {
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.SysProcAttr = &syscall.SysProcAttr {
        Cloneflag: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
    }
    must(cmd.Run())
}

func child() {
    fmt.Printf("runnig %v as PID %d\n", os.Args[2:], os.Getpid())
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout

    must(syscall.Chroot("home/rootfs"))
    must(os.Chdir("/"))
    must(syscall.Mount("proc", "proc", "proc", 0, ""))
    
    must(cmd.Run())
}


func must(err error) {
    if err != nil {
        panic(err)
    }
}
