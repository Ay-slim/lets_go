package main

import "fmt"
// Structs, think of these as objects or classes with properties and methods
type ServerState int64

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
	
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}
