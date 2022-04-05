# go-dynamic-lib

# Example build command

> go build -gcflags='all=-N -l' -buildmode=plugin -o ./plugins/my-printer/my-printer.so ./plugins/my-printer/

We need the flag `-gcflags='all=-N -l'` to allow debugging (details [here](https://github.com/go-delve/delve/issues/865#issuecomment-480766102))

When we run the main application we also need to add `-gcflags='all=-N -l'`:

> go run -gcflags='all=-N -l' main.go

**Note:** The plugins have to be build from inside the main application. Otherwise a "plugin versions do not match" error will be given when trying to open the .so file. More [here](https://github.com/golang/go/issues/31354)