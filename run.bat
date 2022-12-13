taskkill /f /im main.exe
cd "cmd\server"
go run main.go --sqlite true
pause