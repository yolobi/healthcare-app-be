package services

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
)

func newClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx, option.WithCredentialsFile("./files/cloud/keys.json"))
}

func UploadToBucket(ctx context.Context, file multipart.File, directory string) (*string, error) {
	bucket := os.Getenv("BUCKET")
	var err error

	storageClient, err = newClient(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	sw := storageClient.Bucket(bucket).Object(directory).NewWriter(ctx)

	_, err = io.Copy(sw, file)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	if err := sw.Close(); err != nil {
		return nil, apperror.NewServerError(err)
	}
	u, err := url.Parse("/" + bucket + sw.Attrs().Name)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	path := u.EscapedPath()
	return &path, nil
}

func UploadCertificate(ctx context.Context, file *multipart.FileHeader, name string) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	directory := fmt.Sprintf("/certificate/%s.pdf", name)
	return UploadToBucket(ctx, f, directory)
}

func UploadIcon(ctx context.Context, file *multipart.FileHeader, name string) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	log.Println("nama file : ", name)
	directory := fmt.Sprintf("/public/category_icon/%s", name)
	return UploadToBucket(ctx, f, directory)
}

func UploadProductImage(ctx context.Context, file *multipart.FileHeader, name string) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	directory := fmt.Sprintf("/product/image/%s", name)
	return UploadToBucket(ctx, f, directory)
}

func UploadUserImage(ctx context.Context, file *multipart.FileHeader, name string) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	directory := fmt.Sprintf("/user/image/%s.png", name)
	return UploadToBucket(ctx, f, directory)
}

func UploadPaymentFile(ctx context.Context, file *multipart.FileHeader, name string) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	directory := fmt.Sprintf("/payment/%s", name)
	return UploadToBucket(ctx, f, directory)
}
