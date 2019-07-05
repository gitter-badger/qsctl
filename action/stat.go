package action

import (
	"github.com/c2h5oh/datasize"

	"github.com/yunify/qsctl/helper"
	"github.com/yunify/qsctl/utils"
)

// Stat will handle all stat actions.
func Stat(remote string) (err error) {
	objectKey, err := ParseQsPath(remote, true)
	if err != nil {
		return
	}

	return StatRemoteObject(objectKey)
}

// StatRemoteObject will stat a remote object.
func StatRemoteObject(objectKey string) (err error) {
	om, err := helper.HeadObject(objectKey)
	if err != nil {
		return
	}

	content := []string{
		"Key: " + om.Key,
		"Size: " + datasize.ByteSize(om.ContentLength).String(),
		"Type: " + om.ContentType,
		"Modify: " + om.LastModified.String(),
		"StorageClass: " + om.StorageClass,
	}

	if om.ETag != "" {
		content = append(content, "MD5: "+om.ETag)
	}

	println(utils.AlignPrintWithColon(content...))
	return
}