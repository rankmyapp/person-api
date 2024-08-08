package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rpuglielli/person-api/internal/domain/person/entity"
	"github.com/rpuglielli/person-api/internal/domain/person/usecase"
	e "github.com/rpuglielli/person-api/internal/domain/response/entity"
	vo "github.com/rpuglielli/person-api/internal/domain/response/vo"
	"github.com/rpuglielli/person-api/pkg/errors"
)

type PersonHandler struct {
	useCase *usecase.PersonUseCase
}

func NewPersonHandler(useCase *usecase.PersonUseCase) *PersonHandler {
	return &PersonHandler{useCase: useCase}
}

func (h *PersonHandler) ListPersons(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	persons, total, err := h.useCase.FindAll(c.Request.Context(), page, pageSize)
	if err != nil {
		handleError(c, err)
		return
	}

	response := vo.NewPaginatedResponse(persons, page, pageSize, total, "/persons")
	c.JSON(http.StatusOK, response)
}

func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		handleError(c, err)
		return
	}

	err := h.useCase.Create(c.Request.Context(), &person)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (h *PersonHandler) GetPerson(c *gin.Context) {
	id := c.Param("id")

	person, err := h.useCase.FindByID(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")

	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person.ID = id

	err := h.useCase.Update(c.Request.Context(), &person)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	err := h.useCase.Delete(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	response := e.NewDeletedResponse(id, "person")
	c.JSON(http.StatusOK, response)
}

func handleError(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	if errors.IsValidationError(err) {
		status = http.StatusBadRequest
	} else if errors.IsNotFoundError(err) {
		status = http.StatusNotFound
	} else if errors.IsConflictError(err) {
		status = http.StatusConflict
	}
	c.JSON(status, gin.H{"error": err.Error()})
}
