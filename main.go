package main

import (
	pages "gopg/templates/pages"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func staticCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.URL.Path, "/assets/") {
			c.Header("Cache-Control", "public, max-age=86399")
		}

		c.Next()
	}
}

func setupEngine(engine *gin.Engine) {
	err := engine.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal(err)
	}

	engine.Use(staticCacheMiddleware())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Static("/assets", "./public")

	ginHtmlRenderer := engine.HTMLRender
	engine.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
}

func setupRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.Home())
		c.Render(http.StatusOK, r)
	})
}

func main() {
	engine := gin.Default()

	setupEngine(engine)
	setupRoutes(engine)

	engine.Run("0.0.0.0:3000")
}
