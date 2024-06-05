package v1

import (
	"Farmish/internal/controller/http/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp"
	"net/http"
	"time"

	"Farmish/internal/entity"
	"Farmish/internal/usecase"
	"Farmish/pkg/logger"
)

type animalRoutes struct {
	t usecase.Animal
	l logger.Interface
}

type animalWebsocketRoutes struct {
	t usecase.Animal
	l logger.Interface
}

func newAnimalRoutes(handler *gin.RouterGroup, t usecase.Animal, l logger.Interface) {
	r := &animalRoutes{t, l}
	//rws := &animalWebsocketRoutes{}

	h := handler.Group("/animal")
	{
		h.POST("/create", r.CreateAnimal)
		h.PUT("/update/:id", r.UpdateAnimal)
		h.GET("/ws/notify", r.CheckHungryStatusOfAnimal)
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

func (rws *animalRoutes) CheckHungryStatusOfAnimal(c *gin.Context) {
	// Upgrade HTTP request to WebSocket connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		rws.l.Error(err, "websocket upgrade failed")
		return
	}
	defer conn.Close()

	// Create a timer for notification checks
	// timer := time.NewTimer(time.Second * 5)

	// for {
	// 	select {
	// 	case <-timer.C:
	// 		// Every 5 seconds, check animal status and send notification
	// 		result, err := rws.t.NotifyAnimalStatus(c.Request.Context())
	// 		if err != nil {
	// 			rws.l.Error(err, "http - v1 - notify status animal")
	// 			errorResponse(c, http.StatusInternalServerError, "database problems")
	// 			continue
	// 		}
	// 		err = conn.WriteJSON(result) // Send notification data as JSON
	// 		if err != nil {
	// 			rws.l.Error(err, "failed to send animal status update")
	// 		}
	// 		// Reset the timer for the next check
	// 		timer.Reset(time.Second * 5)
	//   case msgType, msg, ok := <-conn.ReadMessage():
	// 	// Handle incoming messages from client (optional)
	// 	if !ok {
	// 	  // Connection closed, handle cleanup
	// 	  return
	// 	}
	// Process received message based on msgType and msg
	// case err := <-conn.ReadPong():
	// 	if err != nil {
	// 		// Handle pong error
	// 	}
	// }
	// }

	for {
		result, err := rws.t.NotifyAnimalStatus(context.Background())
		if err != nil {
			rws.l.Error(err, "http - v1 - notify status animal")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		if result != nil {
			err := conn.WriteJSON(result)
			if err != nil {
				pp.Println(err)
			}
			if err != nil {
				rws.l.Error(err, "failed to send animal status update")
			}
		} else {
			pp.Println("No animal status update available")
		}

		time.Sleep(time.Second * 5)

	}
}
