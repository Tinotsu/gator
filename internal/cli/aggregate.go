package cli

import(
	"context"
	"fmt"
	"github.com/Tinotsu/gator/internal/database"
	"github.com/Tinotsu/gator/internal/config"
	"github.com/Tinotsu/gator/internal/rss"
	"time"
)

func ScrapeFeeds(s *State) error {
	ctx := context.Background()

	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		config.HandleError(err)
		return err
	}

	markParam := new(database.MarkFeedFetchedParams)
	markParam.ID = feed.ID
	markParam.LastFetchedAt.Time = time.Now()
	markParam.LastFetchedAt.Valid = true

	s.DB.MarkFeedFetched(ctx, *markParam)

	rssFeed, err := rss.FetchFeed(ctx, feed.Url)	
	if err != nil {
		config.HandleError(err)
		return err
	}

	for i := range rssFeed.Channel.Item {
		fmt.Print(string(rssFeed.Channel.Item[i].Title), "\n")
	}
	fmt.Print("\n\n")

	return nil
}


