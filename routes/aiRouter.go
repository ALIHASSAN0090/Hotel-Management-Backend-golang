package routes


func AIRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/chat/query" , controllers.AiQuery())
}