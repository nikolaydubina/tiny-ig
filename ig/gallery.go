package ig

import (
	"io"
	"strings"
)

type Renderable interface {
	Render() string
	RenderInto(out io.StringWriter)
	Date() string
}

type Gallery struct {
	Posts []Renderable
}

func (g Gallery) Render() string {
	var b strings.Builder
	g.RenderInto(&b)
	return b.String()
}

func (g Gallery) RenderInto(out io.StringWriter) {
	out.WriteString(`<div align="center">`)
	out.WriteString("\n")
	for _, p := range g.Posts {
		out.WriteString(p.Render())
		out.WriteString("\n")
	}
	out.WriteString("</div>")
	out.WriteString("\n")
}
