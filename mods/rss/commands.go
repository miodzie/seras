package rss

import (
	"fmt"
	"strings"

	"github.com/miodzie/seras"
)

func (mod *RssMod) showFeeds(msg seras.Message) {
	feeds, err := mod.feeds.All()
	if err != nil {
		fmt.Println(err)
		mod.actions.Send(seras.Message{Content: "Failed to fetch feeds.", Channel: msg.Channel})
		return
	}
	if len(feeds) == 0 {
		mod.actions.Send(seras.Message{Content: "No feeds available.", Channel: msg.Channel})
		return
	}

	var reply = seras.Message{Channel: msg.Channel}
	var parsed []string
	for _, feed := range feeds {
		parsed = append(parsed, fmt.Sprintf("%s: %s", feed.Name, feed.Url))
	}
	reply.Content = strings.Join(parsed, "\n")
	reply.Content += fmt.Sprintf("\n\nTo subscribe to a feed, use %ssubscribe {name} {keywords}, keywords being comma separated (spaces are ok, e.g. \"spy x family, comedy\")", seras.Token())

	mod.actions.Send(reply)
}

// !add_feed {name} {url}
func (mod *RssMod) addFeed(msg seras.Message) {
	// TODO: validate.
	feed := &Feed{
		Name: msg.Arguments[1],
		Url:  msg.Arguments[2],
	}
	fmt.Println(feed.Name, feed.Url)
	err := mod.feeds.Save(feed)
	if err != nil {
		fmt.Println(err)
	}
}

// !subscribe {feed name} {keywords, comma separated}
func (mod *RssMod) subscribe(msg seras.Message) {
	feed, err := mod.feeds.GetByName(msg.Arguments[1])
	if err != nil {
		mod.actions.Send(seras.Message{Content: "Unknown feed.", Channel: msg.Channel})
		fmt.Println(err)
		return
	}
	// TODO: parse, test
	keywords := strings.Join(msg.Arguments[2:], " ")
	fmt.Println(keywords)
	sub := &Subscription{
		FeedId:   feed.Id,
		Channel:  msg.Channel,
		Keywords: keywords,
		User:     msg.AuthorId,
	}
	err = mod.subs.Save(sub)
	if err != nil {
		fmt.Println(err)
		mod.actions.Send(seras.Message{
			Content: "Failed to save feed, likely one already exists for this channel and feed.",
			Channel: msg.Channel,
		})
	}
}
