package seeder

import (
	"context"
	"fmt"
	"newsfeed/ent"
)

func Seed(db *ent.Client) {

	comments := make([]*ent.CommentCreate, 0)
	for i := 0; i < 50; i++ {
		comment := db.Comment.Create().AddPostIDs(1).SetContent("sdda").AddReactIDs(1,2,3)
		comments = append(comments, comment)
	}
	err := db.Comment.CreateBulk(comments...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
