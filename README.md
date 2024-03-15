# INSTALLATION
- Clone this repository
```bash
git clone https://github.com/syauqqii/translator
```
- Direct into translator directory
```bash
cd translator
```
- Rename env file
```bash
mv .env.example .env
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
- list language code [click here](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes)
# NOTE
Im using ```go version go1.21.5```
