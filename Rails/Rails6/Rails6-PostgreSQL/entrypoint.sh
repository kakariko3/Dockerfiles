#!/bin/bash

# エラーが発生した場合、スクリプトを終了する
set -e

# server.pidが存在している場合、削除する
rm -f /myapp/tmp/pids/server.pid

# コンテナのメインプロセス（DockerfileのCMDで渡されたコマンド）を実行
exec "$@"
