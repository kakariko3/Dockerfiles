# ROS Melodic Setup
このDockerfileは、[dorowu/ubuntu-desktop-lxde-vnc](https://hub.docker.com/r/dorowu/ubuntu-desktop-lxde-vnc/)及び、[Tiryoh/docker_ros-desktop-vnc](https://github.com/Tiryoh/docker-ros-desktop-vnc)をベースに、下記の機能を実装している。
- Ubuntu 18.04 LXQt
- VNCサーバ
- VNCクライアント (noVNC)、HTTPサーバ
- ROS Melodic

## 1. Dockerイメージのビルド
Dockerfileがあるディレクトリに移動し、下記コマンドを実行する。
```
$ docker build -t ubuntu-melodic .
```

## 2. Dockerコンテナの起動
`docker image ls` コマンドで、`ubuntu-melodic` というDockerイメージが作成されていることを確認し、下記コマンドを実行する。
```
$ docker run -d -p 6080:80 -p 5900:5900 -v ~/workspace:/home/ubuntu/workspace --name ubuntu_ros --shm-size=512m ubuntu-melodic
```
- `-p 6080:80` <br>
  Webブラウザで http://localhost:6080 へアクセスすることで、noVNCに接続できる。<br>
- `-p 5900:5900` <br>
  VNCクライアントで http://localhost:5900 へアクセスすることで、VNCサーバに接続できる。VNC Viewerで接続を確認。<br>
- `-v <作業ディレクトリ>/workspace:/home/ubuntu/workspace` <br>
