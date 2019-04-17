# golang-kafka-example

#### Golang 環境構築（mac os and VSCode）
1. Golangインストール
    *     brew install go
2. PATHの設定
    * view ~/.bash_profile

          export GOPATH=$HOME/go
          export PATH=$GOPATH/bin:$PATH
    * source ~/.bash_profile
3. Visual Studio Code のセットアップ
    1. Go Extensionをインストール
    2. 右下にエラーが出てるので、install を選択

4. 動作確認
    * go run main.go

#### Apache kafka のDocker環境構築
    cd docker
    docker-compose up -d
    docker-compose ps
    docker-compose down


| headerA | headerB                |
| ------- | ---------------------- |
| short   | very very long content |
