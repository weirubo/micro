package template

var (
	DeliveryHttpHandlerSRV = `package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"{{.Dir}}/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string 
}

// ArticleHandler  represent the httphandler for article
type {{title .Alias}}Handler struct {
	AUsecase domain.{{title .Alias}}Usecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func New{{title .Alias}}Handler(e *echo.Echo, us domain.{{title .Alias}}Usecase) {
	handler := &{{title .Alias}}Handler{
		AUsecase: us,
	}
	e.GET("/{{.Alias}}s", handler.Fetch{{title .Alias}})
	e.POST("/{{.Alias}}s", handler.Store)
	e.GET("/{{.Alias}}s/:id", handler.GetByID)
	e.DELETE("/{{.Alias}}s/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
func (a *{{title .Alias}}Handler) Fetch{{title .Alias}}(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listAr, nextCursor, err := a.AUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	
	return c.JSON(http.StatusOK, listAr)
}

// GetByID will get article by given id
func (a *{{title .Alias}}Handler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.AUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *domain.{{title .Alias}}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *{{title .Alias}}Handler) Store(c echo.Context) (err error) {
	var {{.Alias}} domain.{{title .Alias}}
	err = c.Bind(&{{.Alias}})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&{{.Alias}}); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.AUsecase.Store(ctx, &{{.Alias}}
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, {{.Alias}})
}

// Delete will delete article by given param
func (a *{{title .Alias}}Handler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.AUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}`
	DeliveryRpcHandlerSRV = `package rpc

import (
	"context"

	"github.com/micro/go-micro/util/log"

	{{.Alias}} "{{.Alias}}/proto/{{.Alias}}"
)

type {{title .Alias}} struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *{{title .Alias}}) Call(ctx context.Context, req *{{.Alias}}.Request, rsp *{{.Alias}}.Response) error {
	log.Log("Received {{title .Alias}}.Call request")
	example := &domain.Example{
		Name: "test",
	}
	repo := mysql.NewMysqlExampleRepository(internal.MysqlConn())
	ucase := usecase.NewExampleUsecase(repo)
	err := ucase.Create(ctx, example)
	if err != nil {
		log.Errorf("Call() || req=%+v || example=%+v || err=%v", req, example, err)
		return err
	}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *{{title .Alias}}) Stream(ctx context.Context, req *{{.Alias}}.StreamingRequest, stream {{.Alias}}.{{title .Alias}}_StreamStream) error {
	log.Logf("Received {{title .Alias}}.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&{{.Alias}}.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *{{title .Alias}}) PingPong(ctx context.Context, stream {{.Alias}}.{{title .Alias}}_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&{{.Alias}}.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
`
)
