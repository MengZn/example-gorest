package exchange

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/go-rest/pkg/utils"
)

type ExChanger struct {
	exMaper *ExMaper
	pool    *utils.Pool
}

var maxWorker = 10
var maxQueue = 10

func NewExChanger() *ExChanger {
	p := utils.NewWorkPool(maxWorker, maxQueue)
	p.Start()
	return &ExChanger{
		exMaper: NewExMaper(),
		pool:    p,
	}
}

func (ex *ExChanger) CreateExChange(request *restful.Request, response *restful.Response) {
	d := make(chan error)
	dollar := &Dollar{
		Name:    request.PathParameter("name"),
		Value:   request.PathParameter("value"),
		Action:  Create,
		ExMaper: ex.exMaper,
		ErrChan: d,
	}
	request.ReadEntity(dollar)
	ex.pool.Run(dollar)
	select {
	case err := <-d:
		if err != nil {
			fmt.Printf("Create %s:%s is ERROR %s!!\n", dollar.Name, dollar.Value, err)
			response.WriteHeader(http.StatusOK)
		} else {
			fmt.Printf("Create %s:%s success\n", dollar.Name, dollar.Value)
			response.WriteEntity(dollar)
			response.WriteHeader(http.StatusCreated)
		}
	}
}

func (ex *ExChanger) DelExChange(request *restful.Request, response *restful.Response) {
	d := make(chan error)
	dollar := &Dollar{
		Name:    request.PathParameter("name"),
		Action:  Delete,
		ExMaper: ex.exMaper,
		ErrChan: d,
	}
	ex.pool.Run(dollar)
	select {
	case err := <-d:
		if err != nil {
			fmt.Printf("Delete %s ERROR %s!!\n", dollar.Name, err)
			response.WriteHeader(http.StatusOK)
		} else {
			fmt.Printf("Delete %s success\n", dollar.Name)
			response.WriteHeader(http.StatusOK)
		}
	}
}
