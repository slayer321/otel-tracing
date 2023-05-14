package web

import (
	"context"
	"net/http"
	"test/db"
	"test/log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type WebDB struct {
	db.Helper
}

func (hp WebDB) SetupWebServer(client *mongo.Client) {

	r := gin.Default()
	r.Use(otelgin.Middleware("todo-server"))

	r.GET("/todo", func(c *gin.Context) {
		hp.DB = client
		collection := hp.GetCollection("todo", "todos")

		cur, findErr := collection.Find(c.Request.Context(), bson.D{})
		if findErr != nil {
			c.AbortWithError(500, findErr)
		}

		result := make([]interface{}, 0)
		curErr := cur.All(c, &result)
		if curErr != nil {
			c.AbortWithError(500, curErr)
		}
		c.JSON(http.StatusOK, result)

	})
	err := client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	log.Log.Println("Server is running on port 8080")
	_ = r.Run(":8080")
}
