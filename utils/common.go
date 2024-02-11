package utils

import (
	"fmt"
	"healthcare-capt-america/pkg/configs"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetConnection(config *configs.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.DbUser,
		config.Postgres.DbPass,
		config.Postgres.DbName,
	)
}

func ParseTime(t string) (time.Duration, error) {
	var mins, hours int
	var err error
	parts := strings.SplitN(t, ":", 2)
	switch len(parts) {
	case 1:
		mins, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
	case 2:
		hours, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}

		mins, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("invalid time: %s", t)
	}
	if mins > 59 || mins < 0 || hours > 23 || hours < 0 {
		return 0, fmt.Errorf("invalid time: %s", t)
	}
	return time.Duration(hours)*time.Hour + time.Duration(mins)*time.Minute, nil
}

func TimeParseString(t time.Duration) string {
	hours := int(t.Hours())
	minutes := int(t.Minutes()) % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func FileTypeValidation(fh *multipart.FileHeader) (bool, error) {
	extensions := map[string]string{
		"image/jpeg":      "jpg",
		"image/png":       "png",
		"application/pdf": "pdf",
	}
	file, err := fh.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}
	fileType := http.DetectContentType(buffer)
	if _, ok := extensions[fileType]; !ok {
		return false, nil
	}
	return true, nil
}

func GetFileType(fh *multipart.FileHeader) (string, error) {
	extensions := map[string]string{
		"image/jpeg":      "jpg",
		"image/png":       "png",
		"application/pdf": "pdf",
	}

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	fileType := http.DetectContentType(buffer)
	if res, ok := extensions[fileType]; ok {
		return fmt.Sprintf(".%s", res), nil
	}
	return "", nil
}
