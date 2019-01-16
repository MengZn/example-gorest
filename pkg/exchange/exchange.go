package exchange

import (
	"fmt"

	"example-gorest/pkg/utils"

	restful "github.com/emicklei/go-restful"
)

type ExChanger struct {
	exMaper  *ExMaper
	recorder *Recorder
	pool     *utils.Pool
}

var maxWorker = 10
var maxQueue = 10

func NewExChanger() *ExChanger {
	p := utils.NewWorkPool(maxWorker, maxQueue)
	p.Start()
	return &ExChanger{
		exMaper:  NewExMaper(),
		recorder: NewRecorder(),
		pool:     p,
	}
}

func (ex *ExChanger) CreateExChange(request *restful.Request, response *restful.Response) {
	responseChan := make(chan *Response)
	dollar := &Dollar{
		Name:         request.PathParameter("name"),
		Value:        request.PathParameter("value"),
		Action:       Create,
		ExMaper:      ex.exMaper,
		ResponseChan: responseChan,
	}
	request.ReadEntity(dollar)
	ex.pool.Run(dollar)
	select {
	case resp := <-responseChan:
		if resp.Error != "" {
			fmt.Printf("Create %s:%s is ERROR %s!!\n", dollar.Name, dollar.Value, resp.Error)
		} else {
			fmt.Printf("Create %s:%s success\n", dollar.Name, dollar.Value)
		}
		response.WriteEntity(resp)
	}
}

func (ex *ExChanger) DelExChange(request *restful.Request, response *restful.Response) {
	responseChan := make(chan *Response)
	dollar := &Dollar{
		Name:         request.PathParameter("name"),
		Action:       Delete,
		ExMaper:      ex.exMaper,
		ResponseChan: responseChan,
	}
	ex.pool.Run(dollar)
	select {
	case resp := <-responseChan:
		if resp.Error != "" {
			fmt.Printf("Delete %s ERROR %s!!\n", dollar.Name, resp.Error)
		} else {
			fmt.Printf("Delete %s success\n", dollar.Name)
		}
		response.WriteEntity(resp)
	}
}

func (ex *ExChanger) GetExChange(request *restful.Request, response *restful.Response) {
	responseChan := make(chan *Response)
	dollar := &Dollar{
		Name:         request.PathParameter("name"),
		Action:       Get,
		ExMaper:      ex.exMaper,
		ResponseChan: responseChan,
	}
	ex.pool.Run(dollar)
	select {
	case resp := <-responseChan:
		if resp.Error != "" {
			fmt.Printf("Get %s ERROR %s!!\n", dollar.Name, resp.Error)
		} else {
			fmt.Printf("Get %s success value is %s \n", dollar.Name, dollar.Value)
		}
		response.WriteEntity(resp)
	}
}

func (ex *ExChanger) EditExChange(request *restful.Request, response *restful.Response) {
	responseChan := make(chan *Response)
	dollar := &Dollar{
		Name:         request.PathParameter("name"),
		Action:       Edit,
		ExMaper:      ex.exMaper,
		Recorder:     ex.recorder,
		ResponseChan: responseChan,
	}
	request.ReadEntity(dollar)
	ex.pool.Run(dollar)
	select {
	case resp := <-responseChan:
		if resp.Error != "" {
			fmt.Printf("Get %s ERROR %s!!\n", dollar.Name, resp.Error)
		} else {
			fmt.Printf("Edit %s success value is %s \n", dollar.Name, dollar.Value)
		}
		response.WriteEntity(resp)
	}
}
