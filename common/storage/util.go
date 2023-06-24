package storage

import (
	"fmt"
	"github.com/rohanchauhan02/clean/common/util"
	"path/filepath"
)

func sanitizeFileNameForUpload(filename string) string{
	dirPath := filepath.Dir(filename)
	basePath := filepath.Base(filename)
	basePath = util.SanitizeString(basePath)
	return fmt.Sprintf("%s/%s", dirPath, basePath)
}

func stringpointer(s string) *string {
	return &s
}

func int64pointer(i int64) *int64 {
	return &i
}
