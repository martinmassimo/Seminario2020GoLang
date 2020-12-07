package houses

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/houses",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/houses/:ID",
		function: getById(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/houses",
		function: addHouse(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/houses/:ID",
		function: deleteById(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/houses/:ID",
		function: setSoldById(s),
	})
	return list
}

func getAll(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		houses,err := s.FindAll()
			if err != nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"error": "Bad Request",
				})
			return
			}
		c.JSON(http.StatusOK, gin.H{
			"houses": houses,
		})
	}
}

func addHouse(s Service) gin.HandlerFunc {
	var h Houses
	return func(c *gin.Context) {
		if c.BindJSON(&h) == nil {
			id,err := s.AddHouse(h)
			if err != nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"error": "Bad Request",
				})
			return
    }
		c.JSON(http.StatusOK, gin.H{
			"id": id,
			})
		}
	}
}

func getById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "id debe ser numero",
			})
		}
		var house Houses 
		house.Id = -1
		fmt.Println(house)
		house, err = s.FindByID(i)
		if (house.Id == 0) {
			c.JSON(http.StatusOK, gin.H{
				"houses": "No existe registro con el id solicitado",
			})		
			return	
		}
		c.JSON(http.StatusOK, gin.H{
			"houses": house,
		})
	}
}


func deleteById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
    	if err != nil {
        	c.JSON(http.StatusBadRequest,gin.H{
				"error": "id debe ser numero",
			})
    	    return
		}
		 
		rows, err := s.DeleteByID(i)
		if (rows == 0) {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "no se pudo eleminar",
				"err" : err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"borrado registro id": i,
		})
		return
	}
}

func setSoldById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "id debe ser numero",
			})
		}
		rows, err := s.SetSoldByID(i)
		if (rows == 0) {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "no se pudo actualizar como vendido",
				"err" : err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"actualizado registro id": i,
		})
		return
	}
}
// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
