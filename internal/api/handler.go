package api

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/leviakerman9/TODOLIST-API/internal/ethereum"
)

type APIHandler struct {
	Eth *ethereum.Todolist
}

func (h *APIHandler) CreateTask(c *gin.Context) {
	var req struct {
		Content  string `json:"content"`
		Assignee string `json:"assignee"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Eth.Todo.CreateTask(h.Eth.Auth, req.Content, common.HexToAddress(req.Assignee))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

func (h *APIHandler) GetTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	content, status, assignee, err := h.Eth.Todo.GetTask(nil, big.NewInt(int64(taskID)))
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

func (h *APIHandler) UpdateTask(c *gin.Context) {
	var req struct {
		ID         uint64 `json:"id"`
		NewContent string `json:"new_content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Eth.Todo.UpdateTask(h.Eth.Auth, big.NewInt(int64(req.ID)), req.NewContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (h *APIHandler) DeleteTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	_, err := h.Eth.Todo.DeleteTask(h.Eth.Auth, big.NewInt(int64(taskID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
