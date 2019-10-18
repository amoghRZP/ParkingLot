package models

import (
	"fmt"
)

// A slot in parking lot
type Slot struct {
	Number uint
	Car    *Car
}

func New() *Slot {
	return &Slot{}
}

func (s *Slot) Occupy(cr Car) error {
	if s.Car != nil {
		return fmt.Errorf("slot: Slot already occupied")
	}
	s.Car = &cr
	return nil
}
