package product

import (
	"context"
	"io/ioutil"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

type gcVisionRoundRobin struct {
	clietns []*vision.ImageAnnotatorClient
	current int
}

func newGcVisionRoundRobin(gcVisionPathToKeys string) (*gcVisionRoundRobin, error) {
	files, err := ioutil.ReadDir(gcVisionPathToKeys)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(files) == 0 {
		return nil, errors.New("must be at least one key in dir")
	}

	rb := &gcVisionRoundRobin{
		clietns: make([]*vision.ImageAnnotatorClient, 0, len(gcVisionPathToKeys)),
	}

	for _, f := range files {
		cli, err := vision.NewImageAnnotatorClient(context.Background(), option.WithCredentialsFile(gcVisionPathToKeys+"/"+f.Name()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		rb.clietns = append(rb.clietns, cli)
	}
	return rb, nil
}

func (g *gcVisionRoundRobin) next() *vision.ImageAnnotatorClient {
	n := g.clietns[g.current]

	g.current++
	if g.current >= len(g.clietns) {
		g.current = 0
	}
	return n
}

func (g *gcVisionRoundRobin) DetectTexts(ctx context.Context, img *pb.Image) ([]*pb.EntityAnnotation, error) {
	return g.next().DetectTexts(ctx, img, nil, 0)
}
