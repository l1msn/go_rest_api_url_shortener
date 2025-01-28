package delete

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"url_shortner/internal/lib/api/response"
	"url_shortner/internal/lib/logger/sl"
)

type Request struct {
	Alias string `json:"alias" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type URLDeleter interface {
	DeleteURL(alias string) error
	CheckAliasExist(alias string) (bool, error)
}

func Delete(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.Delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to deserialize request", sl.Err(err))

			render.JSON(w, r, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			var validateErr validator.ValidationErrors

			errors.As(err, &validateErr)

			log.Error("failed to validate request", sl.Err(err))

			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}

		url := req.Alias

		isAliasExist, err := urlDeleter.CheckAliasExist(url)

		if err != nil {
			log.Error("cannot check existed alias", sl.Err(err))

			render.JSON(w, r, response.Error("cannot check existed alias"))

			return
		}

		if !isAliasExist {
			log.Error("is alias doesnt exist")

			render.JSON(w, r, response.Error("is alias doesnt exist"))

			return
		}

		err = urlDeleter.DeleteURL(url)

		if err != nil {
			log.Error("cannot delete url", sl.Err(err))

			render.JSON(w, r, response.Error("cannot delete url"))

			return
		}

		log.Info("url deleted", slog.String("url", url))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Status: response.StatusOK,
	})
}
