package desc_test

import (
	"github.com/LiterMC/wasm-jdk/desc"

	"testing"
)

func TestParseDesc(t *testing.T) {
	var datas = []struct {
		S string
		E bool
	}{
		{"I", false},
		{"[I", false},
		{"V", false},
		{"[V", false},
		{"[[[[[[[[J", false},
		{"[[", true},
		{"L", true},
		{"Ljava/lang/Object;", false},
		{"LTest$x123;", false},
		{"LIL;", false},
		{"[LBZ;", false},
	}
	for _, d := range datas {
		dc, err := desc.ParseDesc(d.S)
		if d.E {
			if err == nil {
				t.Errorf("unexpectedly successful parsed invalid descriptor %q as %#v", d.S, dc)
			}
		} else if err != nil {
			t.Errorf("failed parse %q: %v", d.S, err)
		} else if dc.String() != d.S {
			t.Errorf("parsed desc %q not match %q", dc.String(), d.S)
		}
	}
}

func TestParseMethodDesc(t *testing.T) {
	var datas = []struct {
		S string
		E bool
	}{
		{"()V", false},
		{"()", true},
		{"(V)", true},
		{"(V", true},
		{"V", true},
		{"(V)V", false},
		{"([V)I", false},
		{"([[[[[[[[J)Ljava/lang/Object;", false},
		{"[[", true},
		{"L", true},
		{"(Ljava/lang/Object;IJ)D", false},
		{"(ILTest$x123;Ljava/lang/Object;VLID;)D", false},
	}
	for _, d := range datas {
		dc, err := desc.ParseMethodDesc(d.S)
		if d.E {
			if err == nil {
				t.Errorf("unexpectedly successful parsed invalid descriptor %q as %#v", d.S, dc)
			}
		} else if err != nil {
			t.Errorf("failed parse %q: %v", d.S, err)
		} else if dc.String() != d.S {
			t.Errorf("parsed desc %q not match %q", dc.String(), d.S)
		}
	}
}
