package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	//"github.com/eternnoir/gmrn/apis"
	"github.com/eternnoir/gmrn/notifier"
	"os"
	"time"
)

func initLog(debug bool) {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	log.Info("GMRN Start.")
	configPtr := flag.String("c", "", "Config file path.")
	debugPtr := flag.Bool("d", false, "Debug flag")

	flag.Parse()

	initLog(*debugPtr)
	if len(*configPtr) < 1 {
		log.Error("Config file path must set.Use -h to get some help.")
	}

	gitlatsite, err := LoadConfig(*configPtr)
	if err != nil {
		log.Fatalf("Load config %s fail.", *configPtr)
		return
	}
	RunNotifier(gitlatsite)
}

func RunNotifier(config *GitLabConfig) {
	nf := notifier.InitGitLabNotifier(config.Url, config.Token, config.Projects, config.PollingInterval.Duration, config.NotifyInterval.Duration, config.NotifyCommand)
	log.Infof("Start Notifier site: %s", nf.Url)
	nf.Run()
}

func LoadConfig(path string) (*GitLabConfig, error) {
	log.Infof("Load config file from %s.", path)
	var config GitLabConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Error(err)
		return nil, err
	}

	log.Infof("Url: %s.", config.Url)
	return &config, nil
}

type GitLabConfig struct {
	Url             string
	Token           string
	PollingInterval duration
	NotifyInterval  duration
	Projects        []string
	NotifyCommand   string
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
