package event

import (
	"github.com/VectorsOrigin/web"
)

type (
	iBefore interface {
		Before(hd *web.THandler)
	}

	iAfter interface {
		After(hd *web.THandler)
	}

	iPanic interface {
		Panic(hd *web.THandler)
	}

	/*
		iFinally interface {
			Finally(hd *web.THandler)
		}
	*/
	TEvent struct {
	}
)

func NewEvent() *TEvent {
	return &TEvent{}
}

func (self *TEvent) Request(act interface{}, hd *web.THandler) {
	if act != nil {
		if a, ok := act.(iBefore); ok {
			a.Before(hd)
		}
	}
}

func (self *TEvent) Response(act interface{}, hd *web.THandler) {
	if act != nil {
		if a, ok := act.(iAfter); ok {
			a.After(hd)
		}
	}
}

func (self *TEvent) Panic(act interface{}, hd *web.THandler) {
	if act != nil {
		if a, ok := act.(iPanic); ok {
			a.Panic(hd)
		}
	}
}

/*
func (self *TEvent) Finally(act interface{}, hd *web.THandler) {
	if act != nil {
		if a, ok := act.(iFinally); ok {
			a.After(hd)
		}
	}
}
*/
