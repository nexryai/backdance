package news

import (
	"errors"
	"github.com/nexryai/backdance/internal/model"
)

var (
	ErrUrlIsUnsafe = errors.New("requested address is not allowed")
	ErrFeedIsNil   = errors.New("parser returned nil")
)

type FeedProxyService interface {
	Fetch(url string) (*model.Feed, error)
}
