著作権は下記の書籍に帰属する
https://www.oreilly.co.jp/books/9784873118048/

### set up for mac

```sh
// please check you already install homebrew
$ which brew
/usr/local/bin/brew

// AND set your $GOPATH ~/go
$ echo $GOPATH
/Users/yourname/go
```

```sh
$ mkdir -p ~/go/src/github.com/ShuzoN && cd $_
$ git clone git@github.com:ShuzoN/real_world_http.git real_world_http && cd $_
$ make install
```

```sh
// boot server
$ make server
```
