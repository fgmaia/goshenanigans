package asyncloop

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type IndexedOpFunc[C any] func(ctx context.Context, i int) (C, error)

type OpFunc[C any] func(ctx context.Context) (C, error)

func RangedAsyncLoop[C any](ctx context.Context, n, size int, opFunc IndexedOpFunc[*C]) (retValues []*C, errs []error) {
	ranges := Ranges(n, size)
	errs = make([]error, n)
	retValues = make([]*C, n)

	_, _ = AsyncLoop(ctx, len(ranges), func(ctx context.Context, i int) (*C, error) {
		r := ranges[i]
		for j := r.Start; j < r.End; j++ {
			retValues[j], errs[j] = opFunc(ctx, j)
		}
		return nil, nil
	})
	return retValues, errs
}

func AsyncLoop[C any](ctx context.Context, len int, opFunc IndexedOpFunc[C]) (retValues []C, errs []error) {
	errs = make([]error, len)
	retValues = make([]C, len)

	var g errgroup.Group
	for i := 0; i < len; i++ {
		i := i
		g.Go(func() error {
			retValues[i], errs[i] = opFunc(ctx, i)
			return nil
		})
	}

	_ = g.Wait()
	return retValues, errs
}

func AsyncOps[C any](ctx context.Context, opsFunc ...OpFunc[C]) ([]C, []error) {
	return AsyncLoop(ctx, len(opsFunc), func(ctx context.Context, i int) (C, error) {
		return opsFunc[i](ctx)
	})
}

func AnyErr(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func GetLastErrAndCount(errs []error) (int, error) {
	var lastErr error
	c := 0
	for _, err := range errs {
		if err != nil {
			lastErr = err
			c++
		}
	}
	return c, lastErr
}
