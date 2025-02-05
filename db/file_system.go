package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/silas-ss/task-tracker-cli/model"
)


type FileSystem struct {}

const filePath string = "./tasks.json"

func (fs *FileSystem) createDb() {
	if _, err := os.Create(filePath); err != nil {
		log.Fatalf("failed on create file tasks.json: %s", err)
	}
}

func (fs *FileSystem) existDb() bool {
	_, err := os.Stat(filePath)

	return err == nil
}

func (fs *FileSystem) runMigrate() {
	if err := fs.StoreData([]model.Task{}); err != nil {
		log.Fatalf("failed on run migrate data. error: %s", err)
	}
}

func (fs *FileSystem) InitProgram() {
	if ok := fs.existDb(); !ok {
		fs.createDb()
		fs.runMigrate()
		return
	}
}

func (fs *FileSystem) StoreData(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed on marshal tasks: %s", err)
	}
	err = os.WriteFile(filePath, data, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed on write file: %s", err)
	}

	return nil
}

func (fs *FileSystem) LoadData() ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, fmt.Errorf("failed on read file: %s", err)
	}

	return data, nil
}