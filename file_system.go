package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type FileSystem struct {}

func (fs *FileSystem) createDb() {
	if _, err := os.Create(filePath); err != nil {
		log.Fatalf("failed on create file tasks.json: %s", err)
	}

	if err := fs.storeData(); err != nil {
		log.Fatal(err)
	}

	if err := fs.loadData(); err != nil {
		log.Fatal(err)
	}
}

func (fs *FileSystem) existDb() bool {
	_, err := os.Stat(filePath)

	return err == nil
}

func (fs *FileSystem) initProgram() {
	if ok := fs.existDb(); !ok {
		fs.createDb()
		return
	}
	

	if err := fs.loadData(); err != nil {
		log.Fatal(err)
	}
}

func (fs *FileSystem) storeData() error {
	data, err := json.Marshal(tasksdb)
	if err != nil {
		return fmt.Errorf("failed on marshal tasks: %s", err)
	}
	err = os.WriteFile(filePath, data, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed on write file: %s", err)
	}

	return nil
}

func (fs *FileSystem) loadData() error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed on read file: %s", err)
	}

	err = json.Unmarshal(data, &tasksdb)
	if err != nil {
		return fmt.Errorf("failed on unmarshal: %s", err)
	}

	return nil
}