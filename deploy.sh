APP=home

templ generate
GOOS=linux GOARCH=386 go build -o $APP
ssh linode "service $APP stop"
scp $APP linode:/root/apps/
ssh linode "service $APP start"

