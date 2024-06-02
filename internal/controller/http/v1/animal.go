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
	}
}

// @Summary     Create Animal
// @Description Api for creating animal
// @ID          animal-create
// @Tags  	    animal
// @Accept      json
// @Produce     json
// @Param       request body models.LoginRequest true "Admin credentials for logging in"
// @Success     200 {object} models.LoginResponse
// @Failure     500 {object} response
// @Router      /admin/login [post]
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

//
//// @Summary     Translate
//// @Description Translate a text
//// @ID          do-translate
//// @Tags  	    translation
//// @Accept      json
//// @Produce     json
//// @Param       request body doTranslateRequest true "Set up translation"
//// @Success     200 {object} entity.Translation
//// @Failure     400 {object} response
//// @Failure     500 {object} response
//// @Router      /translation/do-translate [post]
//func (r *translationRoutes) doTranslate(c *gin.Context) {
//	var request doTranslateRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusBadRequest, "invalid request body")
//
//		return
//	}
//
//	translation, err := r.t.Translate(
//		c.Request.Context(),
//		entity.Translation{
//			Source:      request.Source,
//			Destination: request.Destination,
//			Original:    request.Original,
//		},
//	)
//	if err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusInternalServerError, "translation service problems")
//
//		return
//	}
//
//	c.JSON(http.StatusOK, translation)
//}
