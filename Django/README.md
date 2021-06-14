# Django

各ディレクトリ内の`README.md`を参照。

## Djangoコマンド

```
# プロジェクトの作成
django-admin startproject <プロジェクト名> .

# アプリケーションの作成
python manage.py startapp <アプリ名>

# マイグレーションファイルを作成
python manage.py makemigrations

# マイグレーションファイルをデータベースに適用
python manage.py migrate

# 開発用のWebサーバーを起動
python manage.py runserver 0.0.0.0:8000

# 管理者ユーザーの作成
python manage.py createsuperuser
```

## 参考資料

https://freeheroblog.com/django-command/<br>
https://32imuf.com/django/manage/<br>
https://qiita.com/STAR77/items/ffcf622e640dfd2f7a21
