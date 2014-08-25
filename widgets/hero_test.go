package widgets
import (
	"github.com/remkade/scribble/widgets"
	"testing"
)

func TestHeroRender(t *Testing) {
	hero := Hero{CSSClasses: "testClass", ImageURL: "//example.com/test/image.png"}
	rendered := hero.Render()
	expected := `<img class="testClass" src="//example.com/test/image.png">`
	if rendered != expected {
		t.Fail("Expected: %s to match %s", rendered, expected)
	}
}
