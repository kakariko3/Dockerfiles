# Django Setup

## 1. Dockerイメージの作成
```
$ docker build -t django .
```

## 2. Dockerコンテナの起動
```
$ docker run -itd -p 8000:8000 -v <作業ディレクトリ>:/work --name <コンテナ名> django
```

## 3. Djangoサーバー構築
### プロジェクトの作成
---
```
$ cd <作業ディレクトリ>
```
```
$ docker exec <コンテナ名> django-admin startproject <プロジェクト名> .
```

### アプリの作成
---
```
$ docker exec <コンテナ名> python3 manage.py startapp <アプリ名>
```

### 開発用サーバを起動
---
```
$ cd ~/
```
```
$ docker exec -d <コンテナ名> python3 manage.py runserver 0.0.0.0:8000
```

## 4. その他
### 参考資料
---
https://qiita.com/homines22/items/2730d26e932554b6fb58
