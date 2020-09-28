package product

import (
	"bytes"
	"context"
	"io"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
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
func (p Product) newTextRecognitionByUUIDFuture(ctx context.Context, uuid string) textRecognitionByUUIDFuture {
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
		r, err := p.textRecognitionByUUID(ctx, uuid)
		if err != nil {
			future.err <- err
		} else {
			future.resp <- r
		}
	}(ctx, uuid, f)

	return f
}

type recognizedText struct {
	text string
}

func (p Product) textRecognitionByUUID(ctx context.Context, uuid string) (*recognizedText, error) {
	buff := bytes.NewBuffer([]byte{})
	streamGetFile, err := p.fileCli.GetFile(ctx, &file.FileUUID{
		UUID: uuid,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var meta *file.Chunk_Meta
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
			meta = ch
		case *file.Chunk_Chunk:
			_, _ = buff.Write(ch.Chunk)
		default:
			return nil, errors.New("unknown type of chunk")
		}
	}
	if meta != nil && meta.Meta != nil && len(meta.Meta.RecognizedText) > 0 {
		return &recognizedText{
			text: meta.Meta.RecognizedText,
		}, nil
	}
	image, err := vision.NewImageFromReader(buff)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	annotations, err := p.visionRR.DetectTexts(ctx, image)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := recognizedText{}
	if len(annotations) > 0 && len(annotations[0].Description) > 0 {
		resp.text = strings.Replace(annotations[0].Description, "\n", " ", -1)
		_, err := p.fileCli.UpdateMetadata(ctx, &file.UpdateFileMetadata{
			UUID: uuid,
			Meta: &file.FileMetadata{
				RecognizedText: resp.text,
			},
		})
		if err != nil {
			ctxzap.Extract(ctx).Error(err.Error())
		}
	}

	return &resp, nil
}
