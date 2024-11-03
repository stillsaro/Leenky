package server

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("/home/saro/GeekHub/Go/Leenky/viwes/html/*.html")),
	}
}

func SetJwtCookie(c echo.Context, token string) {
	cookie := http.Cookie{
		Name:     "Jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(&cookie)
	log.Println("Jwt cookie set with value: ", cookie.Value)
}
