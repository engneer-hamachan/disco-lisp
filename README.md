# 型推論付きのLISP方言「DISCO」

## DISCOとは
型推論付きのLISP方言です。
まだまだバグがありそうなのでv0.000001って感じですが、
良かったら覗いてみて下さい。

## DISCOの特徴
### 型推論
[demo.webm](https://github.com/user-attachments/assets/3e4938c1-a68e-4567-9990-0f27be691053)


組み込み関数で型情報を与えて、ひたすら呼び出し元へと伝え続ける「型伝搬法(勝手に命名)」で型情報を管理して、
型エラーを出しています。

### ちょっと今風の構文
``` lisp
(def test (x)
  (with (y (fn (x)
            (if (eq x 10000)
              x
              (funcall #'y (+ x 1)))))
    (funcall #'y x)))

(print (test 1))
```
defunをdefにしたり、代入にもsetfじゃなく=を用いたりしています。
また、condやlet(discoではwith)の余計なカッコを省いたりしています。
*つまりポールグレアムさんのArcをめっちゃ真似てます。

### スタック式VM
discoコードはまずバイトコードにコンパイルされて、その後feverと言うVMで実行されます。
また-tオプションを付ける事で、VMは実行せずにコンパイルだけを行う事が出来ます。（この時型不整合が起きた場合にはエラーで教えてくれます）
``` bash
# 普通に実行
disco sample.dj
# コンパイルのみ実行
disco sample.dj -t
```

## 重要なファイルの説明
- 「disco/lexer」 字句解析器でdiscoコードをtokenにします。
- 「disco/compiler」 コンパイルを行うコードが置いてあります。　compiler/compiler.goが一番mainのファイルになっています。また、１組み込み関数１ファイルで構成されています。
- 「disco/fever」 バイトコードを実行する為のコードが置いてあります。１命令１ファイルで構成されています。

## インストールからハローワールドまで
``` bash
go build disco.go

# editorで以下のコードが書かれたファイルを作成。拡張子は.djです。
# (print "hello world")
vim hello.dj
./disco hello.dj
# => "hello world"
```

## REPL
``` bash
disco
```
とタイプしてEnterを押すとREPLが使えます。

## DISCOのこれから
まだバグに遭遇するので、のんびり一個一個潰しています。
一年くらい遊んでバグが出なくなったら、v1を正式リリースしたい

