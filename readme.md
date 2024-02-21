# anypwa


## 使い方 


```sh
$ ./anypwa.exe
Usage: anypwa [--url URL] [--width WIDTH] [--height HEIGHT] [--title TITLE] [--debug] [--version] <command> [<args>]

Options:
  --url URL, -u URL      開きたいURL
  --width WIDTH, -x WIDTH
                         初期ウィンドウ幅 [default: 1000]
  --height HEIGHT, -y HEIGHT
                         初期ウィンドウ高さ [default: 800]
  --title TITLE, -t TITLE
                         ウィンドウタイトル
  --debug, -d            デバッグ用。ログが詳細になる。
  --version, -v          バージョン情報を出力する。
  --help, -h             ヘルプを出力する。

Commands:
  version                バージョン情報を出力する。

```


開きたいURLが `https://www.google.co.jp/` (これはprivateな私のToDoリストProjects)だとする。  
この時、  
```sh
$ ./anypwa.exe -u "https://www.google.co.jp/"
```
のように実行する。  
恐らくWindowsだとEdgeを使用してURLが開かれる。  
これだけだとブラウザのショートカットと変わらない。  


## PWAの代替

Chromeなどで、`https://github.com/xcd0/anypwa`を開き、PWAとしてインストールすると、  
`https://github.com/`が開かれてしまう。  

ToDoリスト用に作成したGitHubのProjectsを開くPWAを作りたかった。  
しかし、PWAのURLを編集する方法が見つからなかった。  

その為、これを使用する。

windows上で、`anypwa.exe`へのショートカットを作成し、  
ショートカットの名前を`ToDo - GtiHub.com`
ショートカットのプロパティを開く。


リンク先の項目が、`"anypwa.exeへのフルパス"`になっているはずなので、個の末尾に
` -u "ToDoリスト用に作成したGitHubのProjectsのURL" -t "ToDo - GitHub Projects"`
を追加する。  
序に、ショートカットのアイコンを[GitHubのアイコン](./GitHub.ico)に設定する。  

これで、ToDoリスト用に作成したGitHubのProjectsをPWAアプリ風に開くショートカットができる。  

