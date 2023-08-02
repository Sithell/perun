package database

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Image   string
	Command string
}

type Provider struct {
	gorm.Model
	Host   string
	Status string
}

type Run struct {
	gorm.Model
	JobID      uint
	ProviderID uint
	Job        Job
	Provider   Provider
	Status     string
}
