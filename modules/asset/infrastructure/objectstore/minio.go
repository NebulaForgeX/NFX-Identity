package objectstore

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/cors"
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
	if !ok {
		if err := s.client.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}
	return s.client.SetBucketCors(ctx, s.bucket, cors.NewConfig([]cors.Rule{{
		AllowedOrigin: []string{"*"},
		AllowedMethod: []string{"GET", "PUT", "HEAD", "POST", "DELETE"},
		AllowedHeader: []string{"*"},
		ExposeHeader:  []string{"ETag", "Accept-Ranges", "Content-Range", "Content-Length", "Content-Type"},
		MaxAgeSeconds: 3600,
	}}))
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
