package ig

import (
	"fmt"
	"io"
	"strings"
)

type Video struct {
	Path   string
	Width  int
	Height int
}

func (p Video) Date() string { return p.Path }

func (p Video) Render() string {
	var b strings.Builder
	p.RenderInto(&b)
	return b.String()
}

func (p Video) RenderInto(out io.StringWriter) {
	out.WriteString(`<div style="display: contents;">`)
	out.WriteString(fmt.Sprintf(`<a href="%s" target="_blank" style="display: contents;">`, p.Path))
	out.WriteString("\n")
	out.WriteString(fmt.Sprintf(`<video width="%d" height="%d" controls style="margin: 8px; object-fit: cover; width: %dpx; height: %dpx;" loading="lazy">`, p.Width, p.Height, p.Width, p.Height))
	out.WriteString("\n")
	out.WriteString(fmt.Sprintf(`<source src="%s" type="video/mp4">`, p.Path))
	out.WriteString("\n")
	out.WriteString("</video>")
	out.WriteString("\n")
	out.WriteString("</div>")
	out.WriteString("\n")
}
