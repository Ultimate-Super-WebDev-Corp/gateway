package file

import (
	"strings"

	"github.com/pkg/errors"
)

func makeBucketName(bucket, partition string) string {
	return strings.Join([]string{bucket, partition}, ".")
}

func makeFileUUID(uuid, partition string) string {
	return strings.Join([]string{uuid, partition}, ".")
}

func getPartition(str string) (string, error) {
	strs := strings.Split(str, ".")
	if len(strs) != 2 {
		return "", errors.Errorf("cannot get partition from %s", str)
	}
	return strs[1], nil
}
