package news

import (
	"github.com/mmcdole/gofeed"
	"github.com/nexryai/archer"
	"github.com/nexryai/backdance/internal/model"
	"net/http"
)

// CommonFeedProxyService is an implementation of FeedProxyService
// using gofeed library to parse RSS/Atom/JSON feeds
type CommonFeedProxyService struct{}

func (c *CommonFeedProxyService) Fetch(url string) (*model.Feed, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	requester := archer.SecureRequest{
		Request: req,
		TimeOut: 10,
		MaxSize: 1024 * 1024 * 10,
	}

	resp, err := requester.Send()
	if err != nil {
		return nil, ErrUrlIsUnsafe
	}

	parser := gofeed.NewParser()
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		return nil, err
	} else if feed == nil {
		return nil, ErrFeedIsNil
	}

	res := &model.Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		Links:       feed.Links,
		UpdatedAt:   feed.UpdatedParsed,
		PublishedAt: feed.PublishedParsed,
		Authors:     make([]*model.Person, 0),
		Language:    feed.Language,
		ImageUrl:    feed.Image.URL,
		Copyright:   feed.Copyright,
		Items:       make([]*model.Item, 0),
	}

	for _, author := range feed.Authors {
		res.Authors = append(res.Authors, &model.Person{
			Name:  author.Name,
			Email: author.Email,
		})
	}

	for _, item := range feed.Items {
		res.Items = append(res.Items, &model.Item{
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			Links:       item.Links,
			UpdatedAt:   item.UpdatedParsed,
			PublishedAt: item.PublishedParsed,
			ImageUrl:    item.Image.URL,
		})
	}

	return res, nil
}
