# Dockerfiles

各ディレクトリ内の`README.md`を参照。

## dockerコマンド
```
# コンテナ一覧の表示
docker ps -a

# イメージ一覧
docker images -a

# 停止中のコンテナを削除
docker container prune

# <none>タグのイメージを一括削除
docker image prune

# コンテナ、ボリューム、ネットワーク、イメージを一括削除
docker system prune -a
```

## docker-composeコマンド
```
# コンテナの作成と起動
docker-compose up

# バックグラウンドで起動
docker-compose up -d

# 起動
docker-compose start

# 停止
docker-compose stop

# 停止＆削除
docker-compose down

# コンテナのログを表示
docker-compose logs -f

# 稼働中のコンテナに入る
docker-compose exec <サービス名> ash

# コンテナ内のコマンドを実行
docker-compose run <サービス名> <コマンド>
```
