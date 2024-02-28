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
localhost:6666/translator

Format JSON
{
  "text": "..some text..",
  "target_language": "en",
  "source_language": "id"
}
```
