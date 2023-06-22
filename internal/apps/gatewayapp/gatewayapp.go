package gatewayapp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/model"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/proto"
	"google.golang.org/grpc"
)

const HelloLimit = 5

type GatewayApp struct {
	client      proto.BookServiceClient
	helloClient proto.HelloServiceClient
	conn        *grpc.ClientConn
}

func NewGatewayApp() *GatewayApp {
	app := &GatewayApp{}
	return app
}

func (a *GatewayApp) Start() {
	err := a.setupGrpcClient()
	if err != nil {
		fmt.Printf("could not setup grcp client: %v\n", err)
		return
	}

	fmt.Println("Starting gateway")

	router := gin.Default()
	router.POST("/grpc/books/add", a.addBookGrpc)
	router.POST("/rest/books/add", a.addBookRest)

	router.GET("/rest/books/list", a.listBooks)
	router.GET("/grpc/books/list", a.listBooksGrpc)

	router.GET("/rest/hello", a.restHello)
	router.GET("/grpc/hello", a.grpcHello)

	port := "9092"
	if len(os.Getenv("GATEWAY_PORT")) > 0 {
		port = os.Getenv("GATEWAY_PORT")
	}

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: router,
	}

	s.SetKeepAlivesEnabled(false)
	s.ListenAndServe()
}

func (a *GatewayApp) Shutdown() {
	a.conn.Close()
}

func (a *GatewayApp) setupGrpcClient() error {
	var err error

	fmt.Println("Starting grpc client")

	opts := grpc.WithInsecure()
	serverHost := "localhost"
	serverPort := "50051"

	if port, ok := os.LookupEnv("GRPC_SERVER_PORT"); ok {
		serverPort = port
	}

	if service, ok := os.LookupEnv("GRPC_SERVICE_NAME"); ok {
		serverHost = service
	}

	servAddr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	fmt.Println("dialing", servAddr)

	a.conn, err = grpc.Dial(
		servAddr,
		opts,
	)
	if err != nil {
		fmt.Printf("could not connect: %v\n", err)
		return err
	}

	a.client = proto.NewBookServiceClient(a.conn)
	a.helloClient = proto.NewHelloServiceClient(a.conn)

	return nil
}

func (a *GatewayApp) doUnary(isbn int32, author, name string) string {
	req := &proto.BookRequest{
		Isbn:   isbn,
		Author: author,
		Name:   name,
	}

	_, err := a.client.AddBook(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling AddBook rpc : %v", err)
	}

	return ""
}

func (a *GatewayApp) addBookGrpc(c *gin.Context) {
	fmt.Println("got request - Gateway")

	var br proto.BookRequest
	if err := c.BindJSON(&br); err != nil {
		return
	}

	resp := a.doUnary(br.Isbn, br.Author, br.Name)
	c.String(http.StatusOK, resp)
}

func (a *GatewayApp) addBookRest(c *gin.Context) {
	var err error

	fmt.Println("Starting rest client")

	serverHost := "localhost"
	serverPort := "9090"

	if port, ok := os.LookupEnv("REST_SERVER_PORT"); ok {
		serverPort = port
	}

	if service, ok := os.LookupEnv("REST_SERVICE_NAME"); ok {
		serverHost = service
	}

	servAddr := fmt.Sprintf("http://%s:%s", serverHost, serverPort)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/books/add", servAddr), c.Request.Body)
	if err != nil {
		fmt.Println("error while creating rest request", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("error while sending rest request", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

func (a *GatewayApp) listBooks(c *gin.Context) {
	var err error

	fmt.Println("fetching books")

	serverHost := "localhost"
	serverPort := "9090"

	if port, ok := os.LookupEnv("REST_SERVER_PORT"); ok {
		serverPort = port
	}

	if service, ok := os.LookupEnv("REST_SERVICE_NAME"); ok {
		serverHost = service
	}

	servAddr := fmt.Sprintf("http://%s:%s", serverHost, serverPort)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/books/list", servAddr), nil)
	if err != nil {
		fmt.Println("error while creating rest request", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while sending rest request", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while reading books response", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var resp json.RawMessage
	_ = json.Unmarshal(bodyBytes, &resp)

	c.JSON(http.StatusAccepted, resp)
}

func (a *GatewayApp) listBooksGrpc(c *gin.Context) {
	res, err := a.client.ListBooks(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("error while calling ListBooks rpc : %v", err)
	}

	// var resp json.RawMessage
	// _ = json.Unmarshal(bodyBytes, &resp)

	c.JSON(http.StatusAccepted, res)
}

func (a *GatewayApp) restHello(c *gin.Context) {
	fmt.Println("fetching hello")

	serverHost := "localhost"
	serverPort := "9090"

	if port, ok := os.LookupEnv("REST_SERVER_PORT"); ok {
		serverPort = port
	}

	if service, ok := os.LookupEnv("REST_SERVICE_NAME"); ok {
		serverHost = service
	}

	servAddr := fmt.Sprintf("http://%s:%s", serverHost, serverPort)

	allResp := model.AllHelloResponse{
		Messages: make([]string, 0),
	}

	for {

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/hello", servAddr), nil)
		if err != nil {
			fmt.Println("error while creating rest request", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		req.Header.Add("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("error while sending rest request", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		defer res.Body.Close()
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("error while reading books response", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		var resp model.HelloResponse
		_ = json.Unmarshal(bodyBytes, &resp)

		respParts := strings.Split(resp.Message, "-")

		if len(respParts) != 2 {
			fmt.Println("invalid response from rest server", resp.Message)
			break
		}

		i, err := strconv.Atoi(respParts[1])
		if err != nil {
			fmt.Println("invalid response from rest server", resp.Message)
			break
		}

		allResp.Messages = append(allResp.Messages, resp.Message)

		if i == HelloLimit {
			break
		}
	}

	c.JSON(http.StatusOK, allResp.Messages)
}

func (a *GatewayApp) grpcHello(c *gin.Context) {
	t1 := time.Now()
	allResp := model.AllHelloResponse{
		Messages: make([]string, 0),
	}

	stream, err := a.helloClient.SayHello(context.Background(), &proto.HelloEmpty{})

	if err != nil {
		fmt.Printf("error while streaming hello : %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("error while streaming hello : %v", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		allResp.Messages = append(allResp.Messages, resp.Message)
	}

	fmt.Println("time taken", time.Since(t1).Seconds())

	c.JSON(http.StatusOK, allResp.Messages)
}
