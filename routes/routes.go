package routes

import (
	"gopkg.in/gin-gonic/gin.v1"

	"themis/utils"
)

// Init initializes the routing.
func Init(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing auxiliary REST service routes..")
	initLocalRoutes(engine)
}

func initLocalRoutes(engine *gin.Engine) {
	engine.GET("/ping", keepAlive)
	engine.PATCH("/api/reorder", reorder)
	// fallback bogus /render endpoint, remove that when the core API gets sane
	engine.POST("/api/render", renderMarkdown)
	engine.GET("/api/filters", filters)
}

func keepAlive(c *gin.Context) {
	c.JSON(200, gin.H { "message": "n/a", })
}

// {"data":{"attributes":{"renderedContent":"\u003cp\u003eSee: \u003ca href=\"https://github.com/fabric8io/fabric8-planner/issues/537\" rel=\"nofollow\"\u003ehttps://github.com/fabric8io/fabric8-planner/issues/537\u003c/a\u003e\u003c/p\u003e\n"},"id":"a791cd56-3320-498f-b885-9d19555b8616","type":"rendering"}}
func renderMarkdown(c *gin.Context) {
	c.JSON(200, gin.H { 
		"data": gin.H { 
			"attributes": gin.H { 
				"renderedContent": "<b>Rendering markdown on server side is not supported by this API version.</b>",
			},
		},
		"id": "renderID",
		"type": "rendering",
	})
}

func reorder(c *gin.Context) {
	// { data: [ WORKITEM ], position: { direction: "SOMEDIRECTION", id: "SOMEOTHER_WI_ID" } }
	// WORKITEM only has id, type, attributes.version
	// Response: data[WORKITEM] with WORKITEM being updated WorkItem
	c.JSON(501, "Reordering using /reorder is not supported by this service.")
}

func filters(c *gin.Context) {
	c.JSON(200, gin.H { "data": []gin.H {
		gin.H { 
			"attributes": gin.H { 
				"description": "Filter by workitemtype",
				"query": "filter[workitemtype]={id}",
				"title": "Workitem type",
				"type": "workitemtypes",
			},
			"type": "filters",
		},
	}})
}
