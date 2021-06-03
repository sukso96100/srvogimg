![](./img/example.png)
# srvogimg

Simple Open graph card image server

## Usage

- Configure `IMG_CACHE_PATH` environment variable, if you want to cache image on specific directory
- Run `go run .` to run test server
- Send `GET` request to `/render` with following params
    - `text`: Text to display.
    - `imgurl`: Web URL of the image to display
    - `startcolor`: Gradient start color in **hex** without **#** (e.g. `E95420`)
    - `endcolor`: Gradient start color in **hex** without **#** (e.g. `E95420`)


Example:
```
http://localhost:8080/render?imgurl=https://path.to.img/img.png&text=HelloWorld&startcolor=E95420&endcolor=E95420
```

## Credits
- [Default icon - from icons8.com](https://icons8.com/icon/65355/document)
- [Font - Spoqa Han Sans Neo + Emoji One]
  - [Spoqa Han Sans Neo](https://spoqa.github.io/spoqa-han-sans/)
  - [Adobe Emoji One](https://github.com/adobe-fonts/emojione-color)