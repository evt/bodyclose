# bodyclose
This repository includes 2 almost identical scripts - `cmd/noclose/main.go` and `cmd/withclose/main.go`.
The only difference if `defer resp.Body.Close()` missing in `noclose` version.

Running them both for a while (at least 5-10 minutes) will show how memory leaks if you don't do `defer resp.Body.Close()`.
