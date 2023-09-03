package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func UsernameMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	cl := meta.(*Client)
	usernames := cl.Spec.Usernames
	metas := make([]schema.ClientMeta, len(usernames))
	for i, username := range usernames {
		metas[i] = &Client{
			Backend:        cl.Backend,
			Spec:           cl.Spec,
			ChessComPubAPI: meta.(*Client).ChessComPubAPI,
			Username:       strings.ToLower(username),
			logger:         cl.logger,
		}
	}
	return metas
}
