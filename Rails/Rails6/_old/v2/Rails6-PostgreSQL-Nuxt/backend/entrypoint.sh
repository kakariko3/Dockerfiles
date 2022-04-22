#!/bin/sh

# エラーが発生した場合、スクリプトを終了する
set -e

# server.pidが存在している場合、削除する
rm -f /app/tmp/pids/server.pid

# コンテナのメインプロセス（DockerfileでCMDとして設定されているコマンド）を実行
exec "$@"
