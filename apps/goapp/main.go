package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
    "time"
)

// HealthResponse structure for health endpoint
type HealthResponse struct {
    Status    string    `json:"status"`
    Timestamp time.Time `json:"timestamp"`
    Hostname  string    `json:"hostname"`
    Uptime    string    `json:"uptime"`
}

// InfoResponse structure for info endpoint
type InfoResponse struct {
    AppName     string `json:"app_name"`
    Version     string `json:"version"`
    Hostname    string `json:"hostname"`
    GoVersion   string `json:"go_version"`
    OS          string `json:"os"`
    Architecture string `json:"architecture"`
}

var startTime time.Time

func init() {
    startTime = time.Now()
}

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    
    html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Go Kubernetes Demo</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
        }
        .container {
            background: rgba(255, 255, 255, 0.1);
            padding: 30px;
            border-radius: 10px;
            backdrop-filter: blur(10px);
        }
        h1 { color: #fff; margin-bottom: 20px; }
        .info { 
            background: rgba(255, 255, 255, 0.2); 
            padding: 15px; 
            margin: 10px 0;
            border-radius: 5px;
        }
        .endpoint {
            background: rgba(0, 0, 0, 0.3);
            padding: 10px;
            margin: 5px 0;
            border-radius: 5px;
            font-family: 'Courier New', monospace;
        }
        a { color: #ffd700; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Go Kubernetes Demo Application</h1>
        
        <div class="info">
            <strong>Pod Hostname:</strong> %s<br>
            <strong>Go Version:</strong> %s<br>
            <strong>OS/Arch:</strong> %s/%s<br>
            <strong>Uptime:</strong> %s
        </div>

        <h2>Available Endpoints:</h2>
        <div class="endpoint">
            <a href="/health">GET /health</a> - Health check endpoint
        </div>
        <div class="endpoint">
            <a href="/ready">GET /ready</a> - Readiness check endpoint
        </div>
        <div class="endpoint">
            <a href="/info">GET /info</a> - Application info (JSON)
        </div>
        <div class="endpoint">
            <a href="/metrics">GET /metrics</a> - Simple metrics
        </div>
    </div>
</body>
</html>
    `, hostname, runtime.Version(), runtime.GOOS, runtime.GOARCH, time.Since(startTime).Round(time.Second))
    
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, html)
}

// Health check handler - for liveness probe
func healthHandler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    
    response := HealthResponse{
        Status:    "healthy",
        Timestamp: time.Now(),
        Hostname:  hostname,
        Uptime:    time.Since(startTime).Round(time.Second).String(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Readiness check handler - for readiness probe
func readyHandler(w http.ResponseWriter, r *http.Request) {
    // In real app, check database connections, dependencies, etc.
    hostname, _ := os.Hostname()
    
    response := map[string]interface{}{
        "ready":    true,
        "hostname": hostname,
        "message":  "Application is ready to serve traffic",
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Info handler - application information
func infoHandler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    
    response := InfoResponse{
        AppName:      "go-k8s-demo",
        Version:      "1.0.0",
        Hostname:     hostname,
        GoVersion:    runtime.Version(),
        OS:           runtime.GOOS,
        Architecture: runtime.GOARCH,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Metrics handler - simple metrics
func metricsHandler(w http.ResponseWriter, r *http.Request) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    metrics := map[string]interface{}{
        "goroutines":      runtime.NumGoroutine(),
        "memory_alloc_mb": m.Alloc / 1024 / 1024,
        "memory_sys_mb":   m.Sys / 1024 / 1024,
        "uptime_seconds":  time.Since(startTime).Seconds(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metrics)
}

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
        next(w, r)
        log.Printf("Completed in %v", time.Since(start))
    }
}

func main() {
    // Register handlers
    http.HandleFunc("/", loggingMiddleware(homeHandler))
    http.HandleFunc("/health", loggingMiddleware(healthHandler))
    http.HandleFunc("/ready", loggingMiddleware(readyHandler))
    http.HandleFunc("/info", loggingMiddleware(infoHandler))
    http.HandleFunc("/metrics", loggingMiddleware(metricsHandler))
    
    port := "8080"
    log.Printf("Starting Go server on port %s...", port)
    log.Printf("Access at http://localhost:%s", port)
    
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
