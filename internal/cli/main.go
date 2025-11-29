// Package cli
package cli

import(
	"context"
	"fmt"
	"github.com/Tinotsu/gator/internal/config"
	"github.com/Tinotsu/gator/internal/rss"
	"github.com/Tinotsu/gator/internal/database"
)

type State struct {
	DB *database.Queries
	Config *config.Config
}

type Command struct {
	Name string
	Arguments []string
}

func NewState () *State {
	s := State{}
	return &s
}

func Reset(s *State, cmd Command) error {
	ctxt := context.Background()
	err := s.DB.DeleteUsers(ctxt)
	if err != nil {
		return err
	}
	fmt.Print("table user reset")
	return nil
}

func Users(s *State, cmd Command) error {
	ctxt := context.Background()
	l, err := s.DB.GetUsers(ctxt)
	if err != nil {
		return err
	}

	for _, user := range l {
		if s.Config.Username == user.Name {
			fmt.Print("- ", user.Name, " (current)\n")
		} else {
			fmt.Print("- ", user.Name, "\n")
		}
	}

	return nil
}

func RSS(s *State, cmd Command) error {
	ctxt := context.Background()
	rss.FetchFeed(ctxt, "")
	return nil
}
