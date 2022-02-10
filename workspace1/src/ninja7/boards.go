package main

import "fmt"

type Board struct {
	NailsNeeded int
	NailsDriven int
}

type Mallet struct{}

// DriveNail pounds a nail into the specified board.
func (Mallet) DriveNail(nailSupply *int, b *Board) {
	// Take a nail out of the supply.
	*nailSupply--

	// Pound a nail into the board.
	b.NailsDriven++

	fmt.Println("Mallet: pounded nail into the board.")
}

// Crowbar is a tool that removes nails.
type Crowbar struct{}

// PullNail yanks a nail out of the specified board.
func (Crowbar) PullNail(nailSupply *int, b *Board) {
	// Yank a nail out of the board.
	b.NailsDriven--

	// Put that nail back into the supply.
	*nailSupply++

	fmt.Println("Crowbar: yanked nail out of the board.")
}

type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// Contractor carries out the task of securing boards.
type Contractor struct{}

// Fasten will drive nails into a board.
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)

		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

// Toolbox can contains any number of tools.
type Toolbox struct {
	NailDriver
	NailPuller

	nails int
}

func displayState(tb *Toolbox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}

func main() {
	// Inventory of old boards to remove, and the new boards
	// that will replace them.
	boards := []Board{
		// Rotted boards to be removed.
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},

		// Fresh boards to be fastened.
		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
	}

	tb := Toolbox{
		NailDriver: &Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)

	// Display the new state of our toolbox and boards.
	displayState(&tb, boards)
}
