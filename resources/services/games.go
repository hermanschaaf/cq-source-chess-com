package services

import (
	"context"
	"fmt"
	"time"

	"github.com/agoblet/chesscompubapi"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/hermanschaaf/cq-source-chess-com/client"
)

func Games() *schema.Table {
	return &schema.Table{
		Name:      "chess_com_games",
		Resolver:  fetchGames,
		Multiplex: client.UsernameMultiplex,
		Columns: []schema.Column{
			{Name: "username", Type: arrow.BinaryTypes.String, PrimaryKey: true, Resolver: client.ResolveUsername},
		},
		Transform: client.TransformWithStruct(
			&chesscompubapi.Game{},
			transformers.WithPrimaryKeys("URL"),
		),
		IsIncremental: true,
	}
}

func fetchGames(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	archives, err := cl.ChessComPubAPI.ListArchives(cl.Username)
	if err != nil {
		return err
	}

	key := "username:" + cl.Username
	val, err := cl.Backend.GetKey(ctx, key)
	if err != nil {
		return err
	}
	var lastArchive time.Time
	if val != "" {
		lastArchive, err = time.Parse("2006-01", val)
		if err != nil {
			return fmt.Errorf("failed to parse state time: %w", err)
		}
	}

	for _, arch := range archives {
		date, err := time.Parse("2006-01", fmt.Sprintf("%d-%02d", arch.Year, arch.Month))
		if err != nil {
			return fmt.Errorf("failed to parse archive time: %w", err)
		}
		if !lastArchive.IsZero() && date.Before(lastArchive) {
			// if we have already processed this archive in a previous sync, skip it
			cl.Logger().Info().Str("key", key).Time("archive", date).Time("previous", lastArchive).Msg("skipping archive, already processed in previous sync")
			continue
		}
		err = client.RetryOnError(fetchGame)(ctx, meta, &schema.Resource{Item: arch}, res)
		if err != nil {
			return err
		}
	}

	// update state
	if len(archives) > 0 {
		lastArchive := archives[len(archives)-1]
		err := cl.Backend.SetKey(ctx, key, fmt.Sprintf("%d-%02d", lastArchive.Year, lastArchive.Month))
		if err != nil {
			return err
		}
	}
	return nil
}

func fetchGame(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	archive := resource.Item.(chesscompubapi.Archive)
	games, err := cl.ChessComPubAPI.ListGames(archive)
	if err != nil {
		return err
	}

	res <- games
	return nil
}
