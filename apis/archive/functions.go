package archive

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"io"
	"mime/multipart"
)

func createZip(files []*multipart.FileHeader, buf *bytes.Buffer) (err error) {
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()
	for _, fileHeader := range files {
		var file multipart.File
		file, err = fileHeader.Open()
		if err != nil {
			return
		}
		defer file.Close()
		var w io.Writer
		w, err = zipWriter.Create(fileHeader.Filename)
		if err != nil {
			return
		}
		_, err = io.Copy(w, file)
		if err != nil {
			return
		}
	}
	return
}

func createTarGz(files []*multipart.FileHeader, buf *bytes.Buffer) (err error) {
	gzipWriter := gzip.NewWriter(buf)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	for _, fileHeader := range files {
		var file multipart.File
		file, err = fileHeader.Open()
		if err != nil {
			return
		}
		defer file.Close()
		hdr := &tar.Header{
			Name: fileHeader.Filename,
			Mode: 0600,
			Size: fileHeader.Size,
		}
		err = tarWriter.WriteHeader(hdr)
		if err != nil {
			return
		}
		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return
		}
	}
	return
}
