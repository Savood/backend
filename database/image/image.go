package image

import (
	"io"
	"git.dhbw.chd.cx/savood/backend/database"
	"image"
	"log"
	"github.com/nfnt/resize"
	"bytes"
	"image/jpeg"
)

const fs = "fs"

// UploadImage creates a GridFS file and writes to it
func UploadImage(filename string, inputFile io.ReadCloser) error {

	db := database.GetDatabase()

	gridFSfile, err := db.GridFS(fs).Create(filename)

	if err != nil {
		return err
	}

	defer gridFSfile.Close()
	_, err = io.Copy(gridFSfile, inputFile)

	if err != nil {
		return err
	}

	err = inputFile.Close()

	if err != nil {
		return err
	}

	return nil
}

type closingBuffer struct {
	*bytes.Buffer
}

func (cb *closingBuffer) Close() (err error) {
	//we don't actually have to do anything here, since the buffer is  just some data in memory
	//and the error is initialized to no-error
	return
}

// ResizeImage resizes the image
func ResizeImage(origImg io.ReadCloser, width, height uint) (io.ReadCloser, error) {
	var (
		originalImage image.Image
		newImage      image.Image
		err           error
	)

	defer origImg.Close()

	if originalImage, _, err = image.Decode(origImg); nil != err {
		log.Printf("image decode error! %v", err)

		return nil, err

	}

	newImage = resize.Resize(width, height, originalImage, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, newImage, nil); nil != err {
		log.Printf("image encode error! %v", err)
		return nil, err
	}

	return &closingBuffer{buf}, nil
}

// GetImage gets Image from gridFS
func GetImage(filename string) (io.ReadCloser, error) {

	db := database.GetDatabase()

	gridFSfile, err := db.GridFS(fs).Open(filename)

	if err != nil {
		return nil, err
	}

	return gridFSfile, nil

}

// DeleteImage deletes an Image from gridFS
func DeleteImage(filename string) error {

	db := database.GetDatabase()

	err := db.GridFS(fs).Remove(filename)

	return err

}
