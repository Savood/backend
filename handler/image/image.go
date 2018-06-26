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

func uploadImage(filename string, inputFile io.ReadCloser) error {

	db := database.GetDatabase()

	gridFSfile, err := db.GridFS("fs").Create(filename)

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

func resizeImage(origImg io.ReadCloser, width, height uint) (io.ReadCloser, error) {
	var (
		original_image image.Image
		new_image      image.Image
		err            error
	)

	defer origImg.Close()

	if original_image, _, err = image.Decode(origImg); nil != err {
		log.Printf("image decode error! %v", err)

		return nil, err

	}

	new_image = resize.Resize(width, height, original_image, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, new_image, nil); nil != err {
		log.Printf("image encode error! %v", err)
		return nil, err
	}

	return &closingBuffer{buf}, nil
}

func getImage(filename string) (io.ReadCloser, error) {

	db := database.GetDatabase()

	gridFSfile, err := db.GridFS("fs").Open(filename)

	if err != nil {
		return nil, err
	}

	return gridFSfile, nil

}
