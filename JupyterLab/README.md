# JupyterLab Setup

## 1. Dockerイメージの作成
```
$ docker build -t jupyter-lab .
```

## 2. Dockerコンテナの起動
```
$ docker run -d --rm -p 8888:8888 -v <作業ディレクトリ>:/work --name <コンテナ名> jupyter-lab
```
