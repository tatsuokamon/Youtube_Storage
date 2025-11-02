package main

import (
	"iter"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (s *Storage)Default(file string) (error){
	_, err := os.Stat(file)
	if err != nil {
		return err
	}
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return err
	}

	err = s.init()
	if err != nil {
		return err
	}

	s.DB = db
	return nil
}

func (s *Storage)init() (error){
	err := s.DB.AutoMigrate(&Item{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage)Add(item Item) (error){
	tx := s.DB.FirstOrCreate(item)
	return tx.Error
}

func (s *Storage)GetAll()([]Item, error) {
	var result []Item
	tx := s.DB.Statement.Find(&result)
	return result, tx.Error
}
