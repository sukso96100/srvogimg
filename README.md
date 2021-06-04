![](./img/example.png)
# srvogimg

Simple Open graph card image server

## Usage

### Environment variables
- `IMG_CACHE_PATH`: If you want to cache image on specific directory
- `APP_HOST`: Hostname (or domain) for the deployed app. Required to see RESTful API doc properly

### Run server

```bash
go get
go run .
```

Visit `/swagger/index.html`(If running on local environment, `localhost:8080/swagger/index.html`) to see available image render api


Example:
```
http://localhost:8080/render?imgurl=https://path.to.img/img.png&text=HelloWorld&startcolor=E95420&endcolor=E95420
```

## Credits
- [Default icon - from icons8.com](https://icons8.com/icon/65355/document)
- Default font - Merged following 2 fonts
  - [Spoqa Han Sans Neo](https://spoqa.github.io/spoqa-han-sans/)
  - [Adobe EmojiOne Color](https://github.com/adobe-fonts/emojione-color)