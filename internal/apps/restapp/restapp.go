package restapp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/model"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/repo"
)

type RestApp struct {
	bookRepo repo.BookRepo
}

func NewRestApp() RestApp {
	return RestApp{bookRepo: repo.NewBookRepo()}
}

func (a *RestApp) Start() {
	router := gin.Default()
	router.POST("/books/add", a.addBook)
	router.GET("/books/list", a.listBooks)

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
