package robots

import (
	"context"
	"errors"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
)

type Client struct {
	conn   *grpc.ClientConn
	client proto_game.GameServiceClient
	robot  Robot
}

func NewClient(conn *grpc.ClientConn, robot Robot) *Client {
	x := &Client{
		conn:   conn,
		robot:  robot,
		client: proto_game.NewGameServiceClient(conn),
	}
	return x
}

func Run(ctx context.Context, conn *grpc.ClientConn, token string, robot Robot) error {
	c := NewClient(conn, robot)
	return c.Run(ctx, token)
}

func (c *Client) Run(ctx context.Context, token string) error {
	pcontext := metadata.NewOutgoingContext(ctx, metadata.MD{
		"token": []string{token},
	})
	stream, err := c.client.Connect(pcontext)
	if err != nil {
		return err
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		s := req.GetSuggest()
		answerExpected := s != nil && !s.GetCanceled()
		res := c.robot.Request(req)
		if res == nil {
			if answerExpected {
				return errors.New("Expected robot to answer")
			}
		} else {
			if !answerExpected {
				return errors.New("Answer unexpected")
			}
			res.SuggestId = s.GetSuggestId()
			err = stream.Send(res)
			if err != nil {
				return err
			}
		}
	}
}
