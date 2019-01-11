package exchange

import (
	"errors"
)

type action int

const (
	Get action = iota
	Create
	Edit
	Delete
)

type ExMaper struct {
	ExMap map[string]string
}

func NewExMaper() *ExMaper {
	return &ExMaper{
		ExMap: make(map[string]string),
	}
}

type Dollar struct {
	Name    string
	Value   string
	Action  action
	ExMaper *ExMaper
	ErrChan chan error
}

func (d *Dollar) Task() {
	switch d.Action {
	case Create:
		d.create()
}

func (d *Dollar) create() {
	_, ok := d.ExMaper.ExMap[d.Name]
	if !ok {
		d.ExMaper.ExMap[d.Name] = d.Value
		d.ErrChan <- nil
	} else {
		d.ErrChan <- errors.New("this dollar is alreay exist")
	}
}
