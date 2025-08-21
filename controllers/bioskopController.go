package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Bioskop struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Lokasi string `json:"lokasi"`
	Rating string `json:"rating"`
}

var Tempat = []Bioskop{}

func AddBioskop(ctx *gin.Context) {
	var newTempat Bioskop

	if err := ctx.ShouldBindJSON(&newTempat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, newTempat.Nama, newTempat.Lokasi, newTempat.Rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Err().Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newTempat)
}
