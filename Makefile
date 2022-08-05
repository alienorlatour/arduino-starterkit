flash-00:
	tinygo flash -target=arduino 00-blinky/main.go

flash-02:
	tinygo flash -target=arduino 02-spaceship/main.go

flash-street-full:
	tinygo flash -target=arduino streetlight/full/main.go

flash-street-route:
	tinygo flash -target=arduino -scheduler=tasks streetlight/route/main.go
	#tinygo flash -target=arduino streetlight/route/main.go
