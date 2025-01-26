package get

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"url_shortner/internal/lib/api/response"
	"url_shortner/internal/lib/logger/sl"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	URL    string `json:"alias,omitempty"`
}

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func Get(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Error("cant get alias from url or is nil")

			render.JSON(w, r, response.Error("cant get alias from url or is nil"))

			return
		}

		url, err := urlGetter.GetURL(alias)

		if err != nil {
			log.Error("not found url with that alias", sl.Err(err))

			render.JSON(w, r, response.Error("not found url with that alias"))

			return
		}

		log.Info("url getted", slog.String("url", url))

		responseOK(w, r, url)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, url string) {
	render.JSON(w, r, Response{
		Status: response.StatusOK,
		URL:    url,
	})
}
