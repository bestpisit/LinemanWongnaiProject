package main

import (
	"covid-19/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/covid/summary", func(c *gin.Context) {
		covidSummary, error := utils.Covid19SummaryFromURL("https://static.wongnai.com/devinterview/covid-cases.json")
		if error != nil {
			c.JSON(500, gin.H{
				"Error: ": "Internal Server Error",
			})
		}
		c.JSON(200, gin.H{
			"Province": covidSummary.Province,
			"AgeGroup": covidSummary.AgeGroup,
		})
	})
	r.Run()
}

/* the GIN framework use utilities form utils folder */
/* Tests cases is in tests folder */