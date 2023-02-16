package route

import (
	"html/template"
	"io"
	config "learn-echo-renderer/Config"
	"learn-echo-renderer/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Payload struct {
	DBGorm *gorm.DB
	Config *config.Config
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitRoute(payload *Payload) *echo.Echo {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("", controller.Index)

	user := e.Group("/user")
	user.GET("/registration-page", controller.RegistrationPage)
	user.POST("/login", controller.Login)
	user.POST("/register", controller.Register)
	user.GET("/dashboard/:uuid", controller.Dashboard)

	return e
}
