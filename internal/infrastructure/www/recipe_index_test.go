package www_test

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lonepeon/food/internal/infrastructure/www"
	"github.com/lonepeon/golib/web/webtest"
)

func TestRecipeIndexSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := webtest.NewMockContext(ctrl)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	expected := webtest.MockedResponse("ok response")
	ctx.EXPECT().Response(200, gomock.Any(), nil).Return(expected)

	actual := www.RecipeIndex()(ctx, w, r)

	webtest.AssertResponse(t, expected, actual, "unexpected response")
}
