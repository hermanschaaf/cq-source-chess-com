package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ResolveUsername(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Username)
}
