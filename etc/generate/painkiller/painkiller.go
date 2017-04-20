//go:generate stringer -type=Pill

package painkiller

import (
	"fmt"
)

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func (p Pill) String() string {
	switch p {
	case Placebo:
		return "Placebo"
	case Aspirin:
		return "Aspirin"
	case Ibuprofen:
		return "Ibuprofen"
	case Paracetamol:
		return "Paracetamol"
	}
	return fmt.Sprintf("Pill(%d)", p)
}
