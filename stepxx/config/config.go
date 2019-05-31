package config

import (
	"encoding/json"
	"os"
	"time"
)

type config struct {
	Player    string        `json:"player"`
	Ghost     string        `json:"ghost"`
	Wall      string        `json:"wall"`
	Dot       string        `json:"dot"`
	Pill      string        `json:"pill"`
	Death     string        `json:"death"`
	Space     string        `json:"space"`
	Chaser    string        `json:"chaser"`
	UseEmoji  bool          `json:"use_emoji"`
	FrameRate time.Duration `json:"frame_rate"`
	Lives     int           `json:"lives"`
}

var cfg config

func Player() string {
	return cfg.Player
}

func Ghost() string {
	return cfg.Ghost
}

func Wall() string {
	return cfg.Wall
}

func Dot() string {
	return cfg.Dot
}

func Pill() string {
	return cfg.Pill
}

func Death() string {
	return cfg.Death
}

func Space() string {
	return cfg.Space
}

func Chaser() string {
	return cfg.Chaser
}

func UseEmoji() bool {
	return cfg.UseEmoji
}

func FrameRate() time.Duration {
	if cfg.FrameRate <= 0 {
		return 5
	}
	return cfg.FrameRate
}

func Lives() int {
	if cfg.Lives <= 0 {
		return 3
	}
	return cfg.Lives
}

// Load a configuration file
func Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
