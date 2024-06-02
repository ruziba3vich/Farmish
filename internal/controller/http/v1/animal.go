package v1

import (
	"Farmish/internal/controller/http/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"Farmish/internal/entity"
	"Farmish/internal/usecase"
	"Farmish/pkg/logger"
)

type animalRoutes struct {
	t usecase.Animal
	l logger.Interface
}

func newAnimalRoutes(handler *gin.RouterGroup, t usecase.Animal, l logger.Interface) {
	r := &animalRoutes{t, l}

	h := handler.Group("/animal")
	{
		h.POST("/create", r.CreateAnimal)
		h.PUT("/update/:id", r.UpdateAnimal)
	}
}

// @Summary     Create Animal
// @Description Api for creating animal
// @ID          animal-create
// @Tags  	    animal
// @Accept      json
// @Produce     json
// @Param       request body models.CreateAnimalRequest true "Admin credentials for logging in"
// @Success     200 {object} entity.Animal
// @Failure     500 {object} response
// @Router      /animal/create [post]
func (r *animalRoutes) CreateAnimal(c *gin.Context) {
	var (
		body models.CreateAnimalRequest
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		r.l.Error(err, "http - v1 - create-animal")
		errorResponse(c, http.StatusBadRequest, "request body is not matching")
	}

	response, err := r.t.CreateAnimal(c.Request.Context(), &entity.Animal{
		ID:       "",
		Name:     body.Name,
		Weight:   body.Weight,
		IsHungry: body.IsHungry,
	})

	if err != nil {
		r.l.Error(err, "http - v1 - create-animal")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary     Update Animal
// @Description Api for updating animal
// @ID          animal-update
// @Tags  	    animal
// @Accept      json
// @Produce     json
// @Param id path string true "Animal ID"
// @Param       request body models.CreateAnimalRequest true "Admin credentials for logging in"
// @Success     200 {object} entity.Animal
// @Failure     500 {object} response
// @Router      /animal/update/{id} [put]
func (r *animalRoutes) UpdateAnimal(c *gin.Context) {
	var (
		body models.CreateAnimalRequest
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		r.l.Error(err, "http - v1 - update-animal")
		errorResponse(c, http.StatusBadRequest, "request body is not matching")
	}

	id := c.Param("id")
	response, err := r.t.UpdateAnimal(c.Request.Context(), &entity.Animal{
		ID:       id,
		Name:     body.Name,
		Weight:   body.Weight,
		IsHungry: body.IsHungry,
	})

	if err != nil {
		r.l.Error(err, "http - v1 - update-animal")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, response)
}
