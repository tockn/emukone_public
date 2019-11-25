# emukone

## CI/CD
CircleCI使ってる

## 環境構築

- Server Side: Go (gin)
- Frontend: Nuxt.js
- DB: MySQL8.0 (docker)

DBはdev用とtest用でコンテナ分けて２つ動いてる

まずはDB環境つくる
```
$ make docker_up
```

マイグレーションする
```
$ make migrate_up
```

サーバーだけ動かす
```
$ make server
```

フロントだけ動かす
```
$ make front
```

MySQL(dev)に潜る
```
$ make mysql
```
passは`password`

なんと、swagger生成もできます。
```
$ make swagger
```
これで`docs/swagger/`にswaggerファイルができるので、適当なツールで閲覧してください

### DBテーブル設計

- リリースするまではガンガン変えていきます。

![db](https://user-images.githubusercontent.com/29294540/54074940-37ec9280-42dc-11e9-8d28-ccd1fbddd6fd.png)
