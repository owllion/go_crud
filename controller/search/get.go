package searchRoute

import (
	"context"
	"log"
	"net/http"
	"practice/gRPC/client"
	"practice/gRPC/service/search"
	"strconv"

	"github.com/gin-gonic/gin"
)



func Search(c *gin.Context) {
	searchClient := search.NewSearchServiceClient(client.ClientConn)
	
	query := c.Query("query")
	pageNumber,_ := strconv.Atoi(c.Query("pageNumber"))

	response, err := searchClient.Search(context.Background(), &search.SearchRequest{Query: query, PageNumber: int32(pageNumber)})


	if err != nil {
		log.Printf("Error calling search service: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"returned_msg": response.ReturnedMsg,
		"page_number": response.PageNumber,
	})
}