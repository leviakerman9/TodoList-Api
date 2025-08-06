package api

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/leviakerman9/TODOLIST-API/internal/ethereum"
)

func CreateTask(c *gin.Context) {
	var req struct {
		Content  string `json:"content"`
		Assignee string `json:"assignee"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_,err := ethereum.Todo.CreateTask(ethereum.Auth, req.Content, common.HexToAddress(req.Assignee))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

func GetTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	content, status, assignee, err := ethereum.Todo.GetTask(nil, big.NewInt(int64(taskID)))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       taskID,
		"content":  content,
		"status":   status,
		"assignee": assignee.Hex(),
	})
}

func UpdateTask(c *gin.Context) {
	var req struct {
		ID         uint64 `json:"id"`
		NewContent string `json:"new_content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_,err := ethereum.Todo.UpdateTask(ethereum.Auth, big.NewInt(int64(req.ID)), req.NewContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	_,err := ethereum.Todo.DeleteTask(ethereum.Auth, big.NewInt(int64(taskID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
