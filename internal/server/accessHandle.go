package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stillsaro/Leenky/internal/database/DBgo"
	"golang.org/x/crypto/bcrypt"
)

type SignupUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) RootHandler(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}

func (s *Server) Signup(c echo.Context) error {
	return c.Render(200, "signup.html", nil)
}

func (s *Server) Login(c echo.Context) error {
	return c.Render(200, "login.html", nil)
}

func (s *Server) LeenkyChat(c echo.Context) error {
	type Data struct {
		Username string `json:"username"`
	}
	d := Data{
		Username: c.Get("username").(string),
	}

	return c.Render(200, "leenkychat.html", d)
}

func (s *Server) SignupHandler(c echo.Context) error {
	formValues := SignupUserInput{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	_, err := s.Queries.GetUserByEmail(c.Request().Context(), formValues.Email)
	if err == nil {
		return c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"This email is already registerd. Try loggin or use different email</div>")
	}
	if err != sql.ErrNoRows {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			" An error occurred. Please try again later.</div>")
	}

	_, err = s.Queries.GetUserByUsername(c.Request().Context(), formValues.Username)
	if err == nil {
		return c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"Username taken. Please choose another one</div>")
	}
	if err != sql.ErrNoRows {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"An error occurred. Please try again later.</div>")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formValues.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"An error occurred. Please try again later.</div>")
	}

	err = s.Queries.InsertUser(c.Request().Context(), DBgo.InsertUserParams{
		Username:     formValues.Username,
		Email:        formValues.Email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"An error occurred. Please try again later.</div>")
	}

	return c.HTML(http.StatusOK, "<div id='signup-response' class='success'>"+
		"Your account has been created successfully. Redirecting to login...</div>"+
		"<script>setTimeout(function() { window.location.href = '/login'; }, 2000);</script>")
}

func (s *Server) LoginHandler(c echo.Context) error {
	log.Println("we are in login handler")
	formValues := LoginUserInput{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	user, err := s.Queries.GetUserByUsername(c.Request().Context(), formValues.Username)
	if err == sql.ErrNoRows {
		return c.HTML(http.StatusOK, "<div id='login-response' class='error'>"+
			"Invalid username or password.</div>")

	} else if err != nil {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"An error occurred. Please try again later.</div>")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(formValues.Password))
	if err != nil {
		return c.HTML(http.StatusOK, "<div id='login-response' class='error'>"+
			"Invalid username or password.</div>")
	}

	token, err := s.Auth.JwtGenerator(formValues.Username)
	if err != nil {
		c.Logger().Error(err)
		c.HTML(http.StatusOK, "<div id='signup-response' class='error'>"+
			"An error occurred. Please try again later.</div>")
	}
	SetJwtCookie(c, token)

	return c.HTML(http.StatusOK, "<div id='login-response' class='success'>"+
		"You've logged in successfully! Redirecting to your area...</div>"+
		"<script>setTimeout(function() { window.location.href = '/leenkychat'; }, 2000);</script>")
}
