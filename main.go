package main

import (
    "fmt"
    "os"
    "flag"
    "strings"
    "strconv"
    "syscall"
)

import "github.com/cfrco/goportal/core"

const (
    POSTFIX = ".goportal"
    ROOT_DIR_NAME = ".goportal"
)

func fifoPath(filename string) string {
    if filename[0] == '/'{
        return filename
    }
    home_path := os.Getenv("HOME")
    return home_path+"/"+ROOT_DIR_NAME+"/"+filename+POSTFIX
}

var (
    flagReceiver bool
    flagOriginal bool
    flagInternal bool
    narg int
)

func init() {
    flag.BoolVar(&flagReceiver,"r",false,"execute as a receiver.")
    flag.BoolVar(&flagOriginal,"o",false,"original arguments.")
    flag.BoolVar(&flagInternal,"i",false,"send goportal command.")
}

func runReceiver(){
    var fifoName string
    if narg == 0 {
        fifoName = fifoPath(strconv.Itoa(syscall.Getpid()))
    }else if narg == 1 {
        fifoName = fifoPath(flag.Arg(0)) 
    }else {
        return 
    }

    receiver,err := core.NewReceiver(fifoName,core.MODE_RDWR)
    if err != nil {
        fmt.Println(err)
        return 
    }
    defer receiver.Close()

    fmt.Println("goportal receiver have startd. -> "+strconv.Itoa(syscall.Getpid()))
    for {
        fmt.Println(">>>")
        message := strings.Trim(receiver.ReadMessage()," ")
        if message == "" {
            break
        }

        if strings.HasPrefix(message,"\"#cmd:") {
            err := core.RunInternalCommand(message)
            if err != nil {
                if err.Error() == "command:end" {
                    break
                }
                fmt.Println(err)
            }
        }else {
            core.LastRet = core.CallSystem(message)
        }
    }
}

func runSender(){
    if narg < 2 {
        return 
    }else{
        fifoName := fifoPath(flag.Arg(0)) 
        sender,err := core.NewSender(fifoName)
        if err != nil {
            fmt.Println(err)
            return 
        }
        defer sender.Close()
        
        var cmdline string
        if flagOriginal || flagInternal {
            cmdline = strings.Join(flag.Args()[1:]," ")

            if flagInternal {
                cmdline = "\"#cmd:"+cmdline 
            }
        }else {
            cmdline = core.ArgsToCmdline(flag.Args()[1:])
        }
        sender.SendMessage(cmdline)
    }
}

func main() {
    flag.Parse()
    narg = flag.NArg() 

    if flagReceiver{
        runReceiver()
    }else {
        runSender()
    }
}
