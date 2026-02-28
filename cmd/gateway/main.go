package main

 import (
  	"log"
  	"os"
  	"strconv"
  	"github.com/gin-gonic/gin"	
  )

  func envOrDefault(key, fallback string) string {
  	v := os.Getenv(key)
  	if v == "" {
  		return fallback
  	}
  	return v
  }

  func envIntOrDefault(key string, fallback int) int {
  	v := os.Getenv(key)
  	if v == "" {
  		return fallback
  	}
  	n, err := strconv.Atoi(v)
  	if err != nil {
  		return fallback
  	}
  	return n
  }

  func main() {
  	workerBaseURL := envOrDefault("WORKER_BASE_URL", "http://localhost:8081")
  	httpTimeoutMS := envIntOrDefault("HTTP_TIMEOUT_MS", 2000)

  	r := gin.Default()
  	r.GET("/health", func(c *gin.Context) {
  		c.JSON(200, gin.H{"status": "ok"})
  	})

  	log.Printf("gateway start addr=:8080 worker=%s timeout_ms=%d", workerBaseURL, httpTimeoutMS)
  	r.Run(":8080")
  }
