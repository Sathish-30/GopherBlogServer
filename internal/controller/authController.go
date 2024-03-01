package controller

import (
	"net/http"

	"github.com/sathish-30/GopherBlogServer/internal/database"
	"github.com/sathish-30/GopherBlogServer/internal/model"
	"github.com/sathish-30/GopherBlogServer/internal/util"
)

func Register(w http.ResponseWriter, r *http.Request) error {
	data, err := util.ConvertJson(r)
	data = model.User(data)
	if err != nil {
		return err
	}
	var existingUser model.User
	database.DB.First(&existingUser, "email = ?", data.Email)
	if existingUser.ID == 0 {
		result := database.DB.Create(&data)
		if result.Error != nil {
			util.WriteJson(w, http.StatusAccepted, model.Response{
				Message: "Unable to add user to the database",
			})
		} else {
			util.WriteJson(w, http.StatusAccepted, model.Response{
				Message: "User Registered",
			})
		}
	} else {
		util.WriteJson(w, http.StatusAccepted, model.Response{
			Message: "User Already existed",
		})
	}
	return nil
}
