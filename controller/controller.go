package controller

import (
	config "learn-echo-renderer/Config"
	"learn-echo-renderer/constants"
	"learn-echo-renderer/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func Dashboard(c echo.Context) error {

	uuid := c.Param("uuid")

	user := model.User{}
	if err := config.DB.First(&user).Where("uuid = ?", uuid).Error; err != nil {
		return c.String(http.StatusNotFound, "Tidak ditemukan")
	}

	return c.Render(http.StatusOK, "dashboard", echo.Map{
		"username": user.Username,
	})
}
func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	err := config.DB.First(&user).
		Where("username = ? AND password = ?", user.Username, user.Password).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Render(http.StatusNotFound, "confirmation-page", echo.Map{
				"message":       "user " + username + " tidak ditemukan",
				"buttonMessage": "kembali",
				"success":       false,
				"login":         true,
				"data":          model.User{},
			})
		} else {
			return c.Render(http.StatusInternalServerError, "confirmation-page", echo.Map{
				"message":       "kesalahan server, silahkan coba lagi",
				"buttonMessage": "kembali",
				"success":       false,
				"login":         true,
				"data":          model.User{},
			})
		}
	}

	return c.Render(http.StatusOK, "confirmation-page", echo.Map{
		"message":       "halo! selamat datang " + username,
		"buttonMessage": "lanjut",
		"success":       true,
		"login":         true,
		"data": model.User{
			Username: username,
			UUID:     user.UUID,
		},
	})
}

func RegistrationPage(c echo.Context) error {
	return c.Render(http.StatusOK, "registration-page", nil)
}

func Register(c echo.Context) error {

	var newUser model.User
	newUser.Username = c.FormValue("username")
	newUser.Password = c.FormValue("password")

	if err := config.DB.Find(&newUser).Where("username = ?", newUser.Username).Error; err != nil {
		if newUser.UUID != constants.EmptyString {
			c.Render(http.StatusConflict, "confirmation-page", echo.Map{
				"message":       "Maaf! nama pengguna yang anda pilih sudah digunakan, silahkan gunakan nama pengguna yang lain",
				"buttonMessage": "kembali",
				"login":         false,
				"success":       false,
			})
		}
	}

	newUser.UUID = uuid.NewString()

	if err := config.DB.Create(&newUser).Error; err != nil {
		return c.Render(http.StatusInternalServerError, "confirmation-page", echo.Map{
			"message":       "Maaf ada kesalahan server, silahkan coba lagi nanti.",
			"buttonMessage": "Kembali",
			"login":         false,
			"success":       false,
		})
	}

	return c.Render(http.StatusOK, "confirmation-page", echo.Map{
		"message":       "Pendaftaran berhasil!",
		"buttonMessage": "Login",
		"login":         false,
		"success":       true,
	})
}
