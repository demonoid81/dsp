package cmd

import (
	"fmt"
	"github.com/demonoid81/dsp/app"
	"github.com/demonoid81/dsp/pkg/server"
	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	envFile    string
	configFile string

	DSPCmd = &cobra.Command{
		Use:   "dsp",
		Short: "dsp",
		Long:  "dsp server",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := Env.NewMongoClient()
			if err != nil {
				return err
			}

			server.HTTPServer(Env)
			return nil
		},
	}

	Env *app.Env
)

func Execute() error {
	if err := DSPCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func init() {
	Env = &app.Env{
		Cfg: &app.Config{
		},
	}

	cobra.OnInitialize(func() {
		// Load values from .env into the system
		godotenv.Load(envFile)

		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			viper.AddConfigPath("/etc")
			viper.AddConfigPath(home)
			viper.AddConfigPath("./")
			viper.SetConfigName("dsp")
		}

		// Setup environment variables
		viper.SetEnvPrefix("DSP")
		viper.AutomaticEnv()

		// Read configuration file.
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				// The configuration file does not found in search path.
				// Skip reading the configuration file.
			default:
				// Failed to read the configuration file.
				// Exit the program.
				_, _ = fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	})

	DSPCmd.Flags().StringVar(&envFile, "env-file", app.DefaultEnvFile, "path to the environment variables file. if omitted, .env in the current directory will be set.")
	//sengineCmd.Flags().StringVar(&configFile, "config-file", defaultConfigFile, "path to the configuration file. if omitted, sengine.yaml will be searched for in the /etc directory, then the home directory, and will be set if found.")

	// APP
	DSPCmd.Flags().StringVar(&Env.Cfg.Crypto, "crypto", app.DefaultCrypto, "Crypto key used to encrypt")
	DSPCmd.Flags().StringVar(&Env.Cfg.ClickURL, "click-url", app.DefaultClickURL, "Click URL")
	DSPCmd.Flags().StringVar(&Env.Cfg.MediaURL, "media-url", app.DefaultMediaURL, "Media URL")
	DSPCmd.Flags().StringVar(&Env.Cfg.RedirectURL, "redirect-url", app.DefaultRedirectURL, "Redirect URL")
	DSPCmd.Flags().Float64Var(&Env.Cfg.Revshare, "revshare", app.DefaultRevshare, "Revshare cost")

	// PG
	//DSPCmd.Flags().StringVar(&App.Cfg.PG.Host, "postgres-host", app.DefaultPGHost, "PostgresSQL host")
	//DSPCmd.Flags().IntVar(&App.Cfg.PG.Port, "postgres-port", app.DefaultPGPort, "PostgresSQL port")
	//DSPCmd.Flags().StringVar(&App.Cfg.PG.User, "postgres-user", app.DefaultPGUser, "PostgresSQL user")
	//DSPCmd.Flags().StringVar(&App.Cfg.PG.Password, "postgres-password", app.DefaultPGPassword, "PostgresSQL password")
	//DSPCmd.Flags().StringVar(&App.Cfg.PG.Database, "postgres-database", app.DefaultPGDatabase, "PostgresSQL database")

	// Mongo DB
	DSPCmd.Flags().StringVar(&Env.Cfg.MongoURL, "mongo-url", app.DefaultMongoURL, "PostgresSQL password")
	DSPCmd.Flags().StringVar(&Env.Cfg.MongoDatabase, "mongo-db", app.DefaultMongoDatabase, "PostgresSQL database")

	// Redis database
	DSPCmd.Flags().StringVar(&Env.Cfg.RedisHost, "redis-host", app.DefaultRedisHost, "Redis database host")
	DSPCmd.Flags().IntVar(&Env.Cfg.RedisPort, "redis-port", app.DefaultRedisPort, "Redis database port")

	// CORS
	DSPCmd.Flags().StringSliceVar(&Env.Cfg.CorsAllowedMethods, "cors-allowed-methods", []string{}, "CORS allowed methods (e.g. GET,PUT,DELETE,POST)")
	DSPCmd.Flags().StringSliceVar(&Env.Cfg.CorsAllowedOrigins, "cors-allowed-origins", []string{}, "CORS allowed origins (e.g. http://localhost:8080,http://localhost:80)")
	DSPCmd.Flags().StringSliceVar(&Env.Cfg.CorsAllowedHeaders, "cors-allowed-headers", []string{}, "CORS allowed headers (e.g. content-type,x-some-key)")

	//
	//sengineCmd.Flags().StringVar(&logLevel, "log-level", defaultLogLevel, "log level")
	//sengineCmd.Flags().StringVar(&logFile, "log-file", defaultLogFile, "log file")
	//sengineCmd.Flags().IntVar(&logMaxSize, "log-max-size", defaultLogMaxSize, "max size of a log file in megabytes")
	//sengineCmd.Flags().IntVar(&logMaxBackups, "log-max-backups", defaultLogMaxBackups, "max backup count of log files")
	//sengineCmd.Flags().IntVar(&logMaxAge, "log-max-age", defaultLogMaxAge, "max age of a log file in days")
	//sengineCmd.Flags().BoolVar(&logCompress, "log-compress", defaultLogCompress, "compress a log file")

	DSPCmd.Flags().BoolP("version", "v", false, "show version")

	DSPCmd.Flags().SortFlags = false

	_ = viper.BindPFlag("crypto", DSPCmd.Flags().Lookup("crypto"))
	_ = viper.BindPFlag("click-url", DSPCmd.Flags().Lookup("click-url"))
	_ = viper.BindPFlag("media-url", DSPCmd.Flags().Lookup("media-url"))
	_ = viper.BindPFlag("redirect-url", DSPCmd.Flags().Lookup("redirect-url"))
	_ = viper.BindPFlag("revshare", DSPCmd.Flags().Lookup("revshare"))

	//_ = viper.BindPFlag("postgres-host", DSPCmd.Flags().Lookup("postgres-host"))
	//_ = viper.BindPFlag("postgres-port", DSPCmd.Flags().Lookup("redirect-url"))
	//_ = viper.BindPFlag("postgres-user", DSPCmd.Flags().Lookup("postgres-user"))
	//_ = viper.BindPFlag("postgres-password", DSPCmd.Flags().Lookup("postgres-password"))
	//_ = viper.BindPFlag("postgres-database", DSPCmd.Flags().Lookup("postgres-database"))

	_ = viper.BindPFlag("mongo-url", DSPCmd.Flags().Lookup("mongo-url"))
	_ = viper.BindPFlag("mongo-db", DSPCmd.Flags().Lookup("mongo-db"))

	_ = viper.BindPFlag("redis-host", DSPCmd.Flags().Lookup("redis-host"))
	_ = viper.BindPFlag("redis-port", DSPCmd.Flags().Lookup("redis-port"))

	_ = viper.BindPFlag("cors_allowed_methods", DSPCmd.Flags().Lookup("cors-allowed-methods"))
	_ = viper.BindPFlag("cors_allowed_origins", DSPCmd.Flags().Lookup("cors-allowed-origins"))
	_ = viper.BindPFlag("cors_allowed_headers", DSPCmd.Flags().Lookup("cors-allowed-headers"))

}
