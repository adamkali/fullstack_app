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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
    "github.com/swaggo/echo-swagger"
	_ "github.com/adamkali/fullstack_app/docs"
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
    e.Use(middleware.Logger())
    e.GET("/_health", func(ctx echo.Context) error {
        return ctx.JSON(http.StatusOK, map[string]interface{} { "ok": true })
    })
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

    ctx := context.Background()
    db, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
    if err != nil {
        e.Logger.Fatal(err.Error())
    }

    params := cnts.ControllerParams{
        CTX: &ctx,
        DB: db,
    }
    cnts.AttatchControllers(e,
        cnts.BuildUserController(&params),
    )



    e.Logger.Fatal(e.Start(":5052"))
}
