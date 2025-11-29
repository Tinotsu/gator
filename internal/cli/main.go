// Package cli
package cli

import(
	"errors"
	"github.com/Tinotsu/gator/internal/config"
)

type State struct {
	PtrConfig *config.Config
}

type Command struct {
	Name string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error {
	if cmd.Arguments == nil {
		return errors.New("login handler expects a single argument, the username")
	}
	return nil
}
