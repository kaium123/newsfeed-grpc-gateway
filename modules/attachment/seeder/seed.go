package seeder

import (
	"context"
	"fmt"

	"newsfeed/ent"
	"newsfeed/ent/attachment"

	fakeData "github.com/brianvoe/gofakeit/v6"
)

func Seed(db *ent.Client) {
	attachments := make([]*ent.AttachmentCreate, 0)
	ctx := context.Background()
	for i := 0; i < 50; i++ {
		url := fakeData.ImageURL(1024, 768)
		isPresent, err := db.Attachment.Query().Where(attachment.Path(url)).Exist(ctx)
		if err != nil || isPresent {
			continue
		}

		attachment := db.Attachment.Create().
			SetPath(url)
		attachments = append(attachments, attachment)
	}
	err := db.Attachment.CreateBulk(attachments...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
