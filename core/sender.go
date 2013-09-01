package core

import(
    "syscall"
    "os"
    "bufio"
    "strings"
)

type Sender struct{
    FifoName string
    fd *os.File
    writer *bufio.Writer
}

func NewSender(fifoname string) (Sender,error){
    fd,err := os.OpenFile(fifoname,syscall.O_WRONLY,0) 
    if err != nil{
        return Sender{},err 
    }

    sen := Sender{fifoname,fd,bufio.NewWriter(fd)}
    return sen,nil
}

func (sender *Sender) SendMessage(message string) error{
    line := message+"\n" // Ending with '\n'
    _,err := sender.writer.WriteString(line)
    sender.writer.Flush()
    return err
}

func (sender *Sender) Close(){
    sender.fd.Close() 
    sender.fd = nil
    sender.writer = nil
}

func ArgsToCmdline(args []string) string{
    var cmdline string

    for i:=0;i<len(args);i++{
        cmdline = cmdline+" \""+strings.Replace(args[i],"\"","\\\"",-1)+"\" "
    }
    return cmdline
}
