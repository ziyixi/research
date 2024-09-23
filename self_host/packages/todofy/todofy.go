package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"
)

var (
	requiredEnvs = []string{"DIDA_EMAIL", "GEMINI_API_KEY", "CMI_USER", "CMI_PASSWD", "PORT", "MJ_APIKEY_PUBLIC", "MJ_APIKEY_PRIVATE"}
)

func loadEnv() map[string]string {
	envs := make(map[string]string)
	unloadedEnvs := []string{}
	for _, env := range requiredEnvs {
		envs[env] = os.Getenv(env)
		if len(envs[env]) == 0 {
			unloadedEnvs = append(unloadedEnvs, env)
		}
	}
	if len(unloadedEnvs) > 0 {
		panic("The following env variables are not loaded: " + strings.Join(unloadedEnvs, ", "))
	}
	return envs
}

func main() {
	envs := loadEnv()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.RateLimitMiddleware())

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		envs["CMI_USER"]: envs["CMI_PASSWD"],
	}))

	// set up routes
	authorized.POST("/update_todo", func(c *gin.Context) {
		c.Set("envs", envs)
		updateTodoPost(c)
	})

	// start the server
	listenAddr := ":" + envs["PORT"]
	log.Printf("gin mode: %s\n", gin.Mode())
	log.Printf("Listening on %s\n", listenAddr)
	if err := r.Run(listenAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
