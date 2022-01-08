package metanit

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type SSS struct {
	Port        int    `json:"port"`
	DbURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func Openjs() SSS {
	file, err := os.Open("file1.txt")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	sss := SSS{}

	err = json.NewDecoder(file).Decode(&sss)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	if !strings.HasPrefix(sss.JaegerURL, "http://") || strings.HasPrefix(sss.JaegerURL, "https://") {
		sss.JaegerURL = "wrong url"
	}

	if !strings.HasPrefix(sss.SentryURL, "http://") || strings.HasPrefix(sss.SentryURL, "https://") {
		sss.SentryURL = "wrong url"
	}
	return sss
}
