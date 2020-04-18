package middleware

import (
	"app/conf"
	"app/models"
	"context"
	"net/http"
)

func GetUser(ctx context.Context) *models.User {
	user, ok := ctx.Value("user").(*models.User)
	if !ok {
		return nil
	}
	return user
}

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func getSessionUser(r *http.Request) *models.User {
	session, _ := conf.CookieStore.Get(r, "session")
	userid, _ := session.Values["user_id"]
	if userid == nil {
		return nil
	}
	user := &models.User{}
	if err := models.Db.Where("id = ?", userid).First(&user).Error; err == nil {
		if user.CanLogin() {
			return user
		}
	}
	return nil
}

func CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getSessionUser(r)
		if user != nil {
			ctx := WithUser(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r.Context())
		if user == nil {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
