package repository

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/xid"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
	"strings"
)

type staticFile struct {
	session *session.Session
}

func NewStaticFile(sess *session.Session) repository.StaticFile {
	return &staticFile{
		sess,
	}
}

func (repo *staticFile) StoreImage(img *entity.Image) (*entity.Image, error) {
	uploader := s3manager.NewUploader(repo.session)

	bs, err := decode(img.Base64)
	if err != nil {
		return nil, err
	}

	fileName := generateFileName(img.ContentType)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("emukone-dev"),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(bs),
	})
	if err != nil {
		return nil, err
	}
	newImg := &entity.Image{
		URL: result.Location,
	}

	return newImg, nil
}

func decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func generateFileName(contentType string) string {
	return fmt.Sprintf("%s.%s", xid.New().String(), retrieveExtension(contentType))
}

func retrieveExtension(str string) string {
	exs := strings.Split(str, "/")
	if len(exs) != 2 {
		return ""
	}
	return exs[1]
}
