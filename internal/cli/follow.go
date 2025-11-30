package cli

import (
	"errors"
	"time"
	"fmt"
	"context"
	"github.com/Tinotsu/gator/internal/database"
	"github.com/Tinotsu/gator/internal/config"
	"github.com/google/uuid"
)

func Follow(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 3 {
		err := errors.New("follow commands require an url as third arguments")
		config.HandleError(err)
		return err
	}

	ctxt := context.Background()

	feed := FeedParam(cmd.Arguments[2], s, ctxt)

	f := new(database.CreateFeedFollowParams)

	f.ID = uuid.New()
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	f.FeedID = feed.ID
	f.UserID = user.ID

	_, err := s.DB.CreateFeedFollow(ctxt, *f)
	if err != nil {
		config.HandleError(err)
		return err
	}

	fmt.Print("\n=== follow added ===")
	fmt.Print("\n  feed: ", feed.Name)
	fmt.Print("\n  user: ", user.Name)

	return nil
}

func Following(s *State, cmd Command, user database.User) error {
	ctxt := context.Background()

	userID := user.ID

	feeds,err := s.DB.GetFeedFollowsForUser(ctxt, userID)
	if err != nil {
		config.HandleError(err)
		return err
	}

	fmt.Printf("feeds followed by %s :\n", user.Name)
	for i := range feeds {
		fmt.Print("\n  - ", feeds[i].FeedName)
	}
	
	return nil
}
