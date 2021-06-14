# Django x PostgreSQL x React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
api
  - Dockerfile
  - requirements.txt
  - .gitignore
front
  - Dockerfile
docker-compose.yml
```

## 2. Djangoプロジェクトの作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm api django-admin startproject config .
```
`config`ディレクトリと`manage.py`が作成されていることを確認する。

## 3. データベースの設定

`api/config/setting.py`内の`DATABASES`を下記のように書き換える。
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

## 4. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm front npx create-react-app react_app --template typescript
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。

## 5. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up
```
Webブラウザを起動して以下にアクセスし、http://localhost:8000 でDjango、http://localhost:3000 でReactが起動していることを確認する。

## 6. アプリケーションの作成

新規ターミナルを開いて下記コマンドを実行し、アプリケーションを作成する。
```
docker-compose exec api python3 manage.py startapp <アプリ名>
```

## 7. その他

## 参考資料

https://qiita.com/greenteabiscuit/items/c40ba038703c9f33499b<br>
https://qiita.com/Kobayashi2019/items/66c03b29d8f5effd2341
