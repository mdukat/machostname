machostname: machostname.go
	go fmt machostname.go
	go build machostname.go

install: machostname
	install -m 755 machostname /usr/bin/machostname
	install -m 644 machostname.conf.example /etc/machostname.cfg
	install -m 644 machostname.service /etc/systemd/system/machostname.service
