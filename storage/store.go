package storage

import (
	"cloud.google.com/go/bigquery"
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc/shared"
	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"io"
)

type Store struct {
	bqClient      *bigquery.Client
	writer        spanstore.Writer
	reader        spanstore.Reader
	archiveWriter spanstore.Writer
	archiveReader spanstore.Reader
}

func NewStore(logger hclog.Logger, cfg Configuration) (*Store, error) {
	return &Store{}, nil
}

func (s *Store) Close() error {
	return s.bqClient.Close()
}

func (s *Store) SpanReader() spanstore.Reader {
	return s.reader
}

func (s *Store) SpanWriter() spanstore.Writer {
	return s.writer
}

func (s *Store) DependencyReader() dependencystore.Reader {
	//TODO implement me
	panic("implement me")
}

func (s *Store) ArchiveSpanReader() spanstore.Reader {
	return s.archiveReader
}

func (s *Store) ArchiveSpanWriter() spanstore.Writer {
	return s.archiveWriter
}

var (
	_ shared.ArchiveStoragePlugin = (*Store)(nil)
	_ shared.StoragePlugin        = (*Store)(nil)
	_ io.Closer                   = (*Store)(nil)
)
