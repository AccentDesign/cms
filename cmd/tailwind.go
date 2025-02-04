package cmd

import (
	"context"
	"echo.go.dev/pkg/storage/db/dbx"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
	"os"
)

var cmdTailwind = &cobra.Command{
	Use:   "tailwind",
	Short: "Generate tailwind safelist file using html content in the database",
	Run: func(cmd *cobra.Command, args []string) {
		runTailwind()
	},
}

func runTailwind() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.Database.URL().String())
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	queries := dbx.New(conn)

	classes, err := queries.GetCSSClasses(ctx)
	if err != nil {
		fmt.Printf("Error getting classes: %v\n", err)
		os.Exit(1)
	}

	if err = genTailwindSafelist(classes); err != nil {
		fmt.Println("Failed to update config:", err)
		os.Exit(1)
	}
}

func genTailwindSafelist(classes []string) error {
	file, err := os.Create("tailwind.safelist.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	for _, c := range classes {
		_, err = file.WriteString(c + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return nil
}
