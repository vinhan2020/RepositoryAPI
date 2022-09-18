package repository

import "gorm.io/gorm"

type postgreRepo struct {
	db *gorm.DB
}

func NewPostgreRepo(db *gorm.DB) *postgreRepo {
	return &postgreRepo{db: db}
}
