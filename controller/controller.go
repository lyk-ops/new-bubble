package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-bubble/models"
)

/*
 url     --> controller  --> logic   -->    model(dao)
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateToDo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.Bind(&todo)
	// 2. 调用业务逻辑
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建失败",
		})
	} else {
		// 3. 返回结果
		c.JSON(http.StatusOK, gin.H{
			"msg": "创建成功",
		})
	}

}

func GetTodoList(c *gin.Context) {
	// 1. 调用业务逻辑
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.Bind(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
	}
	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除清单失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "删除清单成功"})
	}
}
