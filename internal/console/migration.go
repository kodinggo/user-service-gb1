package console

import (
	"database/sql"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"

	"github.com/kodinggo/user-service-gb1/internal/config"
	"github.com/kodinggo/user-service-gb1/internal/helper"
)

var (
	direction string
	step      int = 1
)

func init() {
	rootCmd.AddCommand(migrationCMD)

	migrationCMD.Flags().StringVarP(&direction, "direction", "d", "up", "Migration direction")
	migrationCMD.Flags().IntVarP(&step, "step", "s", 1, "Number of migration steps")
}

var migrationCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run:   migrateDB,
}

func migrateDB(cmd *cobra.Command, args []string) {
	config.LoadWithViper()

	connDB, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Panicf("failed to connect server db, error: %s", err.Error())
	}

	defer connDB.Close()

	migrations := &migrate.FileMigrationSource{Dir: "./db/migrations"}

	var n int

	if direction == "down" {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Up, step)
	}
	if err != nil {
		log.Panicf("failed to run migration, error: %s", err.Error())
	}

	log.Printf("successfully applied %d migration(s)", n)
}
