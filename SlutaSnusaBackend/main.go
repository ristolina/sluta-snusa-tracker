package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DBentry struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	QuitDate       string `form:"quitDate" json:"quitdate"`
	PrillorPerDay  int    `form:"prillorPerDay" json:"prillorperday"`
	PricePerDosa   int    `form:"pricePerDosa" json:"priceperdosa"`
	PrillorPerDosa int    `form:"prillorPerDosa" json:"prillorperdosa"`
}

// var slutasnusadb = []slutasnusaentry{
// 	{ID: "1", Username: "kristoffer", QuitDate: "2022-11-10 15:00:00", PrillorPerDay: 11, DosPrice: 19.99, PrillorPerDosa: 20},
// 	{ID: "2", Username: "tom", QuitDate: "2022-12-24 12:00:00", PrillorPerDay: 8, DosPrice: 29.99, PrillorPerDosa: 22},
// 	{ID: "3", Username: "danijel", QuitDate: "2021-12-28 23:00:00", PrillorPerDay: 15, DosPrice: 39.99, PrillorPerDosa: 24},
// }

func getSlutaSnusData(c *gin.Context) {
	var resArr []DBentry
	// Query for a value based on a single row.
	rows, err := db.Query("SELECT id, username, quitdate, prillorperday, priceperdosa, prillorperdosa from slutasnusa")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No data received.",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var res DBentry
		if err := rows.Scan(&res.Id, &res.Username, &res.QuitDate, &res.PrillorPerDay, &res.PricePerDosa, &res.PrillorPerDosa); err != nil {
			return
		}
		resArr = append(resArr, res)
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  "success",
		"message": resArr,
	})
}

func postSlutaSnusData(c *gin.Context) {
	var userData DBentry

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData.Username = c.Param("username")
	// Query for a value based on a single row.
	result, err := db.Exec("UPDATE slutasnusa SET quitdate = ?, prillorperday = ?, priceperdosa = ?, prillorperdosa = ? WHERE username = ?", userData.QuitDate, userData.PrillorPerDay, userData.PricePerDosa, userData.PrillorPerDosa, userData.Username)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fel!")
	}
	fmt.Println(result)
	return
	// var newslutasnus DBentry

	// Call BINDJSON to bind received JSON to newSlutaSnus
	// if err := c.BindJSON(&newslutasnus); err != nil {
	// 	return
	// }

	//slutasnusadb = append(slutasnusadb, newslutasnus)
}

func getSlutaSnusDataByName(c *gin.Context) {
	username := c.Param("username")
	var res DBentry
	// Query for a value based on a single row.
	row := db.QueryRow("SELECT id, username, quitdate, prillorperday, priceperdosa, prillorperdosa from slutasnusa where username = ?", username)
	if err := row.Scan(&res.Id, &res.Username, &res.QuitDate, &res.PrillorPerDay, &res.PricePerDosa, &res.PrillorPerDosa); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No user with username: " + username,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "success",
			"message": res,
		})
	}
	// c.JSON(http.StatusBadRequest, gin.H{
	// 	"result": "bad request",
	// })
	// jsonres, _ := json.Marshal(res)
	// fmt.Println(string(res2B))
	// fmt.Println(string(res.Id))

	// for _, a := range slutasnusadb {
	// 	if a.Username == username {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")
	router := gin.Default()
	router.GET("/slutasnusa", getSlutaSnusData)
	router.GET("slutasnusa/:username", getSlutaSnusDataByName)
	router.POST("/slutasnusa/:username", postSlutaSnusData)

	router.Run("0.0.0.0:8080")
}
