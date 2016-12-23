package event

import (
	"webgo"
)

type (
	iBefore interface {
		Before(hd *webgo.THandler)
	}

	iAfter interface {
		After(hd *webgo.THandler)
	}

	iPanic interface {
		Panic(hd *webgo.THandler)
	}

	/*
		iFinally interface {
			Finally(hd *webgo.THandler)
		}
	*/
	TEvent struct {
	}
)

func NewEvent() *TEvent {
	return &TEvent{}
}

func (self *TEvent) Request(act interface{}, hd *webgo.THandler) {
	if act != nil {
		if a, ok := act.(iBefore); ok {
			a.Before(hd)
		}
	}
}

func (self *TEvent) Response(act interface{}, hd *webgo.THandler) {
	if act != nil {
		if a, ok := act.(iAfter); ok {
			a.After(hd)
		}
	}
}

func (self *TEvent) Panic(act interface{}, hd *webgo.THandler) {
	if act != nil {
		if a, ok := act.(iPanic); ok {
			a.Panic(hd)
		}
	}
}

/*
func (self *TEvent) Finally(act interface{}, hd *webgo.THandler) {
	if act != nil {
		if a, ok := act.(iFinally); ok {
			a.After(hd)
		}
	}
}
*/
