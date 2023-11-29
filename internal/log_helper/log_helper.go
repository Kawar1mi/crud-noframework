package loghelper

import "github.com/rs/zerolog/log"

func FatalIfError(err error) {
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func InfoMsg(msg string) {
	log.Info().Msg(msg)
}

func AddError(err error) {
	log.Error().Msg(err.Error())
}
