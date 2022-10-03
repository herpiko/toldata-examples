package example2

import (
	"log"

	"github.com/citradigital/toldata"
	context "golang.org/x/net/context"
)

type Example2 struct {
}

func NewExample2() *Example2 {
	this := Example2{}
	return &this
}

func (this *Example2) GetGreeting(ctx context.Context, req *Payload) (*Response, error) {
	log.Println("Payload: " + req.Name)
	return &Response{
		Message: "Hello, " + req.Name,
	}, nil
}

func (this *Example2) ToldataHealthCheck(ctx context.Context, req *toldata.Empty) (*toldata.ToldataHealthCheckInfo, error) {
	return nil, nil
}
