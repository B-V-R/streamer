package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Destination interface {
	Write()
	ErrorCh() chan error
}

type File struct {
	Path    string
	Data    []byte
	errorCh chan error
	log *log.Logger
}

type Stdout struct {
	Data    []byte
	errorCh chan error
	log *log.Logger
}

func (file *File) ErrorCh() chan error {
	return file.errorCh
}

func (stdout *Stdout) ErrorCh() chan error {
	return stdout.errorCh
}

func New(log *log.Logger, data []byte) (Destination, Destination) {
	return &File{"destination_file.txt", data, make(chan error, 1), log}, &Stdout{data, make(chan error, 1), log}
}

func (file *File) Write() {
	err := ioutil.WriteFile(file.Path, file.Data, 777)
	if err != nil {
		file.log.Error(err.Error())
	}

	file.errorCh <- err
	close(file.errorCh)
}

func (stdout *Stdout) Write() {
	person, err := UnmarshalPerson(stdout.Data)
	defer close(stdout.errorCh)
	if err != nil {
		stdout.log.Error(err.Error())
		stdout.errorCh <- err
		return
	}
	_, err = fmt.Println(person)
	stdout.errorCh <- err
}
