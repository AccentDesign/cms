package cmd

import (
	"bytes"
	"context"
	"echo.go.dev/pkg/storage/db/dbx"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var (
	tailwindConfigTemplate = "tailwind.config.js.tmpl"
	tailwindConfig         = "tailwind.config.js"
)

var cmdTailwind = &cobra.Command{
	Use:   "tailwind",
	Short: "Generate tailwind config file with a safelist using html content in the database",
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

	if err = genTailwindConfig(classes); err != nil {
		fmt.Println("Failed to update config:", err)
		os.Exit(1)
	}
}

func genTailwindConfig(classes []string) error {
	file, err := os.ReadFile(tailwindConfigTemplate)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(string(file))
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, classes)
	if err != nil {
		return err
	}

	err = os.WriteFile(tailwindConfig, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Updated Tailwind config written to", tailwindConfig)
	return nil
}
