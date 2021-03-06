// +build OMIT

import (
	"context"
	"sync"

	"github.com/go-kit/kit/endpoint"
)

func ImagePreprocessingPipeline(numOfDatasetLoader, numOfImageDecoders, numOfImagePreprocessors) {
	pipeline := endpoint.Chain(
		// CPU-bound + IO-bound task
		loadImageDataset(numOfDatasetLoaders),

		// CPU-bound task
		decodeImages(numOfImageDecoders),

		// Image Preprocessing
		preprocessingImages(numOfImagePreprocessors),
	//...
	)(progressUpdaterEndpoint)
}

func readImages(numOfImageReaders int) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			in := request.(chan ImageEntry)
			// channel for loaded images
			out := make(chan ImageEntry, 1000)
			go func() {
				wg := sync.WaitGroup{}
				for i := 0; i < numOfImageReaders; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						ImageReader(in, out)
					}()
				}
				wg.Wait()
				close(out)
			}()
			return next(ctx, out)
		}
	}
}
