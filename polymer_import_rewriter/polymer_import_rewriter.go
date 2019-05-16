package polymer_import_rewrite

import (
	"net/http"

	"github.com/volts-dev/volts/server"
)

// root : the dir name of app project
func PolymerServe(root string, hd *server.TWebHandler) {
	p := hd.PathParams()
	ext := p.AsString("ext")

	path := hd.Request().URL.Path
	file_path := filepath.Join(
		server.MODULE_DIR, // c:\project\Modules
		path)
	osfile, e := os.Open(file_path)
	if e != nil {
		hd.Abort(404, e.Error())
		return
	}
	defer osfile.Close()

	info, e := osfile.Stat()
	if e != nil {
		hd.Abort(404, e.Error())
		return
	}

	if ext == "js" || ext == "ts" {
		p := pir.NewParser()
		p.SetRoot(root)
		p.SetPath(path)
		p.Parse(osfile)

		buf := p.Buffer()
		http.ServeContent(hd.Response(), hd.Request(), info.Name(), info.ModTime(), bytes.NewReader(buf.Bytes()))
	} else {
		http.ServeContent(hd.Response(), hd.Request(), info.Name(), info.ModTime(), osfile)
	}
}
