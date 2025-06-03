package main

import "github.com/01-edu/z01"

const (
	OPEN  = true
	CLOSE = false
)

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Is the Door opened?\n")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?\n")
	return !ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)

	if IsDoorClose(door) {
		OpenDoor(door)
	}

	if IsDoorOpen(door) {
		CloseDoor(door)
	}

	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
