/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"net/http"
	"os"

	cnts "github.com/adamkali/fullstack_app/controllers"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
    Long: `Serve the application on the 0.0.0.0:5173 host and port.

The default environment can be used with the -e environment flag. 
This is can be any environment presumably in the config directory.

config 
  development.yaml
  staging.yaml
  production.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
        serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
    e := echo.New()
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

    db, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
    if err != nil {
        e.Logger.Fatal(err.Error())
    }
    params := cnts.ControllerParams{
        DB: db,
    }

    cnts.AttatchControllers(e,
        cnts.BuildAuthController(&params),
        cnts.BuildUserController(&params),
    )

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

    e.GET("/_health", func(ctx echo.Context) error {
        return ctx.JSON(http.StatusOK, map[string]interface{} { "ok": true })
    })

	e.Logger.Fatal(e.Start(":1323"))
    
}
