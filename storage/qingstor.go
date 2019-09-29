package storage

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pengsrc/go-shared/convert"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yunify/qingstor-sdk-go/v3/config"
	sdkerrors "github.com/yunify/qingstor-sdk-go/v3/request/errors"
	"github.com/yunify/qingstor-sdk-go/v3/service"

	"github.com/yunify/qsctl/v2/constants"
	"github.com/yunify/qsctl/v2/pkg/fault"
	"github.com/yunify/qsctl/v2/pkg/types/storage"
)

// QingStorObjectStorage will implement ObjectStorage interface.
type QingStorObjectStorage struct {
	service *service.Service
	bucket  *service.Bucket
}

// NewQingStorObjectStorage will create a new qingstor object storage.
func NewQingStorObjectStorage() (q *QingStorObjectStorage, err error) {
	cfg, err := config.New(
		viper.GetString(constants.ConfigAccessKeyID),
		viper.GetString(constants.ConfigSecretAccessKey),
	)
	if err != nil {
		return nil, fault.NewStorageServiceInitFailed(err)
	}

	cfg.Host = viper.GetString(constants.ConfigHost)
	cfg.Port = viper.GetInt(constants.ConfigPort)
	cfg.Protocol = viper.GetString(constants.ConfigProtocol)
	cfg.ConnectionRetries = viper.GetInt(constants.ConfigConnectionRetries)
	cfg.LogLevel = viper.GetString(constants.ConfigLogLevel)

	q = &QingStorObjectStorage{}

	q.service, err = service.Init(cfg)
	if err != nil {
		log.Errorf("Init service failed [%v]", err)
		return nil, fault.NewStorageServiceInitFailed(err)
	}

	log.Debugf("Init service for access key <%s> succeed", cfg.AccessKeyID)
	return
}

// SetupBucket implements ObjectStorage.SetupBucket
func (q *QingStorObjectStorage) SetupBucket(name, zone string) (err error) {
	if zone == "" {
		zone = viper.GetString(constants.ConfigZone)
	}

	if zone != "" {
		q.bucket, err = q.service.Bucket(name, zone)
		if err != nil {
			return fault.NewStorageBucketInitFailed(err, name, zone)
		}
		return
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf("%s://%s.%s:%d",
		viper.GetString(constants.ConfigProtocol),
		name,
		viper.GetString(constants.ConfigHost),
		viper.GetInt(constants.ConfigPort))

	r, err := client.Head(url)
	if err != nil {
		return fault.NewStorageBucketInitFailed(err, name, "")
	}
	if r.StatusCode != http.StatusTemporaryRedirect {
		err = fmt.Errorf("head status is %d instead of %d", r.StatusCode, http.StatusTemporaryRedirect)
		return fault.NewStorageBucketInitFailed(err, name, "")
	}

	// Example URL: https://bucket.zone.qingstor.com
	zone = strings.Split(r.Header.Get("Location"), ".")[1]
	q.bucket, err = q.service.Bucket(name, zone)
	if err != nil {
		return fault.NewStorageBucketInitFailed(err, name, zone)
	}
	return
}

// HeadObject will head object.
func (q *QingStorObjectStorage) HeadObject(objectKey string) (om *storage.ObjectMeta, err error) {
	resp, err := q.bucket.HeadObject(objectKey, nil)
	if err != nil {
		var e sdkerrors.QingStorError
		// TODO: we should add a common status code check to return related errors.
		if errors.As(err, &e) {
			return nil, fault.NewStorageObjectNotFound(err, objectKey)
		}
		if errors.As(err, &e) {
			return nil, fault.NewStorageObjectNoPermission(err, objectKey)
		}
		return nil, fault.NewUnhandled(err)
	}

	om = &storage.ObjectMeta{
		Key:           objectKey,
		ContentLength: convert.Int64Value(resp.ContentLength),
		ContentType:   convert.StringValue(resp.ContentType),
		ETag:          convert.StringValue(resp.ETag),
		LastModified:  convert.TimeValue(resp.LastModified),
		StorageClass:  convert.StringValue(resp.XQSStorageClass),
	}
	return
}

// InitiateMultipartUpload will initiate a Multipart upload.
func (q *QingStorObjectStorage) InitiateMultipartUpload(objectKey string) (uploadID string, err error) {
	resp, err := q.bucket.InitiateMultipartUpload(objectKey, nil)
	if err != nil {
		return "", fault.NewUnhandled(err)
	}

	uploadID = *resp.UploadID
	return
}

// UploadMultipart will upload a Multipart.
func (q *QingStorObjectStorage) UploadMultipart(
	objectKey, uploadID string, size int64, partNumber int, md5sum []byte, r io.Reader,
) (err error) {
	_, err = q.bucket.UploadMultipart(objectKey, &service.UploadMultipartInput{
		Body:          r,
		ContentLength: convert.Int64(size),
		UploadID:      convert.String(uploadID),
		PartNumber:    convert.Int(partNumber),
		ContentMD5:    convert.String(hex.EncodeToString(md5sum[:])),
	})
	if err != nil {
		return fault.NewUnhandled(err)
	}
	return
}

// CompleteMultipartUpload will complete a Multipart upload.
func (q *QingStorObjectStorage) CompleteMultipartUpload(objectKey, uploadID string, totalNumber int) (err error) {
	parts := make([]*service.ObjectPartType, totalNumber)
	for i := 0; i < totalNumber; i++ {
		parts[i] = &service.ObjectPartType{
			PartNumber: convert.Int(i),
		}
	}

	_, err = q.bucket.CompleteMultipartUpload(
		objectKey, &service.CompleteMultipartUploadInput{
			UploadID:    convert.String(uploadID),
			ObjectParts: parts,
		})
	if err != nil {
		return fault.NewUnhandled(err)
	}
	return nil
}

// GetObject will get an object.
func (q *QingStorObjectStorage) GetObject(objectKey string) (r io.Reader, err error) {
	resp, err := q.bucket.GetObject(objectKey, nil)
	if err != nil {
		return nil, fault.NewUnhandled(err)
	}
	return resp.Body, nil
}

// PutBucket will make a bucket with specific name.
func (q *QingStorObjectStorage) PutBucket() error {
	// Request and create bucket
	_, err := q.bucket.Put()
	if err != nil {
		log.Errorf("Make bucket <%s> in zone <%s> failed [%v]",
			*q.bucket.Properties.BucketName, *q.bucket.Properties.Zone, err)
		return err
	}
	return nil
}

// DeleteBucket will delete a bucket.
func (q *QingStorObjectStorage) DeleteBucket() error {
	// Request and delete bucket
	_, err := q.bucket.Delete()
	if err != nil {
		return fault.NewUnhandled(err)
	}
	return nil
}

// DeleteObject will delete an object with specific key.
func (q *QingStorObjectStorage) DeleteObject(objectKey string) (err error) {
	if _, err = q.bucket.DeleteObject(objectKey); err != nil {
		return fault.NewUnhandled(err)
	}
	return nil
}

// ListBuckets will list all Buckets of the user.
func (q *QingStorObjectStorage) ListBuckets(zone string) (buckets []string, err error) {
	res, err := q.service.ListBuckets(&service.ListBucketsInput{Location: convert.String(zone)})
	if err != nil {
		return nil, fault.NewUnhandled(err)
	}
	log.Debugf("<%d> Buckets found.\n", *res.Count)
	for _, b := range res.Buckets {
		log.Debugf("Bucket <%s>, url <%s>, created <%s>, location <%s>\n",
			*b.Name, *b.URL, *b.Created, *b.Location)
		buckets = append(buckets, *b.Name)
	}
	return buckets, nil
}

// ListObjects will list all objects with specific prefix and delimiter from a bucket.
func (q *QingStorObjectStorage) ListObjects(prefix, delimiter string, marker *string) (oms []*storage.ObjectMeta, err error) {
	for {
		res, err := q.bucket.ListObjects(&service.ListObjectsInput{
			Delimiter: convert.String(delimiter),
			Prefix:    convert.String(prefix),
			Marker:    marker,
		})
		if err != nil {
			return nil, fault.NewUnhandled(err)
		}
		// Add directories into oms (if exists)
		for _, cpf := range res.CommonPrefixes {
			oms = append(oms, &storage.ObjectMeta{
				Key:         convert.StringValue(cpf),
				ContentType: constants.DirectoryContentType,
			})
		}
		// Add objects into oms
		for _, obj := range res.Keys {
			oms = append(oms, &storage.ObjectMeta{
				Key:           convert.StringValue(obj.Key),
				ContentLength: convert.Int64Value(obj.Size),
				ContentType:   convert.StringValue(obj.MimeType),
				ETag:          convert.StringValue(obj.Etag),
				LastModified:  time.Unix(int64(convert.IntValue(obj.Modified)), 0),
				StorageClass:  convert.StringValue(obj.StorageClass),
				Children:      nil,
			})
		}

		// recursively for next marker request
		if !convert.BoolValue(res.HasMore) {
			break
		}

		marker = res.NextMarker
	}
	return
}

// GetBucketACL will get acl from a bucket.
func (q *QingStorObjectStorage) GetBucketACL() (ar *storage.ACLResp, err error) {
	res, err := q.bucket.GetACL()
	if err != nil {
		return nil, fault.NewUnhandled(err)
	}
	ar = &storage.ACLResp{
		OwnerID:   convert.StringValue(res.Owner.ID),
		OwnerName: convert.StringValue(res.Owner.Name),
	}
	ar.ACLs = make([]*storage.ACLMeta, 0)
	for _, acl := range res.ACL {
		ar.ACLs = append(ar.ACLs, &storage.ACLMeta{
			GranteeType: convert.StringValue(acl.Grantee.Type),
			GranteeID:   convert.StringValue(acl.Grantee.ID),
			GranteeName: convert.StringValue(acl.Grantee.Name),
			Permission:  convert.StringValue(acl.Permission),
		})
	}
	return
}

// PutObject will put a object.
func (q *QingStorObjectStorage) PutObject(objectKey string, md5sum []byte, r io.Reader) (err error) {
	_, err = q.bucket.PutObject(objectKey, &service.PutObjectInput{
		ContentMD5: convert.String(hex.EncodeToString(md5sum[:])),
		Body:       r,
	})
	if err != nil {
		return fault.NewUnhandled(err)
	}
	return nil
}

// GetBucketZone will get base info from current bucket.
func (q *QingStorObjectStorage) GetBucketZone() (zone string) {
	return *q.bucket.Properties.Zone
}

// PresignObject will get pre-sign url for given object key and expire second.
func (q *QingStorObjectStorage) PresignObject(objectKey string, expire int) (url string, err error) {
	r, _, err := q.bucket.GetObjectRequest(objectKey, nil)
	if err != nil {
		return "", fault.NewUnhandled(err)
	}
	if err = r.Build(); err != nil {
		return "", fault.NewUnhandled(err)
	}
	if err = r.SignQuery(expire); err != nil {
		return "", fault.NewUnhandled(err)
	}
	return r.HTTPRequest.URL.String(), nil
}