APP=home

templ generate
GOOS=linux GOARCH=386 go build -o $APP
scp $APP linode:/home/rustydoggobytes/apps/

