package actions

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/panwenbin/ghttplog/databases"
	"github.com/panwenbin/ghttplog/settings"
	"github.com/panwenbin/ghttplog/structs/requests"
	"time"
)

func LogPost(c *gin.Context) {
	recvLog := requests.RecvLog{}
	err := c.BindJSON(&recvLog)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "格式错误",
			"err":  err.Error(),
		})
		return
	}

	h := recvLog.ToHttp()
	collection := databases.Mongo.Database(settings.MongodbDatabase).Collection(settings.MongodbCollection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, h)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "插入错误",
			"err":  err.Error(),
		})
		return
	}

	id := res.InsertedID
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "插入成功",
		"id":   id,
	})
}
