package services

import (
	"context"

	"github.com/agoblet/chesscompubapi"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/hermanschaaf/cq-source-chess-com/client"
)

func Archives() *schema.Table {
	return &schema.Table{
		Name:      "chess_com_archives",
		Resolver:  client.RetryOnError(fetchArchives),
		Columns:   []schema.Column{},
		Multiplex: client.UsernameMultiplex,
		Transform: client.TransformWithStruct(
			&chesscompubapi.Archive{},
			transformers.WithPrimaryKeys("Username", "Year", "Month"),
		),
	}
}

func fetchArchives(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	archives, err := cl.ChessComPubAPI.ListArchives(cl.Username)
	if err != nil {
		return err
	}
	res <- archives
	return nil
}
