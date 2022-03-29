package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// Clear clears the screen
func Clear() {
	var (
		cursor coord
		w      dword
		h      = getScreen()
	)

	total := dword(h.size.x * h.size.y)

	xFillConsoleOutputCharacter.Call(
		uintptr(h.handle),
		uintptr(' '),
		uintptr(total),
		*(*uintptr)(unsafe.Pointer(&cursor)),
		uintptr(unsafe.Pointer(&w)),
	)

	xFillConsoleOutputAttribute.Call(
		uintptr(h.handle),
		uintptr(h.attributes),
		uintptr(total), *(*uintptr)(unsafe.Pointer(&cursor)),
		uintptr(unsafe.Pointer(&w)),
	)
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	h := getScreen()

	xSetConsoleCursorPosition.Call(
		uintptr(h.handle),
		*(*uintptr)(unsafe.Pointer(&coord{})),
	)
}

func getScreen() consoleScreenBufferInfoHandle {
	h := consoleScreenBufferInfoHandle{
		handle: syscall.Handle(os.Stdout.Fd()),
	}

	xGetConsoleScreenBufferInfo.Call(
		uintptr(h.handle),
		uintptr(unsafe.Pointer(&h.consoleScreenBufferInfo)),
	)

	return h
}

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	xGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	xSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	xFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
	xFillConsoleOutputAttribute = kernel32.NewProc("FillConsoleOutputAttribute")
)

type (
	wchar uint16
	short int16
	dword uint32
	word  uint16
)

type coord struct {
	x short
	y short
}

type smallRect struct {
	left   short
	top    short
	right  short
	bottom short
}

type consoleScreenBufferInfo struct {
	size              coord
	cursorPosition    coord
	attributes        word
	window            smallRect
	maximumWindowSize coord
}

type consoleScreenBufferInfoHandle struct {
	handle syscall.Handle
	consoleScreenBufferInfo
}

const cLength = 8

func main() {

	// for line := 0; line < len(digits[0]); line++ {
	// 	for di := range digits {
	// 		fmt.Print(digits[di][line], "  ")
	// 	}
	// 	fmt.Println()
	// }

	// var alarmed bool
	clockT, clock := [cLength]placeholder{}, [cLength]placeholder{}
	returned := false

	Clear()

	for shift := 0; shift <= cLength; shift++ {
		Clear()
		MoveTopLeft()
		now := time.Now()

		hour, minutes, seconds := now.Hour(), now.Minute(), now.Second()
		if shift == cLength {
			returned = !returned
			shift = 0
		}

		clockT = [cLength]placeholder{
			digits[hour/10], digits[hour%10], colon,
			digits[minutes/10], digits[minutes%10], colon,
			digits[seconds/10], digits[seconds%10],
		}

		if returned {
			for i := cLength - shift; i < cLength; i++ {
				clock[i] = clockT[i-cLength+shift]
			}
			for i := 0; i < cLength-shift; i++ {
				clock[i] = empty
			}
		} else {
			for i := 0; i < cLength-shift; i++ {
				clock[i] = clockT[i+shift]
			}
			for i := cLength - 1; i > cLength-shift-1; i-- {
				clock[i] = empty
			}
		}

		for line := 0; line < len(clock[0]); line++ {
			for di := range clock {
				next := clock[di][line]
				// if digit == colon && seconds%2 == 0 {
				// 	next = "   "
				// }
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second / 2)
	}

}
