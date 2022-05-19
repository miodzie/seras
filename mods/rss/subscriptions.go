package rss

import (
	"math/rand"
	"strings"
)

type Subscription struct {
	Id        uint64
	FeedId    uint64
	User      string
	Keywords  string
	Channel   string
	Feed      *Feed
	Seen      string          // Item.GUID comma separated
	SeenItems map[string]bool // [guid]bool
}

func (s *Subscription) See(item Item) {
	s.makeSeenMap()
	if _, seen := s.SeenItems[item.GUID]; !seen {
		s.SeenItems[item.GUID] = true
		s.Seen = ""
		var keys []string
		for k := range s.SeenItems {
			keys = append(keys, k)
		}
		s.Seen = strings.Join(keys, ",")
	}
}

func (s *Subscription) HasSeen(item Item) bool {
	s.makeSeenMap()
	_, seen := s.SeenItems[item.GUID]
	return seen
}

func (sub *Subscription) KeywordsSlice() []string {
	return strings.Split(sub.Keywords, ",")
}

func (s *Subscription) makeSeenMap() {
	if s.SeenItems == nil {
		s.SeenItems = make(map[string]bool)
		for _, i := range strings.Split(s.Seen, ",") {
			s.SeenItems[i] = true
		}
	}
}

type Subscriptions interface {
	Add(*Subscription) error
	Update(*Subscription) error
	ByFeedId(id uint64) ([]*Subscription, error)
}

type InMemSubs struct {
	items map[uint64]*Subscription
}

func NewInMemSubs() *InMemSubs {
	return &InMemSubs{items: make(map[uint64]*Subscription)}
}

func (m *InMemSubs) Add(s *Subscription) error {
	if s.Id == 0 {
		s.Id = rand.Uint64()
	}
	m.items[s.Id] = s
	return nil
}

func (m *InMemSubs) Update(s *Subscription) error {
	m.items[s.Id] = s
	return nil
}

func (m *InMemSubs) ByFeedId(id uint64) ([]*Subscription, error) {
	var found []*Subscription

	for _, s := range m.items {
		if s.FeedId == id {
			found = append(found, s)
		}
	}

	return found, nil
}
