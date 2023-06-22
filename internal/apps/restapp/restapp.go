package restapp

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/model"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/repo"
)

const HelloLimit = 5

type RestApp struct {
	bookRepo repo.BookRepo
	ctr      int
}

func NewRestApp() RestApp {
	return RestApp{bookRepo: repo.NewBookRepo(), ctr: 0}
}

func (a *RestApp) Start() {
	router := gin.Default()
	router.POST("/books/add", a.addBook)
	router.GET("/books/list", a.listBooks)
	router.GET("/hello", a.helloProcessor)

	port := "9090"
	if len(os.Getenv("REST_SERVER_PORT")) > 0 {
		port = os.Getenv("REST_SERVER_PORT")
	}

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: router,
	}

	s.SetKeepAlivesEnabled(false)
	s.ListenAndServe()
}

func (a *RestApp) Stop() {

}

func (a *RestApp) addBook(c *gin.Context) {
	var book *model.Book
	if err := c.BindJSON(&book); err != nil {
		fmt.Println("error while binding", err)
		c.JSON(http.StatusInternalServerError, nil)
	}

	a.bookRepo.StoreBook(book)
	c.JSON(http.StatusAccepted, nil)
}

func (a *RestApp) listBooks(c *gin.Context) {
	books := a.bookRepo.GetBooks()
	c.JSON(http.StatusOK, books)
}

func (a *RestApp) helloProcessor(c *gin.Context) {
	time.Sleep(1 * time.Second)
	resp := model.HelloResponse{
		Message: fmt.Sprintf("Hello-%d", a.ctr),
	}

	if a.ctr == HelloLimit {
		a.ctr = 0
	}

	a.ctr++

	c.JSON(http.StatusOK, resp)
}
