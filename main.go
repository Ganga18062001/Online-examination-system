package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin_reg.html", gin.H{"Message": "Admin registration completed"})
	})
	engine.POST("/submit", func(ctx *gin.Context) {
		aadhar_id := ctx.PostForm("id")
		first_name := ctx.PostForm("first_name")
		last_name := ctx.PostForm("last_name")
		gender := ctx.PostForm("gender")
		address := ctx.PostForm("address")
		contact_no := ctx.PostForm("mobile_no")
		email_id := ctx.PostForm("email_id")
		isActive := ctx.PostForm("isactive")
		registration_date := ctx.PostForm("registration_date")
		password := ctx.PostForm("password")

		dbConnectionStr := "host=localhost user=postgres port=5434 sslmode=disable dbname=Project password=Ganga@123"

		db, err := sql.Open("postgres", dbConnectionStr)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "can not connect to database"})
			return
		}
		defer db.Close()
		if err := db.Ping(); err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "database not accessed"})
			return
		}

		insertQuery := `INSERT INTO pro.admin_reg(
		aadhar_id,First_name ,Last_name,Gender ,Address,Contact_no,isActive,Registration_date ,Password,email_id ) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
		statement, err := db.Prepare(insertQuery)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error Prearing the statenment"})
			log.Fatal(err)
			return
		}
		defer statement.Close()
		result, err := statement.Exec(aadhar_id, first_name, last_name, gender, address, contact_no, isActive, registration_date, password,email_id)

		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in executing query"})
			return
		}
		_, err = result.RowsAffected()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fetching Row Affected"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "SuccessFully Recod Inserted"})
	})

	engine.Run(":8080")

}
