package search

import (
	"bytes"
	"context"
	"io"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/pkg/errors"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

type textRecognitionByUUIDFuture struct {
	err  chan error
	resp chan *recognizedText
}

func (f textRecognitionByUUIDFuture) get() (*recognizedText, error) {
	select {
	case res := <-f.resp:
		return res, nil
	case err := <-f.err:
		return nil, err
	}
}
func (s Search) newTextRecognitionByUUIDFuture(ctx context.Context, uuid string) textRecognitionByUUIDFuture {
	f := textRecognitionByUUIDFuture{
		err:  make(chan error, 1),
		resp: make(chan *recognizedText, 1),
	}

	go func(ctx context.Context, uuid string, future textRecognitionByUUIDFuture) {
		defer func() {
			if p := recover(); p != nil {
				future.err <- errors.Errorf("recovering from panic %v", p)
			}
		}()
		r, err := s.textRecognitionByUUID(ctx, uuid)
		if err != nil {
			future.err <- err
		} else {
			future.resp <- r
		}
	}(ctx, uuid, f)

	return f
}

type recognizedText struct {
	text  string
	words []string
}

func (s Search) textRecognitionByUUID(ctx context.Context, uuid string) (*recognizedText, error) {
	buff := bytes.NewBuffer([]byte{})
	streamGetFile, err := s.fileCli.GetFile(ctx, &file.FileUUID{
		UUID: uuid,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for {
		msg, err := streamGetFile.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		switch ch := msg.OneOfChunk.(type) {
		case *file.Chunk_Meta:
			continue
		case *file.Chunk_Chunk:
			_, _ = buff.Write(ch.Chunk)
		default:
			return nil, errors.WithStack(err)
		}
	}

	image, err := vision.NewImageFromReader(buff)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	annotations, err := s.imageAnnotatorClient.DetectTexts(ctx, image, nil, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := recognizedText{
		words: make([]string, 0, len(annotations)),
	}

	for i, a := range annotations {
		if i == 0 {
			resp.text = strings.Replace(a.Description, "\n", " ", -1)
		} else {
			resp.words = append(resp.words, a.Description)
		}
	}
	return &resp, nil
}
