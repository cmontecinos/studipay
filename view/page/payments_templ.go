// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.590
package page

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Payments() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-bs-theme=\"auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = head().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = lightmode().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"navbar sticky-top bg-dark flex-md-nowrap p-0 shadow\" data-bs-theme=\"dark\"><a class=\"navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6 text-white\" href=\"#\">Studipay</a><ul class=\"navbar-nav flex-row d-md-none\"><li class=\"nav-item text-nowrap\"><button class=\"nav-link px-3 text-white\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarSearch\" aria-controls=\"navbarSearch\" aria-expanded=\"false\" aria-label=\"Toggle search\"><svg class=\"bi\"><use xlink:href=\"#search\"></use></svg></button></li><li class=\"nav-item text-nowrap\"><button class=\"nav-link px-3 text-white\" type=\"button\" data-bs-toggle=\"offcanvas\" data-bs-target=\"#sidebarMenu\" aria-controls=\"sidebarMenu\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><svg class=\"bi\"><use xlink:href=\"#list\"></use></svg></button></li></ul><div id=\"navbarSearch\" class=\"navbar-search w-100 collapse\"><input class=\"form-control w-100 rounded-0 border-0\" type=\"text\" placeholder=\"Search\" aria-label=\"Search\"></div></header><div class=\"container-fluid\"><div class=\"row\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Menu().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"col-md-9 ms-sm-auto col-lg-10 px-md-4\"><div class=\"d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom\"><h1 class=\"h2\">Pagos</h1></div><form class=\"row g-3\"><div class=\"col-md-2\"><label for=\"inputAlumno\" class=\"form-label\">Alumno</label> <select class=\"form-select\" id=\"inputAlumno\"><option value=\"alumno1\">Alumno 1</option> <option value=\"alumno2\">Alumno 2</option> <option value=\"alumno3\">Alumno 3</option></select></div><div class=\"col-md-2\"><label for=\"inputMonto\" class=\"form-label\">Monto</label> <input type=\"text\" class=\"form-control input-small\" id=\"inputMonto\" placeholder=\"Monto\"></div><div class=\"col-12\"><button type=\"submit\" class=\"btn btn-primary\">Agregar Pago</button></div></form></main></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
