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

#### docker-conpose 環境構築
    cd docker
    docker-compose up -d
    docker-compose ps
    docker-compose down

#### Golang Standard Project Layout
    https://github.com/golang-standards/project-layout
    https://qiita.com/sueken/items/87093e5941bfbc09bea8

#### MySQL
    docker exec -it sample-project_db_1 bash
    mysql -u user -p
    use sample_db;
    show tables;

| headerA | headerB                |
| ------- | ---------------------- |
| short   | very very long content |
