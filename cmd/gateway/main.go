package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ZoomieMuffin/go-observability/internal/otel"
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
	if _, err := otel.InitTracerProvider(context.Background(), "gateway"); err != nil {
		log.Fatal(err)
	}
	if _, err := otel.InitMeterProvider(context.Background(), "gateway"); err != nil {
		log.Fatal(err)
	}

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
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Printf("close worker response body: %v", err)
			}
		}()

		if resp.StatusCode >= http.StatusInternalServerError {
			c.JSON(http.StatusBadGateway, gin.H{"error": "worker error"})
			return
		}

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Status(resp.StatusCode)
		_, _ = io.Copy(c.Writer, resp.Body)
	})

	log.Printf("gateway start addr=:8080 worker=%s timeout_ms=%d", workerBaseURL, httpTimeoutMS)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
