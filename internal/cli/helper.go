package cli

import(
	"context"
	"github.com/Tinotsu/gator/internal/database"
	"github.com/Tinotsu/gator/internal/config"
)

func UserParam (s *State, ctxt context.Context) database.User {
	var u database.User

	l, err := s.DB.GetUsers(ctxt)
	if err != nil {
		config.HandleError(err)
	}

	for _, user := range l {
		if s.Config.Username == user.Name {
			u = user
		}
	}

	return u
}

func FeedParam (url string, s *State, ctxt context.Context) database.Feed {
	var f database.Feed

	l, err := s.DB.GetFeeds(ctxt)
	if err != nil {
		config.HandleError(err)
	}

	for _, feed := range l {
		if url == feed.Url {
			f = feed
		}
	}

	return f
}
