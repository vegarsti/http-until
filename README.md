# Using `until` to wait for an http server to start

The `until` bash command can be used to run a command until it succeeds, i.e. exits with exit code 0.

`go main.go` will start an HTTP server which exits after the first successful request.

`make` starts this Go program in the background, and waits `until` the server returns an OK response.

Try it by running `make`!
