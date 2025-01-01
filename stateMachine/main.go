package main

import "fmt"

const (
	StateOpen   = "open"
	StateClose  = "close"
	StateLocked = "locked"
)

const (
	EventOpen   = "open"
	EventClose  = "close"
	EventLock   = "lock"
	EventUnlock = "unlock"
)

type StateMachine struct {
	state       string
	transitions map[string]map[string]string
}

func NewStateMachine(
	initialState string,
	transitions map[string]map[string]string) *StateMachine {
	return &StateMachine{
		state:       initialState,
		transitions: transitions,
	}
}

func (sm *StateMachine) State() string {
	return sm.state
}

func (sm *StateMachine) Trigger(event string) error {
	if nextStep, ok := sm.transitions[sm.state][event]; ok {
		fmt.Println("Transitioning from", sm.state, "to", nextStep, "for event", event)
		sm.state = nextStep
		return nil
	}

	return fmt.Errorf("Invalid event %s for current state %s", event, sm.state)
}

func main() {
	transition := map[string]map[string]string{
		StateOpen: {
			EventClose: StateClose,
			EventLock:  StateLocked,
		},
		StateClose: {
			EventOpen: StateOpen,
		},
		StateLocked: {
			EventUnlock: StateClose,
		},
	}

	sm := NewStateMachine(StateOpen, transition)

	event := []string{EventOpen, EventOpen, EventLock, EventUnlock}
	for _, e := range event {
		err := sm.Trigger(e)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Final state is", sm.State())
}
