# executable-example

Simple example to show how to get Go executable path.

Using:
- os.Args[0]
- filepath.Abs(os.Args[0])
- runtime.Caller(0)
- github.com/kardianos/osext => osext.Executable

## Other links
http://stackoverflow.com/questions/12090170/go-find-the-path-to-the-executable/

http://stackoverflow.com/questions/18537257/golang-how-to-get-the-directory-of-the-currently-running-file/
