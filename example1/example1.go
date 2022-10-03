package example1

import (
	example2 "example2"
	"log"

	"github.com/citradigital/toldata"
	context "golang.org/x/net/context"
)

type Example1 struct {
}

func NewExample1() *Example1 {
	this := Example1{}
	return &this
}

func (this *Example1) Say(ctx context.Context, req *Hello) (*Empty, error) {
	client, err := toldata.NewBus(ctx, toldata.ServiceConfiguration{URL: "localhost:4222"})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	example2Svc := example2.NewExample2ToldataClient(client)
	defer client.Close()

	greeting, err := example2Svc.GetGreeting(ctx, &example2.Payload{Name: req.Name})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(greeting.Message)
	return nil, nil
}

func (this *Example1) ToldataHealthCheck(ctx context.Context, req *toldata.Empty) (*toldata.ToldataHealthCheckInfo, error) {
	return nil, nil
}
