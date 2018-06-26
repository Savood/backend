package image

/*

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

func resizeImage(origImg io.ReadCloser, width, height uint) error {
	var (
		original_image image.Image
		new_image      image.Image
		err            error
	)

	defer origImg.Close()

	if original_image, _, err = image.Decode(origImg); nil != err {
		log.Printf("image decode error! %v", err)

		return err

	}

	new_image = resize.Resize(width, height, original_image, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, new_image, nil); nil != err {
		log.Printf("image encode error! %v", err)
		return err
	}
	content = buf.Bytes()

	r.Header.Set("Content-Type", "image/jpeg")
}

func getImage(filename string) error {

	db := database.GetDatabase()

	gridFSfile, err := db.GridFS("fs").Open(filename)

	if err != nil {
		return err
	}

	if len(width_str) > 0 {

		if width, err = strconv.ParseUint(width_str, 10, 32); nil != err {
			status = http.StatusBadRequest
			return
		}

		is_width_set = true
	}

	height_str := r.URL.Query().Get("h")

	var (
		height        uint64
		is_height_set = false
	)
	if len(height_str) > 0 {

		if height, err = strconv.ParseUint(height_str, 10, 32); nil != err {
			status = http.StatusBadRequest
			return
		}

		is_height_set = true
	}

	logger.Info("width and height is [%d, %d], status [%t, %t]", width, height, is_width_set, is_height_set)

	if is_width_set || is_height_set {

	}

}

*/