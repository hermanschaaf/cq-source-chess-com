package client

import "errors"

type Spec struct {
	Usernames []string `json:"usernames"`
}

func (s *Spec) Validate() error {
	if len(s.Usernames) == 0 {
		return errors.New("usernames cannot be empty")
	}
	return nil
}
