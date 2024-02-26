# vercel-go-1-22-repro

Minimal repro of Go 1.22 not supported in Vercel's Go runtime.

## Repro

1. Deploy to Vercel
1. See `Warning: Unknown Go version in /vercel/path0/go.mod` in the build logs
1. Check that Go 1.22 was installed anyway:
	1. Go to /api
	1. See "Hello from go"
	1. Go to /api/one22
	1. See 1,2,3,4,5 (which only works on Go 1.22)
