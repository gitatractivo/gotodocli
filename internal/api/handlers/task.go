package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/gitatractivo/gotodocli/internal/storage/sqlite"
)

type TaskHandler struct{
	storage *sqlite.SQLiteStorage
}

func NewTaskHandler(storage *sqlite.SQLiteStorage) *TaskHandler{
	return &TaskHandler{
		storage: storage,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context){
	var task models.Task
	log.Println("Creating task")
	if err:=c.ShouldBindJSON(&task); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	log.Println("Task created",task)
	if err:=h.storage.CreateTask(&task); err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	log.Println("Task created",task)

	c.JSON(http.StatusCreated,task)
}

func (h *TaskHandler) GetTasks(c *gin.Context){
	tasks,err:=h.storage.GetAllTasks()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context){
 idParam:=c.Param("id")
 id,err := strconv.ParseUint(idParam,10,64)
 if err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return
 }
 task,err:=h.storage.GetTaskById(uint(id))

 if err!=nil{
	c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
	return
 }

 c.JSON(http.StatusOK,task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context){
	idParam:=c.Param("id")
	id,err:=strconv.ParseUint(idParam,10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	var updatedTask models.Task
	
	existingTask,err:=h.storage.GetTaskById(uint(id))
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}	

	updatedTask.ID=existingTask.ID
	if err:=c.ShouldBindJSON(&updatedTask); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err:=h.storage.UpdateTask(&updatedTask); err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,updatedTask)
	
}
 
func (h *TaskHandler) DeleteTask(c *gin.Context){
	idParam:=c.Param("id")
	id,err:=strconv.ParseUint(idParam,10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err:=h.storage.DeleteTask(uint(id)); err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"Task deleted successfully"})
	

}

func (h *TaskHandler) MarkTaskAsDone(c *gin.Context){
	log.Println("Marking task as done")
	idParam:=c.Param("id")
	id,err:=strconv.ParseUint(idParam,10,64)
	log.Println("Marking task as done",id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	task,err:=h.storage.GetTaskById(uint(id))
	log.Println("Task found",task)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}

	task.Completed=true
	if err:=h.storage.UpdateTask(task); err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,task)
}


 
