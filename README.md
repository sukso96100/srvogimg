# srvogimg

Simple Open graph card image server

## Usage

Run `go run .` to run test server

Send `GET` request to `/render` with following params

- `text`: Text to display.
- `imgurl`: Web URL of the image to display

Example:
```
http://localhost:8080/render?imgurl=https://path.to.img/img.png&text=HelloWorld
```

## Credits
- [Default icon - from icons8.com](https://icons8.com/icon/65355/document)
- [Font - Noto Sans CJK KR](https://www.google.com/get/noto/#sans-kore)