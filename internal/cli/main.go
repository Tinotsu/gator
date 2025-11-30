// Package cli
package cli

import(
	"context"
	"os"
	"time"
	"github.com/google/uuid"
	"fmt"
	"errors"
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

func Feeds(s *State, cmd Command) error {
	ctxt := context.Background()

	f, err := s.DB.GetFeeds(ctxt)
	if err != nil {
		config.HandleError(err)
		return err
	}

	l, err := s.DB.GetUsers(ctxt)
	if err != nil {
		config.HandleError(err)
		return err
	}

	for _, feed := range f {
		var username string
		for _, user := range l {
			if feed.UserID == user.ID {
				username = user.Name
			} 
		}
		fmt.Print("\n- ", feed.Name)
		fmt.Print("\n  url: ", feed.Url)
		fmt.Print("\n  user: ", username)
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

func AddFeed(s *State, cmd Command) error {
	if len(cmd.Arguments) < 4 {
		os.Exit(1)
		return errors.New("addfeed command take two additional arguments, the name of the feed and the url of the feed")
	}

	ctxt := context.Background()

	var userUUID uuid.UUID

	l, err := s.DB.GetUsers(ctxt)
	if err != nil {
		config.HandleError(err)
		return err
	}

	for _, user := range l {
		if s.Config.Username == user.Name {
			userUUID = user.ID
		}
	}
	
	var f *database.CreateFeedParams
	f = new(database.CreateFeedParams)

	f.ID = uuid.New()
	f.Name = cmd.Arguments[2]
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	f.Url  = cmd.Arguments[3]
	f.UserID = userUUID

	newFeed,err := s.DB.CreateFeed(ctxt, *f)
	if err != nil {
		config.HandleError(err)
		return err
	}

	feed, err := rss.FetchFeed(ctxt, newFeed.Url)
	if err != nil {
		config.HandleError(err)
		return err
	}
	fmt.Print(feed.Channel.Title)

	return nil
}
