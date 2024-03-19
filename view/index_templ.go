// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.590
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"col-md-9 ms-sm-auto col-lg-10 px-md-4\"><div class=\"d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom\"><h1 class=\"h2\">Dashboard</h1><div class=\"btn-toolbar mb-2 mb-md-0\"><div class=\"btn-group me-2\"><button type=\"button\" class=\"btn btn-sm btn-outline-secondary\">Share</button> <button type=\"button\" class=\"btn btn-sm btn-outline-secondary\">Export</button></div><button type=\"button\" class=\"btn btn-sm btn-outline-secondary dropdown-toggle d-flex align-items-center gap-1\"><svg class=\"bi\"><use xlink:href=\"#calendar3\"></use></svg> This week</button></div></div><canvas class=\"my-4 w-100\" id=\"myChart\" width=\"900\" height=\"380\"></canvas><h2>Section title</h2><div class=\"table-responsive small\"><table class=\"table table-striped table-sm\"><thead><tr><th scope=\"col\">#</th><th scope=\"col\">Header</th><th scope=\"col\">Header</th><th scope=\"col\">Header</th><th scope=\"col\">Header</th></tr></thead> <tbody><tr><td>1,001</td><td>random</td><td>data</td><td>placeholder</td><td>text</td></tr><tr><td>1,002</td><td>placeholder</td><td>irrelevant</td><td>visual</td><td>layout</td></tr><tr><td>1,003</td><td>data</td><td>rich</td><td>dashboard</td><td>tabular</td></tr><tr><td>1,003</td><td>information</td><td>placeholder</td><td>illustrative</td><td>data</td></tr><tr><td>1,004</td><td>text</td><td>random</td><td>layout</td><td>dashboard</td></tr><tr><td>1,005</td><td>dashboard</td><td>irrelevant</td><td>text</td><td>placeholder</td></tr><tr><td>1,006</td><td>dashboard</td><td>illustrative</td><td>rich</td><td>data</td></tr><tr><td>1,007</td><td>placeholder</td><td>tabular</td><td>information</td><td>irrelevant</td></tr><tr><td>1,008</td><td>random</td><td>data</td><td>placeholder</td><td>text</td></tr><tr><td>1,009</td><td>placeholder</td><td>irrelevant</td><td>visual</td><td>layout</td></tr><tr><td>1,010</td><td>data</td><td>rich</td><td>dashboard</td><td>tabular</td></tr><tr><td>1,011</td><td>information</td><td>placeholder</td><td>illustrative</td><td>data</td></tr><tr><td>1,012</td><td>text</td><td>placeholder</td><td>layout</td><td>dashboard</td></tr><tr><td>1,013</td><td>dashboard</td><td>irrelevant</td><td>text</td><td>visual</td></tr><tr><td>1,014</td><td>dashboard</td><td>illustrative</td><td>rich</td><td>data</td></tr><tr><td>1,015</td><td>random</td><td>tabular</td><td>information</td><td>text</td></tr></tbody></table></div></main></div></div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz\" crossorigin=\"anonymous\"></script><script src=\"https://cdn.jsdelivr.net/npm/chart.js@4.3.2/dist/chart.umd.js\" integrity=\"sha384-eI7PSr3L1XLISH8JdDII5YN/njoSsxfbrkCTnJrzXt+ENP5MOVBxD+l6sEG4zoLp\" crossorigin=\"anonymous\"></script><script src=\"../static/js/dashboard.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
