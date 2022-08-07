flash-00:
	tinygo flash -target=arduino 00-blinky/main.go

flash-02:
	tinygo flash -target=arduino 02-spaceship/main.go

flash-08:
	tinygo flash -target=arduino -scheduler=tasks 08-hourglass/main.go