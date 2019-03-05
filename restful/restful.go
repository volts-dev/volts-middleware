package restful

import (
	"vectors/volts/server"
)

type (
	WebBefore interface {
		Before(hd *server.TWebHandler)
	}

	WebAfter interface {
		After(hd *server.TWebHandler)
	}

	WebPanic interface {
		Panic(hd *server.TWebHandler)
	}

	RpcBefore interface {
		Before(hd *server.TRpcHandler)
	}

	RpcAfter interface {
		After(hd *server.TRpcHandler)
	}

	RpcPanic interface {
		Panic(hd *server.TRpcHandler)
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

func (self *TEvent) Request(act interface{}, route *server.TController) {
	if act != nil {
		web := route.GetHttpHandler()
		if web != nil {
			if a, ok := act.(WebBefore); ok {
				a.Before(web)
			}
		}
		rpc := route.GetRpcHandler()
		if rpc != nil {
			if a, ok := act.(RpcBefore); ok {
				a.Before(rpc)
			}
		}
	}
}

func (self *TEvent) Response(act interface{}, route *server.TController) {
	if act != nil {
		web := route.GetHttpHandler()
		if web != nil {
			if a, ok := act.(WebAfter); ok {
				a.After(web)
			}
		}
		rpc := route.GetRpcHandler()
		if rpc != nil {
			if a, ok := act.(RpcAfter); ok {
				a.After(rpc)
			}
		}
	}
}

func (self *TEvent) Panic(act interface{}, route *server.TController) {
	if act != nil {
		web := route.GetHttpHandler()
		if web != nil {
			if a, ok := act.(WebPanic); ok {
				a.Panic(web)
			}
		}
		rpc := route.GetRpcHandler()
		if rpc != nil {
			if a, ok := act.(RpcPanic); ok {
				a.Panic(rpc)
			}
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
