// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package testcssexpression

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "strings"

func className() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#ffffff;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-height:calc(100vh - 170px);`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, red)))
	templ_7745c5c3_CSSID := templ.CSSID(`className`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}
