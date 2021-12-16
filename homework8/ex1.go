package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	port       = flag.Int("port", 80, "Порт")
	db_url     = flag.String("db_url", "", "url1")
	jaeger_url = flag.String("jaeger_url", "", "url2")
	sentry_url = flag.String("sentry_url", "", "url3")
)

func main() {
	flag.Parse()
	fmt.Println(*port)
	fmt.Println(*db_url)

	//проверка на валидность
	if strings.HasPrefix(*jaeger_url, "http://") || strings.HasPrefix(*jaeger_url, "https://") {
		fmt.Println(*jaeger_url)
	}

	fmt.Println(*sentry_url)
}
