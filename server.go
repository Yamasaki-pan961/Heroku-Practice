package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
)

func main(){
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"Hello,Heroku!",
		})
	})

	r.Run(":8080")
	//r.Run(":" + os.Getenv("PORT"))
}