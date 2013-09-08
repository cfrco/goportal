package main

import (
    "fmt"
    "os"
    "flag"
    "strings"
    "strconv"
    "syscall"
)

const (
    POSTFIX = ".goportal"
    ROOT_DIR_NAME = ".goportal"
)

func colorize(text string, style string) string {
    return "\x1B["+style+"m"+text+"\x1B[00m";
}

func fifoPath(filename string) string {
    if filename[0] == '/'{
        return filename
    }
    home_path := os.Getenv("HOME")
    return home_path+"/"+ROOT_DIR_NAME+"/"+filename+POSTFIX
}

var (
    // flags
    flagReceiver bool
    flagOriginal bool
    flagInternal bool

    // global
    narg int
)

func init() {
    flag.BoolVar(&flagReceiver,"r",false,"execute as a receiver.")
    flag.BoolVar(&flagOriginal,"o",false,"original arguments.")
    flag.BoolVar(&flagInternal,"i",false,"send goportal command.")
}

func runReceiver(){
    fifoRootPath := os.Getenv("HOME")+"/.goportal"
    if _,err := os.Stat(fifoRootPath);err != nil {
        if err = os.Mkdir(fifoRootPath,0700); err != nil {
            fmt.Println(err) 
            return 
        }
    }

    var fifoName string
    if narg == 0 {
        fifoName = fifoPath(strconv.Itoa(syscall.Getpid()))
    }else if narg == 1 {
        fifoName = fifoPath(flag.Arg(0)) 
    }else {
        return 
    }

    receiver,err := NewReceiver(fifoName,MODE_RDWR)
    if err != nil {
        fmt.Println(err)
        return 
    }
    defer receiver.Close()

    fmt.Print("goportal receiver have startd. -> ")
    if narg == 0 {
        fmt.Println(strconv.Itoa(syscall.Getpid())) 
    }else {
        fmt.Println(flag.Arg(0)) 
    }

    HistoryCmd.Init()

    for {
        fmt.Println(colorize(">>>","1;32"));
        message := strings.Trim(receiver.ReadMessage()," ")

        if strings.HasPrefix(message,"#cmd:") {
            err := RunInternalCommand(message)
            if err != nil {
                if err.Error() == "command:end" {
                    break
                }
                fmt.Println(err)
            }
        }else {
            if message != ""{
                LastRet = CallSystem(message)
                HistoryCmd.Push(message)
            }else { // redo the previous command
                LastRet = CallSystem(HistoryCmd.Peek())
            }
        }
    }
}

func runSender(){
    if narg < 1 {
        return 
    }else{
        fifoName := fifoPath(flag.Arg(0)) 
        sender,err := NewSender(fifoName)
        if err != nil {
            fmt.Println(err)
            return 
        }
        defer sender.Close()
        
        var cmdline string
        if flagOriginal || flagInternal {
            cmdline = strings.Join(flag.Args()[1:]," ")

            if flagInternal {
                cmdline = "#cmd:"+cmdline 
            }
        }else {
            cmdline = ArgsToCmdline(flag.Args()[1:])
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
