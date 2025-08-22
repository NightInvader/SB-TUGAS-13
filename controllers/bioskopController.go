package controllers

import (
	"database/sql"

	"net/http"

	"fmt"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

var Tempat = []Bioskop{}
var DB *sql.DB

func AddBioskop(ctx *gin.Context) {

	var newTempat Bioskop
	if err := ctx.ShouldBindJSON(&newTempat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := DB.QueryRow(sqlStatement, newTempat.Nama, newTempat.Lokasi, newTempat.Rating).Scan(&newTempat.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newTempat)
}

func GetBioskop(ctx *gin.Context) {
	var results = []Bioskop{}
	sqlStatement := `SELECT * FROM bioskop `

	rows, err := DB.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var bioskop = Bioskop{}

		err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)

		if err != nil {
			panic(err)
		}

		results = append(results, bioskop)
	}

	fmt.Println("Bioskop Datas : ", results)
	ctx.JSON(http.StatusOK, results)
}

func GetBioskopByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var search Bioskop

	sqlStatement := `SELECT * FROM bioskop WHERE id = $1`

	err := DB.QueryRow(sqlStatement, id).Scan(&search.ID, &search.Nama, &search.Lokasi, &search.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Bioskop Datas : ", search)
	ctx.JSON(http.StatusOK, search)
}
func UpdateBioskop(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `
	UPDATE bioskop
	SET nama = $2, lokasi = $3, rating = $4 WHERE id = $1`

	var update Bioskop
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := DB.Exec(sqlStatement, id, update.Nama, update.Lokasi, update.Rating)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Data amount = ", count)
}

func DeleteBioskop(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `
	DELETE FROM bioskop WHERE id = $1`

	result, err := DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Data amount = ", count)

}
