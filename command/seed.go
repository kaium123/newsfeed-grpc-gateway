package command

import (

	//	"newsfeed/ent"

	//"fmt"

	"newsfeed/db"
	attachmentSeeder "newsfeed/modules/attachment/seeder"
	userSeeder "newsfeed/modules/user/seeder"
	postSeeder "newsfeed/modules/post/seeder"
	reactSeeder "newsfeed/modules/reaction/seeder"
	commentSeeder "newsfeed/modules/comment/seeder"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run sync",
	Run: func(cmd *cobra.Command, args []string) {
		db:=db.NewEntDb()
		attachmentSeeder.Seed(db)
		userSeeder.Seed(db)
		postSeeder.Seed(db)
		reactSeeder.Seed(db)
		commentSeeder.Seed(db)
	},
}
