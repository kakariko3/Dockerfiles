# Django x PostgreSQL Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml
- requirements.txt
- .gitignore

## 2. Djangoプロジェクトの作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm web django-admin startproject config .
```
`config`ディレクトリと`manage.py`が作成されていることを確認する。

## 3. データベースの設定

`config/setting.py`内の`DATABASES`を下記のように書き換える。
```
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'postgres',
        'USER': 'postgres',
        'PASSWORD': 'password',
        'HOST': 'db',
        'PORT': 5432,
    }
}
```

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up
```
Webブラウザで http://localhost:8000 へアクセスし、Djangoが起動していることを確認する。

## 5. アプリケーションの作成

新規ターミナルを開いて下記コマンドを実行し、アプリケーションを作成する。
```
docker-compose exec web python manage.py startapp <アプリ名>
```

## 6. その他

## 参考資料

https://docs.docker.jp/compose/django.html<br>
https://datacoach.me/data/tips/docker-django/<br>
https://www.inglow.jp/techblog/docker-django/
