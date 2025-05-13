package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

// RSSFeed represents the top-level RSS feed structure.
type RSSFeed struct {
	Channel RSSChannel `xml:"channel"`
}

// RSSChannel represents the channel section of the RSS feed.
type RSSChannel struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	Items       []RSSItem `xml:"item"`
}

// RSSItem represents a single item in the RSS feed.
type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var rss RSSFeed
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, fmt.Errorf("unmarshaling XML: %w", err)
	}

	// Unescape HTML entities in channel
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	// Unescape HTML entities in each item
	for i := range rss.Channel.Items {
		rss.Channel.Items[i].Title = html.UnescapeString(rss.Channel.Items[i].Title)
		rss.Channel.Items[i].Description = html.UnescapeString(rss.Channel.Items[i].Description)
	}

	return &rss, nil
}
