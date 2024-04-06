// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package aboutviews

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AbouProject(query string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><title>Anybook</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" type=\"text/css\" href=\"/static/css/styles.css\"></head><body><header><div class=\"logo\"><a href=\"/\">Anybook</a></div><nav><ul><li><a href=\"/\">Home</a></li><li><a href=\"/static/docs/api/index.html\">REST API</a></li><li><a href=\"/about\">About Project</a></li></ul></nav><div class=\"search\"><form action=\"/search\" method=\"GET\"><label><input type=\"text\" placeholder=\"Search...\" id=\"q\" name=\"q\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(query)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/about/aboutviews/main.templ`, Line: 27, Col: 92}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></label> <button type=\"submit\">Find!</button></form></div></header><div class=\"main\"><div class=\"main-content\"><div class=\"about-content\"><h3>Used technologies</h3><ul><li><a href=\"https://tip.golang.org/doc/go1.22\">Go 1.22</a></li><li><a href=\"https://entgo.io/\">Ent. An entity framework for Go</a></li><li><a href=\"https://www.sqlite.org/\">SQLite</a></li><li><a href=\"https://echo.labstack.com/\">Echo. High performance, extensible, minimalist Go web framework</a></li><li><a href=\"https://github.com/a-h/templ/\">templ. A language for writing HTML user interfaces in Go.</a></li><li><a href=\"https://openapi-generator.tech/\">OpenAPI Doc Generator</a></li><li><a href=\"https://github.com/deepmap/oapi-codegen/\">OpenAPI Server Code Generator</a></li><li><a href=\"https://letsencrypt.org/\">Let's encrypt for TLS certificates</a></li><li><a href=\"https://www.hetzner.com/\">VPS Provider: Hetzner</a></li><li><a href=\"https://www.godaddy.com/\">Domain Provider: GoDaddy</a></li></ul></div></div></div><footer><p>&copy; 2024 Anybook. All right reserved </p></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
