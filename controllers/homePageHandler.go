package controllers

import (
	"final/initializers"
	"final/middlewares"
	"final/models"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := initializers.InitTemplates()
	headData := models.HeadData{HeadTitle: "Home", StyleName: "Home"}
	headerData := models.HeaderData{CurrentSite: "Home"}
	objectId, err := middlewares.ParseObjectIdFromJWT(r)
	if err == nil {
		headerData.ProfileID = objectId.Hex()
	}
	//
	data := models.PageData{
		HeaderData: headerData,
		HeadData:   headData,
	}
	err = tmpl.ExecuteTemplate(w, "Home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
