build_linux:
	go env -w GOOS=linux
	go build -o modbus main.go
	scp modbus root@192.168.0.20:/root/modbus_util
	ssh root@192.168.0.20 chmod +x /root/modbus_util/modbus
	go env -w GOOS=windows

run_linux: build_linux
	ssh root@192.168.0.20 /root/modbus_util/modbus
