# executable-example

Simple example to show how to get executable path for Go program.

Using:

- `os.Args[0]`
- `filepath.Abs(os.Args[0])`
- `runtime.Caller(0)`
- github.com/kardianos/osext => `osext.Executable`

What you should test to compare behavior:

- running builded binary from different folder
- `go run`
- running from symlinked folder/file
- running in cron
- running this program from other program/script

## Other links

<http://stackoverflow.com/questions/12090170/go-find-the-path-to-the-executable/>

<http://stackoverflow.com/questions/18537257/golang-how-to-get-the-directory-of-the-currently-running-file/>
