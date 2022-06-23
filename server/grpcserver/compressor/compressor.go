package compressor

import (
	"io"
	"sync"

	"github.com/klauspost/compress/snappy"
	"go.uber.org/fx"
	"google.golang.org/grpc/encoding"
)

const (
	Name = "snappy"
)

var (
	compressorPool   sync.Pool
	decompressorPool sync.Pool
)

var Module = fx.Options(
	fx.Provide(newCompressor),
	fx.Invoke(RegisterCompressor),
)

type Compressor struct{}

func RegisterCompressor(c *Compressor) {
	if c == nil {
		c = new(Compressor)
	}

	encoding.RegisterCompressor(c)
}

func newCompressor() *Compressor {
	c := new(Compressor)
	return c
}

func (c *Compressor) Name() string {
	return Name
}

func (c *Compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	wr, inPool := compressorPool.Get().(*writeCloser)
	if !inPool {
		return newWriteCloser(w), nil
	}

	wr.Reset(w)

	return wr, nil
}

func (c *Compressor) Decompress(r io.Reader) (io.Reader, error) {
	dr, inPool := decompressorPool.Get().(*reader)
	if !inPool {
		return newReader(r), nil
	}

	dr.Reset(r)

	return dr, nil
}

type (
	reader      struct{ *snappy.Reader }
	writeCloser struct{ *snappy.Writer }
)

func newWriteCloser(w io.Writer) *writeCloser {
	return &writeCloser{
		Writer: snappy.NewBufferedWriter(w),
	}
}

func (w *writeCloser) Close() error {
	defer func() {
		compressorPool.Put(w)
	}()

	return w.Writer.Close()
}

func newReader(r io.Reader) *reader {
	return &reader{
		Reader: snappy.NewReader(r),
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if err == io.EOF {
		decompressorPool.Put(r)
	}

	return n, err
}
