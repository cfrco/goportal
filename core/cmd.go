package core

import (
    "strings"
    "os"
    "fmt"
)

type CmdError struct{
    ErrorMessage string
}

func (e *CmdError) Error() string {
    return e.ErrorMessage
}

var LastRet int
var PrevCmd string

func RunInternalCommand(cmdline string) error{
    cmdline = cmdline[5:]
    cmd := strings.SplitN(cmdline," ",2)

    cmd[0] = strings.ToLower(strings.Trim(cmd[0],"\" "))
    if len(cmd) > 1 {
        cmd[1] = strings.Trim(cmd[1],"\" ")
        switch cmd[0]{
            case "cd" :
                fmt.Printf("Chdir( %s )\n",cmd[1])
                return os.Chdir(cmd[1])
            default :
                return &CmdError{"No command."} 
        }
    }else{
        switch cmd[0]{
            case "ret" :
                fmt.Printf("%d\n",LastRet)
            case "end" :
                return &CmdError{"command:end"}
            default :
                return &CmdError{"No command."} 
        }
    }

    return nil
}
