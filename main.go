package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tecbot/gorocksdb"
)

func main() {
	opt := gorocksdb.NewDefaultOptions()
	opt.SetCreateIfMissing(true)

	d, err := gorocksdb.OpenDb(opt, "test")
	wo := gorocksdb.NewDefaultWriteOptions()
	ro := gorocksdb.NewDefaultReadOptions()

	r := gin.Default()
	r.GET("/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		val, err := d.Get(ro, []byte(key))

		if err != nil {
			c.JSON(400, gin.H{
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "ok",
				"val":     string(val.Data()),
			})
		}
	})

	r.POST("/key/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		val := c.Param("value")
		err = d.Put(wo, []byte(key), []byte(val))

		if err != nil {

			c.JSON(400, gin.H{
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		}
	})

	r.DELETE("/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		err = d.Delete(wo, []byte(key))

		if err != nil {
			c.JSON(400, gin.H{
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run()
}
