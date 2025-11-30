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
    ctx := context.Background()

    if err := s.DB.DeleteFeeds(ctx); err != nil {
		config.HandleError(err)
        return err
    }

    if err := s.DB.DeleteUsers(ctx); err != nil {
		config.HandleError(err)
        return err
    }

    fmt.Println("Database reset successfully!")
    return nil
}

func Users(s *State, cmd Command) error {
	ctxt := context.Background()
	l, err := s.DB.GetUsers(ctxt)
	if err != nil {
		config.HandleError(err)
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
	feed, err := rss.FetchFeed(ctxt, "https://www.wagslane.dev/index.xml")
	if err != nil {
		config.HandleError(err)
		return err
	}
	fmt.Print(feed)
	return nil
}
