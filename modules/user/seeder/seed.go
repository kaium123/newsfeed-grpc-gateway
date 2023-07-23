package seeder

import (
	"context"
	"fmt"

	"newsfeed/ent"
)

func Seed(db *ent.Client) {

	users := make([]*ent.UserCreate, 0)
	for i := 0; i < 50; i++ {
		s:=fmt.Sprintf("first_name_%v", i)
		user := db.User.Create().
			SetName(s)
			 
			
			
		users = append(users, user)
	}
	err := db.User.CreateBulk(users...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
