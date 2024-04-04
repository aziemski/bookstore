// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package homeviews

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HomeIndex(name string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><title>Bookstore</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" type=\"text/css\" href=\"css/styles.css\"></head><body><header><div class=\"logo\"><a href=\"\">Bookstore</a></div><nav><ul><li><a href=\"\">Home</a></li><li><a href=\"\">Categories</a></li><li><a href=\"\">About Project</a></li><li><a href=\"\">Contact</a></li></ul></nav><div class=\"search\"><label><input type=\"text\" placeholder=\"Search...\"></label> <button type=\"submit\">Find!</button></div></header><div class=\"main\"><div class=\"welcome-content\"><h1>Welcome to Bookstore</h1><p>Find your next favourite book today</p></div><div class=\"main-content\"><h2>Featured Books</h2><div class=\"books\"><div class=\"book\"><img src=\"https://httpcats.com/201.jpg\" alt=\"\"><h3>Book Title 201</h3><p>Author Name</p><p>$5.99</p></div><div class=\"book\"><img src=\"https://httpcats.com/101.jpg\" alt=\"\"><h3>Book Title 101</h3><p>Author Name</p><p>$20.55</p></div><div class=\"book\"><img src=\"https://http.dog/303.jpg\" alt=\"\"><h3>Book Title 303</h3><p>Author Name</p><p>$10.12</p></div><div class=\"book\"><img src=\"https://http.dog/404.jpg\" alt=\"\"><h3>Book Title 404 </h3><p>Author Name</p><p>$14.32</p></div></div></div></div><footer><p>&copy; 2024 Bookstore. All right reserved </p></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
