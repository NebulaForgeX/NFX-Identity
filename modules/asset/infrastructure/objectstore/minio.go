package objectstore

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type Store struct {
	client    *minio.Client
	presigner *minio.Client
	bucket    string
}

func New(client *minio.Client, bucket string) *Store {
	return NewWithPresigner(client, nil, bucket)
}

func NewWithPresigner(client, presigner *minio.Client, bucket string) *Store {
	if presigner == nil {
		presigner = client
	}
	return &Store{client: client, presigner: presigner, bucket: bucket}
}

func (s *Store) EnsureBucket(ctx context.Context) error {
	ok, err := s.client.BucketExists(ctx, s.bucket)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	err = s.client.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{})
	if err == nil {
		return nil
	}
	resp := minio.ToErrorResponse(err)
	if resp.Code == "BucketAlreadyOwnedByYou" || resp.Code == "BucketAlreadyExists" {
		return nil
	}
	return err
}

func (s *Store) PresignPut(ctx context.Context, key string, expiry time.Duration) (*url.URL, error) {
	return s.presigner.PresignedPutObject(ctx, s.bucket, key, expiry)
}

func (s *Store) Stat(ctx context.Context, key string) (minio.ObjectInfo, error) {
	return s.client.StatObject(ctx, s.bucket, key, minio.StatObjectOptions{})
}

func (s *Store) Remove(ctx context.Context, key string) error {
	return s.client.RemoveObject(ctx, s.bucket, key, minio.RemoveObjectOptions{})
}

func (s *Store) Get(ctx context.Context, key string) (*minio.Object, error) {
	return s.client.GetObject(ctx, s.bucket, key, minio.GetObjectOptions{})
}

func (s *Store) Open(ctx context.Context, key string) (io.ReadCloser, error) {
	return s.Get(ctx, key)
}
