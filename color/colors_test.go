package color

import (
	"testing"
)

func TestHilight(t *testing.T) {
	Hilight("Hilight: hello, %v!", "world")
}

func TestDimV(t *testing.T) {
	DimV("DimV (verbose build only): hello, %v!", "world")
}

func TestText(t *testing.T) {
	Text("Text: hello, %v!\n", "world")
}

func TestDim(t *testing.T) {
	Dim("Dim: hello, %v!", "world")
}

func TestToDim(t *testing.T) {
	t.Logf("%v", ToDim("ToDim: hello, %v!", "world"))
}

func TestToColor(t *testing.T) {
	t.Logf("%v", ToColor(Magenta, "ToColor: hello, %v!", "world"))
}

func TestColoredV(t *testing.T) {
	ColoredV(LightMagenta, "ColoredV (verbose build only): hello, %v!", "world")
}

func TestColored(t *testing.T) {
	Colored(LightMagenta, "Colored: hello, %v!", "world")
}
