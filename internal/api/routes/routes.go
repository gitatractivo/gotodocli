package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitatractivo/gotodocli/internal/api/handlers"
	"github.com/gitatractivo/gotodocli/internal/storage/sqlite"
)

func SetupRoutes(router *gin.Engine){
	storage,err:=sqlite.NewSQLiteStorage("tasks.db")
	if err!=nil{
		panic(err)
	}
	
	handler:=handlers.NewTaskHandler(storage)

	v1:=router.Group("/v1")
	
	{
		taskRouter:=v1.Group("/tasks")
		{
			taskRouter.POST("",handler.CreateTask)
			taskRouter.GET("",handler.GetTasks)
			taskRouter.GET("/:id",handler.GetTask)
			taskRouter.PUT("/:id",handler.UpdateTask)
			taskRouter.DELETE("/:id",handler.DeleteTask)
			taskRouter.POST("/done/:id",handler.MarkTaskAsDone)
		}
	}

}
