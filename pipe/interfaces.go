package pipe

// Interface definitions

type PipeReader interface {
	Read(b []byte) (n int, e error)
}

type PipeWriter interface {
	Write(b []byte) (n int, e error)
}

type PipeCloser interface {
	Close() error
}

type PipeReadWriter interface {
	PipeReader
	PipeWriter
}

type PipeReadCloser interface {
	PipeReader
	PipeCloser
}

type PipeWriteCloser interface {
	PipeWriter
	PipeCloser
}

type PipeReadWriteCloser interface {
	PipeReader
	PipeWriter
	PipeCloser
}
