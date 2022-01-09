# Django x PostgreSQL x React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
app
  - Dockerfile
  - requirements.txt
  - .gitignore
front
  - Dockerfile
docker-compose.yml
```
`requirements.txt`を編集し、プロジェクトに必要なPythonパッケージを追記する。

## 2. Djangoプロジェクトの作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm app django-admin startproject config .
```
`config`ディレクトリと`manage.py`が作成されていることを確認する。

## 3. データベースの設定

`app/config/setting.py`の項目を下記のように書き換える。
```
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql_psycopg2',
        'NAME': 'new_db',
        'USER': 'admin',
        'PASSWORD': 'admin',
        'HOST': 'db',
        'PORT': 5432,
    }
}
LANGUAGE_CODE = 'ja'
TIME_ZONE = 'Asia/Tokyo'
```

下記コマンドを実行し、データベースの変更を反映させる。
```
docker-compose run --rm app python manage.py makemigrations
```
```
docker-compose run --rm app python manage.py migrate
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
docker-compose up -d
```
Webブラウザを起動して以下にアクセスし、http://localhost:8000 でDjango、http://localhost:3000 でReactが起動していることを確認する。

## 6. Djangoアプリケーションの作成

新規ターミナルを開いて下記コマンドを実行し、アプリケーションを作成する。
```
docker-compose exec app python manage.py startapp <アプリ名>
```

## 7. その他

## 参考資料

https://qiita.com/greenteabiscuit/items/c40ba038703c9f33499b<br>
https://qiita.com/Kobayashi2019/items/66c03b29d8f5effd2341<br>
https://qiita.com/shiranon/items/b3efd3ed7ce473c6ad83
