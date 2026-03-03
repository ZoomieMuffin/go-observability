package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
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

	r.POST("/work", func(c *gin.Context) {
		client := &http.Client{
			Timeout: time.Duration(httpTimeoutMS) * time.Millisecond,
		}

		resp, err := client.Post(workerBaseURL+"/work", "application/json", nil)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "worker unavailable"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusInternalServerError {
			c.JSON(http.StatusBadGateway, gin.H{"error": "worker error"})
			return
		}

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Status(resp.StatusCode)
		_, _ = io.Copy(c.Writer, resp.Body)
	})

	log.Printf("gateway start addr=:8080 worker=%s timeout_ms=%d", workerBaseURL, httpTimeoutMS)
	r.Run(":8080")
}
