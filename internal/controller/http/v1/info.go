package v1

import (
	"devops_course_app/internal/usecase"
	"devops_course_app/pkg/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type infoRoutes struct {
	i usecase.InfoContract
}

func NewInfoRoutes(routes chi.Router, i usecase.InfoContract) {
	ir := &infoRoutes{i: i}

	routes.Get("/", ir.getCurrencyRate)
}

type resp struct {
	Data    map[string]string `json:"data"`
	Service string            `json:"service"`
}

func (i *infoRoutes) getCurrencyRate(w http.ResponseWriter, r *http.Request) {
	currency := r.URL.Query().Get("currency")
	date := r.URL.Query().Get("date")

	response, err := i.i.GetCurrencyRate(currency, date)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Printf("Rendering error")
			return
		}
		return
	}
	responseJSON := resp{Data: map[string]string{currency: response}, Service: "currency"}
	render.JSON(w, r, responseJSON)
}
