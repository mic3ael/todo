package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/mic3ael/todo/pkg/error"
	"github.com/mic3ael/todo/pkg/model"
)

func (h *Handler) createList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []model.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserId(c)
	
	if err != nil {
		return
	}
	
	lists, err := h.services.TodoList.GetAll(userId)
	
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context){
	userId, err := getUserId(c)
	
	if err != nil {
		return
	}
	
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}


	list, err := h.services.TodoList.GetById(userId, id)

	switch err.(type) {
	case nil:
		break;
	case *custom_error.NotFound:
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return;
	default:
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context){
	userId, err := getUserId(c)
	
	if err != nil {
		return
	}
	
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}


	var input model.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userId, id, input);
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteList(c *gin.Context){
	userId, err := getUserId(c)
	
	if err != nil {
		return
	}
	
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}