package cli

import(
	"github.com/Tinotsu/gator/internal/database"
	"context"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func (s *State, cmd Command) error {
		ctxt := context.Background()
		user, err := s.DB.GetUser(ctxt, s.Config.Username)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
