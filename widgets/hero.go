package widgets

import (
	"fmt"
)

type Hero struct {
	Name string
	CSSClasses []string
	ImageURL string
}

func (h *Hero) Render() string {
	return fmt.Sprintf(`<img class="%s" src="%s">`, h.CSSClasses, h.ImageURL)
}
