build-binary:
		GOOS=darwin GOARCH=amd64 go build -o ./binary/darwin_amd64/todo
		GOOS=linux GOARCH=amd64 go build -o ./binary/linux_amd64/todo
		GOOS=windows GOARCH=amd64 go build -o ./binary/windows_amd64/todo.exe