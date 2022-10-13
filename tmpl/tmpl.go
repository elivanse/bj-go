package tmpl

import "net/http"

func IndexPage(w http.ResponseWriter, r *http.Request) {

	Templates.ExecuteTemplate(w, "ppal", nil)

}
