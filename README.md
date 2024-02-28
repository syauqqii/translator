# INSTALLATION
- Clone this repository
```bash
git clone https://github.com/syauqqii/translator
```
- Direct into translator directory
```bash
cd translator
```
- Inititalization folder name
```go
go mod init translator
```
- Download package
```go
go mod tidy
```
- Running application
```go
go run main.go
```
- Access application with postman or something else
```bash
URL request path:
localhost:6666/translator

Format JSON request:
{
  "text": "..some text..",
  "target_language": "en",
  "source_language": "id"
}
```
# NOTE
Im using ```go version go1.21.5```
