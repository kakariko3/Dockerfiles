# Dockerfiles

各ディレクトリ内の`README.md`を参照。

## Docker コマンド

```
# コンテナ一覧の表示
docker ps -a

# イメージ一覧の表示
docker images -a

# 停止中のコンテナを削除
docker container prune

# <none>タグのイメージを削除
docker image prune

# コンテナ、ネットワーク、イメージを一括削除
docker system prune -a

# コンテナ、ネットワーク、イメージ、ボリュームを一括削除
docker system prune -a --volumes
```

## Docker Compose コマンド

```
# コンテナの作成と起動（バックグラウンドで起動）
docker compose up -d

# コンテナの作成と起動（イメージのビルドを実行）
docker compose up -d --build

# 起動
docker compose start

# 停止
docker compose stop

# 再起動
docker compose restart

# 停止＆削除
docker compose down

# コンテナのログを表示
docker compose logs -f

# 起動中のコンテナに入る（bash）
docker compose exec <サービス名> bash

# 起動中のコンテナに入る（ash）
docker compose exec <サービス名> ash

# コンテナ内のコマンドを実行
docker compose run <サービス名> <コマンド>
```
