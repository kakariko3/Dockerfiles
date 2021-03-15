#!/bin/bash
set -e

# Remove a potentially pre-existing server.pid for Rails.
# Railsのserver.pidが存在している場合、削除する
rm -f /myapp/tmp/pids/server.pid

# Then exec the container's main process (what's set as CMD in the Dockerfile).
# コンテナのメインプロセス（DockerfileでCMDとして設定されているコマンド）を実行
exec "$@"
