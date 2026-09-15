package objectstore

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type Store struct {
	client *minio.Client
	bucket string
}

func New(client *minio.Client, bucket string) *Store {
	return &Store{client: client, bucket: bucket}
}

func (s *Store) EnsureBucket(ctx context.Context) error {
	ok, err := s.client.BucketExists(ctx, s.bucket)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return s.client.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{})
}

func (s *Store) PresignPut(ctx context.Context, key string, expiry time.Duration) (*url.URL, error) {
	return s.client.PresignedPutObject(ctx, s.bucket, key, expiry)
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
