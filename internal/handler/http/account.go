package handler

import (
	"currency-operations/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entity.Account

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Account.Create(userId, input)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllAccountsResponce struct {
	Data []entity.Account `json:"accounts"`
}

func (h *Handler) GetAllAccounts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	accounts, err := h.services.Account.GetAll(userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAccountsResponce{
		Data: accounts,
	})

}

func (h *Handler) GetAccountById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	account, err := h.services.Account.GetById(userId, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) Deposit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateAccount

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Account.Deposit(userId, id, *input.Balance)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
