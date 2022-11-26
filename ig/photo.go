package ig

import (
	"fmt"
	"io"
	"strings"
)

type Photo struct {
	Path   string
	Width  int
	Height int
}

func (p Photo) Date() string { return p.Path }

func (p Photo) Render() string {
	var b strings.Builder
	p.RenderInto(&b)
	return b.String()
}

func (p Photo) RenderInto(out io.StringWriter) {
	out.WriteString(fmt.Sprintf(`<a href="%s" target="_blank" style="display: contents;">`, p.Path))
	out.WriteString("\n")
	out.WriteString(fmt.Sprintf(`<img src="%s" style="margin: 8px; object-fit: cover; width: %dpx; height: %dpx;" loading="lazy">`, p.Path, p.Width, p.Height))
	out.WriteString("\n")
	out.WriteString("</a>")
	out.WriteString("\n")
}
