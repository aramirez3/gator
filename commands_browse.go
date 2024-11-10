package main

import (
	"context"
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/aramirez3/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var readLimit int64
	if len(cmd.arguments) == 1 {
		readLimit, _ = strconv.ParseInt(cmd.arguments[0], 10, 64)
	} else {
		readLimit = 2
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(readLimit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error retrieving posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found for user")
		return nil
	}

	for _, post := range posts {
		post := escapePost(&post)
		fmt.Printf("# %s\n", post.Title)
		fmt.Printf("     - %s\n", post.Description)
		fmt.Printf("     - Published %v\n", post.PublishedAt.Format(time.DateOnly))
		fmt.Printf("     - Url %v\n", post.Url)
		fmt.Println("----------------------")
	}

	return nil
}

func escapePost(post *database.GetPostsForUserRow) *database.GetPostsForUserRow {
	title := html.EscapeString(post.Title)
	description := html.UnescapeString(post.Description)
	post.Title = title
	post.Description = description
	return post
}
