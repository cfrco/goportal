package main

import (
    "container/list"
    "fmt"
)

var (
    maxlen int = 10
    HistoryCmd History
)

type History struct {
    list *list.List
}

func (history *History) Init() {
    history.list = list.New()
}

func (history *History) Push(cmd string) {
    if history.list.Len() >= maxlen {
        history.list.Remove(history.list.Front()) 
    }
    history.list.PushBack(cmd)
}

func (history *History) Peek() string {
    if history.list.Len() > 0 {
        return history.list.Back().Value.(string)
    }
    return ""
}

func (history *History) Display() {
    for e := history.list.Front(); e!=nil; e = e.Next() {
        fmt.Println(e.Value.(string)) 
    }
}
