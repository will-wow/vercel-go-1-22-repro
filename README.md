# vercel-go-1-22-repro

Minimal repro of Go 1.22 not supported in Vercel's Go runtime.

## Repro

- Go to /
- See "Hello from go"
- Go to /one22

Expected:

- See 1,2,3,4,5

Actual:

- Get an error because ranging over integers is not supported in Go 1.21
