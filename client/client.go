package client

import (
	"github.com/agoblet/chesscompubapi"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
)

type Client struct {
	logger         zerolog.Logger
	Spec           Spec
	ChessComPubAPI *chesscompubapi.Client
	Username       string
	Backend        state.Client
}

func (c *Client) ID() string {
	return "chess-com-" + c.Username
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(logger zerolog.Logger, s Spec, backend state.Client) (*Client, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	chess := chesscompubapi.NewClient()
	c := &Client{
		logger:         logger,
		Spec:           s,
		ChessComPubAPI: chess,
		Backend:        backend,
	}
	return c, nil
}
