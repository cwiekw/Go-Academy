package validator

import (
	"MovieManager/internal/web/api"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	"go.uber.org/zap"
)

func ValidateStructMiddleware(log *zap.Logger) api.StrictMiddlewareFunc {
	return func(f strictecho.StrictEchoHandlerFunc, _ string) strictecho.StrictEchoHandlerFunc {
		return func(ctx echo.Context, request interface{}) (response interface{}, err error) {
			method := ctx.Request().Method
			if method != http.MethodPost && method != http.MethodPut {
				log.Info("Skipping body validation", zap.String("method", method), zap.String("url", ctx.Request().URL.Path))
				return f(ctx, request)
			}

			log.Info("Body validator", zap.String("action", "started"))
			validate := validator.New(validator.WithRequiredStructEnabled())

			err = validate.Struct(request)
			if err != nil {
				resp := ctx.Response()
				resp.WriteHeader(http.StatusPreconditionFailed)

				log.Info("Body validator", zap.String("action", "failed"), zap.Error(err))
				return resp, nil
			}

			log.Info("Body validator", zap.String("action", "finished"))
			return f(ctx, request)
		}
	}
}
