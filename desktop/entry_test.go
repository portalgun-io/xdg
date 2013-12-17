package desktop

import (
	"testing"
	"strings"
)

// specExample is the example desktop file given in the spec.
const specExample = `
[Desktop Entry]
Version=1.0
Type=Application
Name=Foo Viewer
Comment=The best viewer for Foo objects available!
TryExec=fooview
Exec=fooview %F
Icon=fooview
MimeType=image/x-foo;
Actions=Gallery;Create;

[Desktop Action Gallery]
Exec=fooview --gallery
Name=Browse Gallery

[Desktop Action Create]
Exec=fooview --create-new
Name=Create a new Foo!
Icon=fooview-new
`

func TestSpecExample(t *testing.T) {
	d, err := New(strings.NewReader(specExample))
	if err != nil {
		t.Error(err)
	}

	if d.Type() != Application {
		t.Error("Type")
	}

	arr := map[string]string {
		d.Version(): "1.0",
		d.Name(): "Foo Viewer",
		d.Comment(): "The best viewer for Foo objects available!",
		d.TryExec(): "fooview",
		d.Exec(): "fooview %F",
		d.Icon(): "fooview",
	}
	for act, exp := range arr {
		if act != exp {
			t.Log("expected: " + exp)
			t.Log("actual:   " + act)
			t.Fail()
		}
	}
}