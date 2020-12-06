package houses

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	// "encoding/json"
	// "strings"
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
		c.JSON(http.StatusOK, gin.H{
			"houses": s.FindAll(),
		})
	}
}

func addHouse(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h Houses
		if c.BindJSON(&h) == nil {
			// dec := json.NewDecoder(&h)
			// err := dec.Decode(&h)
			c.JSON(http.StatusOK, gin.H{
				"houses": s.AddHouse(h),
			})
		}
	}
}

func getById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"houses": s.FindByID(i),
		})
	}
}


func deleteById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"houses": s.DeleteByID(i),
		})
	}
}

func setSoldById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"houses": s.SetSoldByID(i),
		})
	}
}
// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
