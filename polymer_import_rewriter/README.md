# the polymer3 template importer transformate

help you develope polymer3 by golang http server instead of command 'polymer serve'

1. Setup the URL route for all template dir which the polymer app are included in it
`
	web.Url("GET", "/web/template/(*).(string:ext)", staticCtrl.polymer_import) // do not change "ext"
`

2. Setup the controller to serve the route and point out which polymer app dir will be serve.
`
import	pir "github.com/volts-dev/volts-middleware/polymer_import_rewriter"

type(
	staticCtrl struct {
	}
)

func (self staticCtrl) polymer_import(hd *server.TWebHandler) {
	pir.PolymerServe("polymer_app", hd)
}
`

3. done