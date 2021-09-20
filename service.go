package main

import "cchampou.me/logger"

type Service struct {
	Logger *logger.Logger
}

func LoadService(logfunc *logger.Logger) *Service {

	return &Service{
		Logger: logfunc,
	}
}

func (ser *Service) Go() {
	ser.Logger("hello")
}
