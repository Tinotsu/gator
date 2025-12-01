package cli

import(
	"context"
	"fmt"
	"github.com/google/uuid"
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

	for _, item := range rssFeed.Channel.Item {
		post := new(database.CreatePostParams)
		post.ID = uuid.New()
		post.CreatedAt = time.Now()
		post.UpdatedAt = time.Now()
		post.FeedID = feed.ID

		const shortForm = "2006-Jan-02"
		dateTime, err := time.Parse(shortForm, item.PubDate)
		post.PublishedAt = dateTime

		post.Title = item.Title
		post.Url = item.Link
		post.Description = item.Description

		_, err = s.DB.CreatePost(ctx, *post)
		if err != nil {
			fmt.Print("\n⚠️ error: ", err, "\n")
		}
	}
	fmt.Print(rssFeed.Channel.Title)
	fmt.Print("\n\n")

	return nil
}

