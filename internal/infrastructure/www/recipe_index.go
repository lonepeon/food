package www

import (
	"net/http"

	"github.com/lonepeon/golib/web"
)

func RecipeIndex() web.HandlerFunc {
	return func(ctx web.Context, w http.ResponseWriter, r *http.Request) web.Response {
		return ctx.Response(200, "templates/recipes/index.html.tmpl", nil)
	}
}
