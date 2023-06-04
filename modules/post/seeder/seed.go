package seeder

import (
	"context"
	"fmt"
	"newsfeed/ent"
)

func Seed(db *ent.Client) {

	posts := make([]*ent.PostCreate, 0)
	for i := 0; i < 50; i++ {
		post := db.Post.Create(). 
		AddAttachmentIDs(1,2,3). 
		SetContent("post").AddAuthorIDs(1,2,3)	 
			
			
		posts = append(posts, post)
	}
	err := db.Post.CreateBulk(posts...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
