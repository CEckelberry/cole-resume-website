package http

import "context"

// requestCtx is an alias to context.Context so middleware.go can avoid
// repeatedly importing the same package and to keep the function signatures
// readable. There's no behavior difference.
type requestCtx = context.Context

func contextWithValue(ctx requestCtx, key, val any) requestCtx {
	return context.WithValue(ctx, key, val)
}
