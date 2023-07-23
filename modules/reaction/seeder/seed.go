package seeder

import (
	"context"
	"fmt"
	"newsfeed/ent"
)

func Seed(db *ent.Client) {

	reacts := make([]*ent.ReactCreate, 0)
	for i := 0; i < 50; i++ {
		react := db.React.Create().AddPostIDs(1).SetReactType("love")

		reacts = append(reacts, react)
	}
	err := db.React.CreateBulk(reacts...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
