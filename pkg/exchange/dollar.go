package exchange

type action int

const (
	Get action = iota
	Create
	Edit
	Delete
)

type ExMaper struct {
	ExMap map[string]string
}

func NewExMaper() *ExMaper {
	return &ExMaper{
		ExMap: make(map[string]string),
	}
}

type Response struct {
	Name  string
	Value string
	Error string
}

type Dollar struct {
	Name         string
	Value        string
	Action       action
	ExMaper      *ExMaper
	ResponseChan chan *Response
}

func (d *Dollar) Task() {
	switch d.Action {
	case Create:
		d.create()
	case Delete:
		d.delete()
	case Get:
		d.get()
	case Edit:
		d.edit()
	}
}

func (d *Dollar) create() {
	_, ok := d.ExMaper.ExMap[d.Name]
	resp := &Response{
		Name:  d.Name,
		Value: d.Value,
		Error: "",
	}
	if !ok {
		d.ExMaper.ExMap[d.Name] = d.Value
	} else {
		resp.Error = "this dollar is alreay exist"
	}
	d.ResponseChan <- resp
}

func (d *Dollar) get() {
	value, ok := d.ExMaper.ExMap[d.Name]
	resp := &Response{
		Name:  d.Name,
		Value: value,
		Error: "",
	}
	if !ok {
		resp.Error = "this dollar is not exist"
	}
	d.ResponseChan <- resp
}
func (d *Dollar) delete() {
	_, ok := d.ExMaper.ExMap[d.Name]
	resp := &Response{
		Name:  d.Name,
		Value: d.Value,
		Error: "",
	}
	if ok {
		delete(d.ExMaper.ExMap, d.Name)
	} else {
		resp.Error = "this dollar is not exist"
	}
	d.ResponseChan <- resp
}

func (d *Dollar) edit() {
	_, ok := d.ExMaper.ExMap[d.Name]
	resp := &Response{
		Name:  d.Name,
		Value: d.Value,
		Error: "",
	}
	if ok {
		d.ExMaper.ExMap[d.Name] = d.Value
	} else {
		resp.Error = "this dollar is not exist"
	}
	d.ResponseChan <- resp
}
