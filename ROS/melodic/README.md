# ROS Melodic Setup
[dorowu/ubuntu-desktop-lxde-vnc](https://hub.docker.com/r/dorowu/ubuntu-desktop-lxde-vnc/)、[Tiryoh/docker_ros-desktop-vnc](https://github.com/Tiryoh/docker-ros-desktop-vnc)をベースに、下記の機能を実装。
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
`docker image ls` コマンドで、Dockerイメージが作成されていることを確認し、下記コマンドを実行する。
```
$ docker run -d -p 6080:80 -p 5900:5900 -v ~/refro_sim/workspace:/home/ubuntu/workspace --name refro_sim --shm-size=512m ubuntu-melodic
```
- `-p 6080:80` <br>
  Webブラウザで http://localhost:6080 へアクセスすることでnoVNCに接続できる。<br>
- `-p 5900:5900` <br>
  VNCクライアントで http://localhost:5900 へアクセスすることでVNCサーバに接続できる。VNC Viewerで接続を確認。<br>
- `-v ~/<作業ディレクトリ>/workspace:/home/ubuntu/workspace` <br>
  -vオプションで共有ディレクトリを指定。
- `--name <コンテナ名>` <br>
  --nameオプションでコンテナ名を指定。
- `--shm-size=512m` <br>
  共有メモリのサイズを指定。デフォルトでは64MB。
- `-e RESOLUTION=1600x900` <br>
  VNC環境の解像度を指定。オプションで指定しない場合、初回接続時のウィンドウサイズに自動で調整される。

## 3. その他
- `$ docker exec -it <コンテナ名> /bin/bash` <br>
  コンテナ起動時、コンテナ内のbashに接続。
- `$ docker stop <コンテナ名>` <br>
  コンテナを停止。
- `$ docker start <コンテナ名>` <br>
  コンテナ停止時、コンテナを起動。
