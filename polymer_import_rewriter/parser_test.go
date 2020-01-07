package polymer_import_rewrite

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	osfile, e := os.Open("./example1.js")
	if e != nil {
		t.Log(e)
		return
	}
	defer osfile.Close()

	_, e = osfile.Stat()
	if e != nil {
		t.Log(e)
		return
	}

	p := NewParser()
	p.SetRoot("")
	p.SetPath("./example1.js")
	p.Parse(osfile)
	buf := p.Buffer()
	t.Log(buf.String())
}
