package main

import (
	"outback/stack/pipline"
	"outback/stack/spiders"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@192.168.199.104:5432/postgres"), &gorm.Config{})
	if err != nil {
		log.Error().Err(err)
		return
	}
	pip := pipline.NewCreate(db)

	sp := spiders.NewStarkSpider(pip)
	// sp := spiders.NewNameCode(pip)
	//
	// sp.ListSh()
	sp.Start()
}
