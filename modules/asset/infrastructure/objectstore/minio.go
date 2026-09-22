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
	lan    *minio.Client
	public *minio.Client
	bucket string
}

func New(client *minio.Client, bucket string) *Store {
	return NewWithPresigners(client, nil, nil, bucket)
}

func NewWithPresigner(client, presigner *minio.Client, bucket string) *Store {
	return NewWithPresigners(client, presigner, presigner, bucket)
}

func NewWithPresigners(client, lan, public *minio.Client, bucket string) *Store {
	if public == nil {
		public = client
	}
	if lan == nil {
		lan = public
	}
	return &Store{client: client, lan: lan, public: public, bucket: bucket}
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
	return s.pick(ctx).PresignedPutObject(ctx, s.bucket, key, expiry)
}

func (s *Store) PresignGet(ctx context.Context, key string, expiry time.Duration) (*url.URL, error) {
	return s.pick(ctx).PresignedGetObject(ctx, s.bucket, key, expiry, nil)
}

func (s *Store) pick(ctx context.Context) *minio.Client {
	if lanFromCtx(ctx) {
		return s.lan
	}
	return s.public
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
