package polymer_import_rewrite

import (
	"github.com/volts-dev/volts/server"
)

type (
	TPolymerImport struct{}
)

func (self *TPolymerImport) Response(act interface{}, route *server.TController) {
	if act != nil {
		/*
			web := route.GetHttpHandler()
			path := web.Request().URL.Path

			finfo, err := os.Stat(path)
			if err != nil {
				http.NotFound(ctx.ResponseWriter, ctx.Request)
				return
			}

			osfile, e := os.Open(path)
			if e != nil {
				return nil, e
			}
			defer osfile.Close()
		*/
	}
}
