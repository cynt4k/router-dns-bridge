package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile   string
	configFolder string
	rootCmd      = &cobra.Command{
		Use:          "router-dns-bridge",
		SilenceUsage: true,
		Run:          func(c *cobra.Command, args []string) {},
	}
	cfg = config.GetConfig()
)

func Execute() error {
	return rootCmd.Execute()
}

func init() { // nolint:gochecknoinits
	cobra.OnInitialize(initConfig)

	rootCmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		err := c.Usage()

		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	})

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&configFolder, "configfolder", "f", "", "config folder")
}

func initConfig() {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		var fileDir string
		var configDir string
		if configFolder != "" {
			fileDir = configFolder
		} else {
			fileDir = filepath.Dir(os.Args[0])
		}
		fileDir, err := filepath.Abs(fileDir)
		if err != nil {
			log.Fatal(err)
		}
		if configFolder != "" {
			configDir = fileDir
		} else {
			configDir = filepath.Join(fileDir, "config")
		}
		viper.AddConfigPath(configDir)

		err = readDefaultConfig(configDir)

		if err != nil {
			log.Fatal(err)
		}

		switch env := strings.ToLower(os.Getenv("ENV")); env {
		case "dev":
			viper.SetConfigFile(filepath.Join(configDir, "dev.yaml"))
		case "prd":
			viper.SetConfigFile(filepath.Join(configDir, "prd.yaml"))
		case "":
			if err := viper.Unmarshal(cfg); err != nil {
				log.Fatalf("error while unmarshal config %s", err)
			}
			return
		default:
			log.Fatalf("unknown env variable %s", env)
		}
	}

	if err := viper.MergeInConfig(); err != nil {
		log.Printf("error while reading config %v", err)
	}

	if os.Getenv("ENV") == "dev" {
		viper.Set("DevMode", true)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("error while unmarshal config %s", err)
	}
	cfg = config.GetConfig()
}
