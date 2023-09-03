package client

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// RetryOnError will run the given resolver function and retry on rate limit exceeded errors
// or other retryable errors (like internal server errors) after waiting some amount of time.
func RetryOnError(resolver schema.TableResolver) schema.TableResolver {
	retries := 0
	maxRetries := 10
	backoff := 10 * time.Second
	f := func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		c := meta.(*Client)
		var err error
		for err = resolver(ctx, meta, parent, res); retries < maxRetries; err = resolver(ctx, meta, parent, res) {
			if shouldRetry(err) {
				retryAfter := time.Duration((0.9 + rand.Float64()*0.2) * float64(backoff))
				retries++
				c.logger.Info().Msgf("Got retryable error (%v), retrying in %.2fs (%d/%d)", err.Error(), retryAfter.Seconds(), retries, maxRetries)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(retryAfter):
					continue
				}
			}
			return err
		}
		return err
	}
	return f
}

func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "429") || strings.Contains(err.Error(), "500") {
		return true
	}
	return false
}
