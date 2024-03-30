APP=home

templ generate
GOOS=linux GOARCH=386 go build -o $APP
ssh linode "sudo service $APP stop"
scp $APP linode:/home/rustydoggobytes/apps/
ssh linode "sudo service $APP start"

