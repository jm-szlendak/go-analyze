package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	"github.com/jm-szlendak/go-analyze/analyser"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func handleAnalyseFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.IndentedJSON(400, ErrorResponse{Message: "File must be provided"})
		return
	}

	reader, err := file.Open()
	if err != nil {
		c.IndentedJSON(500, ErrorResponse{Message: "Cant open file"})
		return
	}

	decoder := xml.NewDecoder(reader)

	result, err := analyser.Analyse(decoder)

	if err != nil {
		c.IndentedJSON(400, err.Error())
		return
	}

	c.IndentedJSON(200, result)

}

func main() {
	router := gin.Default()
	router.POST("/analyseFile", handleAnalyseFile)

	router.Run("localhost:5001")

}
