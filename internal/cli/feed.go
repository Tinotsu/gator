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

	if len(f) < 1 {
		err = errors.New("there is no feed in the database")
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

func AddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 4 {
		err := errors.New("addfeed command take two additional arguments, the name of the feed and the url of the feed")
		config.HandleError(err)
		os.Exit(1)
		return err	
	}

	ctxt := context.Background()

	u := UserParam(s, ctxt)

	f := new(database.CreateFeedParams)

	f.ID = uuid.New()
	f.Name = cmd.Arguments[2]
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	f.Url  = cmd.Arguments[3]
	f.UserID = u.ID

	newFeed,err := s.DB.CreateFeed(ctxt, *f)
	if err != nil {
		config.HandleError(err)
		return err
	}

	cmdFeed := new(Command)
	cmdFeed.Arguments = append(cmdFeed.Arguments, cmd.Arguments[0],cmd.Arguments[1],cmd.Arguments[3])

	err = Follow(s, *cmdFeed, user)
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
