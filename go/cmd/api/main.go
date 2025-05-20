package main

import "github.com/gin-gonic/gin"

// --- エントリポイント --- //
func main() {
	r := gin.Default()

	// health check だけ定義
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// :8080 で待受
	_ = r.Run(":8080")
}
