[ Tue Oct 01 23:07:03 ~/Dev/gocode/src ] revel package mychat
~
~ revel! http://robfig.github.com/revel
~
2013/10/01 23:07:15 revel.go:253: Loaded module static
2013/10/01 23:07:15 reflect.go:288: Could not find import: mychat/app/routes
2013/10/01 23:07:15 build.go:72: Exec: [/usr/local/go/bin/go build -tags  -o /home/mike/Dev/gocode/bin/mychat mychat/app/tmp]
Abort: Failed to parse template /home/mike/Dev/gocode/src/github.com/robfig/revel/cmd/package_run.sh.template: open /home/mike/Dev/gocode/src/github.com/robfig/revel/cmd/package_run.sh.template: no such file or directory


https://www.digitalocean.com/community/articles/how-to-install-go-and-revel-on-an-ubuntu-13-04-x64-vps
go get -u github.com/robfig/revel/revel


[ Wed Oct 02 21:18:36 /usr/local/go/bin ] sudo su
mike-HP-Pavilion-dv8000-EE944AV bin # export GOPATH=/home/mike/Dev/gocode
mike-HP-Pavilion-dv8000-EE944AV bin # /usr/local/go/bin/go get github.com/robfig/revel/revel
mike-HP-Pavilion-dv8000-EE944AV bin # 

