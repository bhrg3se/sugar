package routes

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"sugar/models"
	"sugar/pkg/users"
	"sugar/store"
	"sugar/utils"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			logrus.Debug(err)
			utils.ErrorResponse(w, "cookie not found", 401)
			return
		}

		token := c.Value

		parsedToken, err := utils.VerifyAndParseToken(token)
		if err != nil {
			logrus.Error(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}

		user, err := users.FetchUser(store.State.DB, parsedToken.UserID)
		if err != nil {
			logrus.Error("invalid user id in jwt token: ", err)
			utils.ErrorResponse(w, "invalid user_id", 500)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	})

}

func AdminOnly(next http.Handler) http.Handler {
	return Authentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if !user.IsAdmin {
			utils.ErrorResponse(w, "not authorised for normal users", 403)
			return
		}
		next.ServeHTTP(w, r)
	}))
}
