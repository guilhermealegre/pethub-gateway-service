package writer

import "bitbucket.org/asadventure/be-infrastructure-lib/domain"

type Fallback struct {
	reader domain.FallbackReader
	writer domain.FallbackWriter
}

type WriterHandler func(message []byte) error
