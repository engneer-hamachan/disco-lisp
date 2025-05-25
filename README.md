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
また、condやlet(discoではwith)の余計なカッコを省いてます。
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
ビルドしてから
``` bash
./disco
```
とタイプしてEnterを押すとREPLが使えます。

## DISCOのこれから
まだバグに遭遇するので、のんびり一個一個潰しています。
一年くらい遊んでバグが出なくなったら、v1を正式リリースしたい


以下、AIにも怒気ぃ書いて貰いました！AIが書いたドキュメントの方が優秀ですね（汗）

# DISCO LISP ドキュメント

## 概要

DISCO（Dynamic Intelligent Statically-typed COmpiler）は型推論機能を持つLISP方言です。独自の「型伝搬法」を用いた静的型検査システムにより、実行前にコンパイル時の型エラーを検出します。また、DISCOコードはFeverと呼ばれるスタックベースのVMで実行されるバイトコードに変換されます。

## 特徴

### 型推論システム

DISCOの最大の特徴は、明示的な型宣言なしに型の整合性を保証する型推論システムです。組み込み関数から型情報を伝搬させ、不整合があればコンパイル時にエラーを検出します。

例：
```lisp
(def add (x y)
  (+ x y))      ; + は数値型の引数を期待するため、xとyは数値型と推論される

(add 1 2)       ; OK
(add "a" 2)     ; コンパイルエラー: "a" is invalid argument type string want int
```

### シンプルな構文

古典的なLispと比較して、余分な括弧を省略した読みやすい構文を採用しています：

```lisp
; 従来のLisp
(defun test (x)
  (let ((y (lambda (x)
             (if (eq x 10000)
                 x
                 (funcall y (+ x 1))))))
    (funcall y x)))

; DISCO
(def test (x)
  (with (y (fn (x)
            (if (eq? x 10000)
              x
              (funcall #'y (+ x 1)))))
    (funcall #'y x)))
```

### バイトコードコンパイラとVM

DISCOはコードをバイトコードにコンパイルし、Feverという専用VMで実行します。これにより、実行時のパフォーマンスが向上します。

## インストール方法

### ビルド

```bash
go build disco.go
```

### システムへのインストール（オプション）

```bash
sudo mv ./disco /usr/local/bin/disco
```

## 使い方

### ファイルの実行

```bash
./disco filename.dj
```

### コンパイルのみ（型チェック）

```bash
./disco filename.dj -t
```

### バイトコードのダンプ

```bash
./disco filename.dj -d
```

### REPL（対話モード）

```bash
./disco
```

## 言語仕様

### 基本データ型

- **整数** `int`: `1`, `42`, `-100`
- **浮動小数点数** `float`: `1.0`, `3.14`, `-0.5`
- **文字列** `string`: `"hello"`, `"world"`
- **シンボル** `symbol`: `x`, `my-var`, `+`
- **リスト** `list`: `(1 2 3)`, `(a b c)`
- **ベクター** `vector`: `(vector 1 2 3)`
- **ハッシュテーブル** `hash`: `(make-hash)`
- **真値** `true`: `t`
- **偽値** `nil`: `nil`

### 特殊形式

- **def**: 関数定義
- **with**: ローカル変数束縛（letに相当）
- **if**: 条件分岐
- **unless**: ifの否定形
- **cond**: 複数条件分岐
- **case**: 値に基づく分岐
- **mac**: マクロ定義
- **global**: グローバル変数定義
- **=**: 変数代入

### マクロシステム

DISCOはLisp-1型のマクロシステムを持っています：

```lisp
(mac when (test &rest body)
  `(if ,test 
     (progn ,@body)))
```

クォート記法：
- `` `form ``: クォート（quasi-quote）
- `,form`: アンクォート（unquote）
- `,@form`: スプライシングアンクォート（unquote-splicing）

## 組み込み関数一覧

### 算術演算子
- **+**: 加算 `(+ 1 2 3)` → `6`
- **-**: 減算 `(- 10 5)` → `5`
- **\***: 乗算 `(* 2 3 4)` → `24`
- **/**: 除算 `(/ 10 2)` → `5`
- **%**: 剰余 `(% 10 3)` → `1`

### 比較演算子
- **eq?**: 等価比較 `(eq? 1 1)` → `t`
- **>**: 大なり `(> 5 3)` → `t`
- **<**: 小なり `(< 3 5)` → `t`
- **>=**: 以上 `(>= 5 5)` → `t`
- **<=**: 以下 `(<= 5 5)` → `t`

### 論理演算子
- **and**: 論理積 `(and t t nil)` → `nil`
- **or**: 論理和 `(or nil nil t)` → `t`
- **not**: 論理否定 `(not t)` → `nil`

### 型述語
- **atom?**: アトムか判定 `(atom? 1)` → `t`
- **pair?**: ドットペアか判定 `(pair? (cons 1 2))` → `t`
- **list?**: リストか判定 `(list? '(1 2))` → `t`
- **nil?**: nilか判定 `(nil? nil)` → `t`
- **str?**: 文字列か判定 `(str? "hello")` → `t`
- **zero?**: ゼロか判定 `(zero? 0)` → `t`
- **even?**: 偶数か判定 `(even? 2)` → `t`
- **odd?**: 奇数か判定 `(odd? 3)` → `t`
- **type**: 値の型を取得 `(type 42)` → `"int"`

### リスト操作
- **car**: リストの先頭要素を取得 `(car '(1 2 3))` → `1`
- **cdr**: リストの残りを取得 `(cdr '(1 2 3))` → `(2 3)`
- **cons**: 要素とリストを結合 `(cons 1 '(2 3))` → `(1 2 3)`
- **list**: リスト作成 `(list 1 2 3)` → `(1 2 3)`
- **append**: リスト結合 `(append '(1 2) '(3 4))` → `(1 2 3 4)`
- **length**: 長さ取得 `(length '(1 2 3))` → `3`
- **nth**: インデックスで要素取得 `(nth 1 '(a b c))` → `b`
- **member**: 要素探索 `(member 'b '(a b c))` → `(b c)`
- **assoc**: 連想リスト探索 `(assoc 'b '((a 1) (b 2)))` → `(b 2)`
- **mapcar**: 各要素に関数適用 `(mapcar #'1+ '(1 2 3))` → `(2 3 4)`

### 文字列操作
- **format**: 文字列フォーマット `(format "Hello ~a" "World")` → `"Hello World"`
- **regexp-match**: 正規表現マッチ `(regexp-match "abc" "a")` → `t`
- **regexp-replace**: 正規表現置換 `(regexp-replace "abc" "a" "X")` → `"Xbc"`
- **split**: 文字列分割 `(split " " "a b c")` → `("a" "b" "c")`
- **subseq**: 部分文字列 `(subseq "abcde" 1 3)` → `"bc"`
- **str-to-list**: 文字列をリストに変換 `(str-to-list "abc")` → `("a" "b" "c")`
- **intern**: 文字列からシンボル作成 `(intern "abc")` → `abc`

### ベクター操作
- **vector**: ベクター作成 `(vector 1 2 3)` → `#VECTOR(1 2 3)`
- **aref**: ベクター要素アクセス `(aref 0 (vector 1 2 3))` → `1`
- **vector-push**: ベクターに要素追加 `(vector-push (vector 1 2) 3)` → `#VECTOR(1 2 3)`
- **vector-pop**: ベクターから要素取得と削除 `(vector-pop (vector 1 2 3))` → `3`
- **vector-len**: ベクター長さ取得 `(vector-len (vector 1 2 3))` → `3`

### ハッシュテーブル操作
- **make-hash**: ハッシュテーブル作成 `(make-hash)`
- **gethash**: ハッシュテーブルから値取得 `(gethash table 'key)` → `value`
- **sethash**: ハッシュテーブルに値設定 `(sethash 'key table value)` → `table`

### 入出力
- **print**: 値の表示（改行付き） `(print "hello")` → `"hello"`
- **princ**: 値の表示（引用符なし） `(princ "hello")` → `hello`
- **open**: ファイルオープン `(open "file.txt" "r")`
- **close**: ファイルクローズ `(close file)`
- **read-file**: ファイル読み込み `(read-file file)` → `"content"`
- **write**: ファイル書き込み `(write file "content")`

### その他
- **random**: 乱数生成 `(random 10)` → `7`
- **time**: 処理時間計測 `(time (+ 1 2))` → `3` と時間表示
- **command**: 外部コマンド実行 `(command "ls" "-l")` → `"..."` (コマンド出力)
- **error**: エラー発生 `(error "message")`
- **json-parse**: JSON解析 `(json-parse "{\"a\":1}")` → `(("a" 1))`
- **funcall**: 関数オブジェクト呼び出し `(funcall #'+ 1 2)` → `3`

### Webサーバー関連
- **defserver**: サーバー定義 `(defserver 'web ":8080")`
- **defhandler**: ハンドラー定義 `(defhandler "/" #'index-handler)`
- **run-server**: サーバー起動 `(run-server 'web)`
- **set-status**: HTTPステータス設定 `(set-status 200)`
- **get-query**: クエリパラメータ取得 `(get-query 'name)` → `"value"`
- **make-request**: HTTPリクエスト作成 `(make-request "GET" "https://example.com")`
- **do-request**: HTTPリクエスト実行 `(do-request request)` → `"response"`
- **add-request-header**: リクエストヘッダー追加 `(add-request-header request "Content-Type" "application/json")`

## コード例

### フィボナッチ数列

```lisp
(def fib (n)
  (if (< n 2)
    n
    (+ (fib (- n 1)) (fib (- n 2)))))

(print (fib 10))  ; 55
```

### テールコール最適化の例

```lisp
(def test (x)
  (with (y (fn (x)
            (if (eq? x 10000)
              x
              (funcall #'y (+ x 1)))))
    (funcall #'y x)))

(print (test 1))  ; 10000
```

### Webサーバーの例

```lisp
(def input-form ()
  (format "<input name='name' type='text'>"))

(def index-handler () 
  (set-status 200)
  (format "<html><body>~a</body></html>" (input-form)))

(def test-handler () 
  (set-status 200)
  (format "Hello, ~a!" (get-query 'name)))

(defserver 'web ":8080")
(defhandler "/" #'index-handler)
(defhandler "/test" #'test-handler)
(run-server 'web)
```

## エラーハンドリング

DISCOでは、コンパイル時と実行時の両方でエラーが検出されます。

### コンパイル時エラー

特に型エラーはコンパイル時に検出されます：

```
test.dj::10::add is int not string
```

### 実行時エラー

実行時エラーはVM（Fever）によって検出・報告されます：

```
test.dj::15::division by zero
```

## 今後の開発予定

DISCOはまだ開発初期段階にあり、以下のような機能追加や改善が予定されています：

1. パフォーマンス最適化
2. より強力な型推論
3. ユーザー定義型のサポート
4. ライブラリシステムの拡充
5. デバッグツールの改善

## 注意点

- DISCOはまだ実験的な言語であり、API変更や仕様変更が頻繁に行われる可能性があります
- プロダクション環境での使用は推奨されていません
- バグ報告や機能提案は歓迎しています

## 参考リンク

- [GitHub リポジトリ](https://github.com/username/disco)
- [チュートリアル](https://github.com/username/disco/docs/tutorial.md)
- [サンプルコード集](https://github.com/username/disco/examples)

