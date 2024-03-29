# Setup Django x PostgreSQL x React

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend
│   ├── .gitignore
│   └── requirements.txt
├── docker
│   ├── frontend
│   │   └── Dockerfile
│   └── backend
│       └── Dockerfile
├── .gitignore
└── docker-compose.yml
```
`requirements.txt`を編集し、プロジェクトに必要なPythonパッケージを追記する。

## 2. Djangoプロジェクトの作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm backend django-admin startproject config .
```
`config`ディレクトリと`manage.py`が作成されていることを確認する。

## 3. データベースの設定

`backend/config/setting.py`の項目を下記のように書き換える。
```
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql_psycopg2',
        'NAME': 'postgres',
        'USER': 'postgres',
        'PASSWORD': 'postgres',
        'HOST': 'db',
        'PORT': 5432,
    }
}

LANGUAGE_CODE = 'ja'

TIME_ZONE = 'Asia/Tokyo'
```

下記コマンドを実行し、データベースの変更を反映させる。
```
docker-compose run --rm backend python manage.py makemigrations
```
```
docker-compose run --rm backend python manage.py migrate
```

## 4. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm frontend npx create-react-app . --template typescript
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
docker-compose exec backend python manage.py startapp <アプリ名>
```

## 7. その他

## VSCodeの自動補完機能を有効化

https://qiita.com/suzuki_sh/items/6bc15446965df20b6c5a#vscode%E3%81%A7docker%E5%86%85%E3%81%AEpython%E3%82%92%E5%88%A9%E7%94%A8%E3%81%99%E3%82%8B<br>
上記のURLを参考に、コンテナ(backend)にリモートアクセスする新規ウィンドウを開き、VSCode拡張機能(ms-python.python)をインストールする。<br>
また、VSCodeで使用するPythonをコンテナ内のものに変更する。(インタープリターを選択)

## 参考資料

https://qiita.com/greenteabiscuit/items/c40ba038703c9f33499b<br>
https://qiita.com/Kobayashi2019/items/66c03b29d8f5effd2341<br>
https://qiita.com/shiranon/items/b3efd3ed7ce473c6ad83<br>
https://qiita.com/mk-tool/items/1c7e4929055bb3b7aeda
