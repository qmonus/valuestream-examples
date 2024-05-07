## アプリケーションの説明
本アプリケーションは[QmonusValueStreamチュートリアル](https://docs.valuestream.qmonus.net/tutorials/)で利用するTodoアプリケーションです。


## フォルダ構成の説明
アプリケーションはチュートリアルリポジトリのtodo-appのフォルダに格納されています。
フォルダには、API を提供するバックエンドアプリケーションと、バックエンドアプリケーションと連携する TODO アプリケーションであるフロントエンドアプリケーションがそれぞれ格納されています。

```bash
.
├── .valuestream
│   ├── assemblyline-staging.yaml  
│   └── qvs.yaml  
└── todo-app
    ├── backend  # バックエンドアプリケーション
    ├── frontend  # フロントエンドアプリケーション
    └── shared-infra 
```

バックエンドアプリケーションのフォルダにはバックエンドアプリケーションの実行に必要なソースコード一式と各種パッケージ類、Dockerでバックエンドアプリケーションを動かすためファイル一式が格納されています。

バックエンドアプリケーションフォルダの説明は以下の通りです。

Dockerfile: バックエンドアプリケーションのDockerイメージの内容を定義したファイル
docker-compose.yml: ローカルで実行するコンテナ設定が書かれています。
main.go: バックエンドアプリケーションのメインプログラムファイル
/pkg: バックエンドアプリケーションのプログラムファイル
scripts: テスト用のスクリプトファイル

```bash
.
└── todo-app
    ├── backend
    │   ├── .valuestream  
    │   ├── Dockerfile  # バックエンドアプリケーションのDockerイメージの内容を定義したファイル
    │   ├── Makefile
    │   ├── docker-compose.yml  # ローカルで実行するためのDocker Composeファイル
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go  # バックエンドアプリケーションのメインプログラムファイル
    │   ├── pkg  #  バックエンドアプリケーションのプログラムファイル
    │   │   └── ...
    │   └── scripts  # テスト用のスクリプトファイル
    │       └── ...
```
    
フロントエンドアプリケーションのフォルダにはフロントエンドアプリケーションの実行に必要なソースコード一式と各種パッケージ類が格納されています。

/src : フロントエンドアプリケーション本体
/public : staticファイルとして公開されるフォルダ

```bash
.
└── todo-app
    ├── frontend
    │   ├── .env.development
    │   ├── .env.production 
    │   ├── .valuestream  
    │   ├── Dockerfile
    │   ├── Makefile
    │   ├── README.md
    │   ├── babel.config.js
    │   ├── jsconfig.json
    │   ├── package.json
    │   ├── public  # staticファイルとして公開されるフォルダ
    │   │   └── ...
    │   ├── src  # フロントエンドアプリケーションのプログラムファイル
    │   │   └── ...
    │   ├── vue.config.js
    │   └── yarn.lock
 ```

## ローカルでのテスト実行
このリポジトリはローカルでも実行可能です。
ローカルで実行する場合は以下のコマンドでアプリケーションを立ち上げた後に`http://localhost:8080`にアクセスすることでデプロイするアプリケーションの動作を確認できます。

```
cd todo-app
docker-compose up
```