package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/notifier"
	"os"
	"time"
)

func initLog() {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	initLog()
	log.Info("GMRN Start.")
	configPtr := flag.String("c", "", "Config file path.")

	flag.Parse()

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
	nf := notifier.InitGitLabNotifier(config.Url, config.Token, config.Interval.Duration)
	log.Infof("Start Notifier site: %s", nf.Url)

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
	Url           string
	Token         string
	Interval      duration
	NotifyCommand string
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
