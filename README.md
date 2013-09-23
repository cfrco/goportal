# goportal
Send command to another terminal that running goportal in Go.

## Motivation
Since continuously, I hope i can run a shell in Vim. But i don't find a good solution. One day, I find `tslime.vim` (https://github.com/jgdavey/tslime.vim). But i don't use `tmux`. So I think I can write a program send command from terminal to terminal.

## Installation
Install or update 
```
go get -u github.com/cfrco/goportal
```

## Usage
### Start a Receiver
```
goportal -r <name>
```
### End a Receiver
Always use this command to end receiver. For sure it remove FIFO files.
```
goportal -i <name> end
```

### Use Sender to send command
#### Normal command
```
goportal <name> <command...>
```
#### Redo previous command
```
goportal <name>
```

#### Send goportal interanl command
```
goportal -i <name> <internal_cmd> <args...>
```

#### Send non-processed command
Normally, `goportal` automatically add `"` for each argument.  
For example,`ls -al` will be `"ls" "-al"` and `echo '"hello world!"'` will be `"echo" "\"hello world\""`.  
It's OK for those case. But it may cause problem in some case.(`ls -al | grep go` => `"ls" "-al" "|" "grep" "go"`).  
And pipeline can't work. In this case, use `-o`.
```
goportal -o test ls -al "|" grep go
```
IMPORTANT : `|` should be `"|"`,or it will be pipeline's redirection for `goportal -o test ls -al`.

### Internal Command
```
goportal -i <name> cd <path>     # change directory.
goportal -i <name> ret           # display `$?`(the last command's return value).
goportal -i <name> history       # display history commands.
goportal -i <name> end           # end the receiver.
```

## Vim:vim-goportal
### Install 
```
cp -r vim ~/.vim/bundle/vim-goportal
```
### Usage
 * `:Gpd` : set/display the default name
 * `:Gpdd` : clear the default name
 * `:Gp` : normal command
 * `:Gpi` : internal command
 * `:Gpo` : original command

 * `function GoPortal(...)`
 * `function GoPortalDefault(...)`
