package routes

import (
	"gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	r.GET("/students", controllers.FindAllStudents)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/documentNumber/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.InsertStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
}
