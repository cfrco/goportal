package main

import (
    "syscall"
    "os"
    "bufio"
    "strings"
)

type Receiver struct {
    FifoName string
    Mode int
    fd *os.File
    reader *bufio.Reader
}

const (
    MODE_RDWR = syscall.O_RDWR
    MODE_RDONLY = syscall.O_RDONLY
)

func NewReceiver(fifoname string, mode int) (Receiver,error) {
    err := syscall.Mkfifo(fifoname,0600)
    if err != nil{
        if err.Error() != "file exists"{
            return Receiver{},err
        }
    }

    fd,err := os.OpenFile(fifoname,mode,0)
    if err != nil{
        return Receiver{},err
    }
    
    rec := Receiver{fifoname,mode,fd,bufio.NewReader(fd)}
    return rec,nil
}

func (receiver *Receiver) ReadMessage() string {
    line,err := receiver.reader.ReadString(byte('\n'))
    if err != nil{
        return "" 
    }
    return strings.Trim(line,"\n")
}

func (receiver *Receiver) Close() {
    receiver.fd.Close()
    receiver.fd = nil
    receiver.reader = nil

    os.Remove(receiver.FifoName)
}
