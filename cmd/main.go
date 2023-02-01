//!

package main

import (
	"library-api/api"
	"library-api/service"
	hotreload "library-api/util/hot_reload"

	"github.com/gin-gonic/gin"
)

func main() {
	hotreload.Go("../", func() {
		service.MustResetTables()
		g := gin.Default()
		api.InitGroup(g)
		g.Run()
	})
}
