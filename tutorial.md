# DISCO LISP チュートリアル

このチュートリアルでは、型推論付きLISP方言「DISCO」の基本的な使い方から応用までを順を追って説明します。DISCOの特徴である型推論システムを活かしたプログラミングを体験しましょう。

## 目次

1. [インストールと初めての実行](#1-インストールと初めての実行)
2. [基本的な構文と型](#2-基本的な構文と型)
3. [関数の定義と使用](#3-関数の定義と使用)
4. [条件分岐と制御構造](#4-条件分岐と制御構造)
5. [リストの操作](#5-リストの操作)
6. [変数とスコープ](#6-変数とスコープ)
7. [高階関数](#7-高階関数)
8. [マクロ](#8-マクロ)
9. [データ構造](#9-データ構造)
10. [入出力と文字列処理](#10-入出力と文字列処理)
11. [Webアプリケーション](#11-webアプリケーション)
12. [型推論とエラー処理](#12-型推論とエラー処理)

## 1. インストールと初めての実行

### インストール

DISCOをインストールするには、まずGoの開発環境が必要です。Goがインストールされていることを確認してから以下の手順に従ってください。

```bash
# リポジトリをクローン
git clone https://github.com/username/disco.git

# ディレクトリに移動
cd disco

# ビルド
go build disco.go

# 動作確認
./disco -h
```

### Hello World

最初のDISCOプログラムを作成しましょう。テキストエディタで `hello.dj` というファイルを作成し、次のコードを入力します：

```lisp
(print "Hello, World!")
```

このプログラムを実行するには：

```bash
./disco hello.dj
```

次のような出力が表示されるはずです：

```
"Hello, World!"
```

### REPLの使用

対話型環境（REPL）でDISCOを使用することもできます：

```bash
./disco
```

REPLが起動し、プロンプト `(disco) =>` が表示されます。ここで式を入力し、Enterキーを押すと評価結果が表示されます：

```
(disco) => (+ 1 2 3)
6
(disco) => (print "Hello!")
"Hello!"
```

REPLを終了するには、Ctrl+Dを押すか、`(exit)` と入力します。

## 2. 基本的な構文と型

### リテラル値

DISCOは以下の基本的なリテラル型をサポートしています：

```lisp
; 整数
42
-10

; 浮動小数点数
3.14
-0.5

; 文字列
"Hello, World!"

; シンボル
hello
my-variable

; 真値
t

; 偽値（空リスト）
nil

; リスト
(1 2 3 4)
(a b c)

; 空リスト
()
```

### 式の評価

DISCOでは、括弧で囲まれた式はリストとして評価されます。最初の要素は関数名、残りの要素は引数として扱われます：

```lisp
; 加算（1 + 2 + 3）
(+ 1 2 3)  ; 結果: 6

; 文字列の連結
(format "Hello, ~a!" "World")  ; 結果: "Hello, World!"
```

### クォート記法

式を評価せずにリテラルとして扱いたい場合は、クォート記法を使用します：

```lisp
; シングルクォート - 式を評価せずにリテラルとして扱う
'(1 2 3)  ; 評価結果: (1 2 3)

; クォートされていない場合は関数呼び出しとして評価
(list 1 2 3)  ; 評価結果: (1 2 3)
```

### 基本的な演算

```lisp
; 算術演算
(+ 1 2 3)  ; 6
(- 10 5)   ; 5
(* 2 3 4)  ; 24
(/ 10 2)   ; 5
(% 10 3)   ; 1（剰余）

; 比較演算
(eq? 1 1)  ; t（真）
(eq? 1 2)  ; nil（偽）
(> 5 2)    ; t
(< 5 10)   ; t
(>= 5 5)   ; t
(<= 5 5)   ; t

; 論理演算
(and t t nil)  ; nil
(or nil t nil) ; t
(not t)        ; nil
(not nil)      ; t
```

### 型述語

値の型を調べるための関数：

```lisp
(atom? 1)       ; t - 数値はアトム
(atom? '(1 2))  ; nil - リストはアトムではない
(pair? '(1 . 2)) ; t - ドットペア
(list? '(1 2))  ; t - リスト
(str? "hello")  ; t - 文字列
(nil? nil)      ; t - nil
(type 42)       ; "int" - 値の型を文字列で取得
```

## 3. 関数の定義と使用

### 関数定義

DISCOでは `def` を使って関数を定義します：

```lisp
; 2つの数値を加算する関数
(def add (x y)
  (+ x y))

; 関数の呼び出し
(add 2 3)  ; 5
```

### 無名関数（ラムダ）

一時的な関数は `fn` を使って作成できます：

```lisp
; 無名関数の定義と即時呼び出し
((fn (x) (+ x 1)) 5)  ; 6

; 変数に関数を代入
(= add-one (fn (x) (+ x 1)))
(add-one 5)  ; 6
```

### オプショナル引数

オプショナル引数は `&optional` キーワードで定義できます：

```lisp
; zはオプショナル引数
(def greet (name &optional greeting)
  (if (nil? greeting)
    (format "Hello, ~a!" name)
    (format "~a, ~a!" greeting name)))

(greet "John")           ; "Hello, John!"
(greet "John" "Welcome") ; "Welcome, John!"
```

### 関数オブジェクトと高階関数

関数を値として扱うには `#'` 記法を使用します：

```lisp
; 関数オブジェクトの取得
#'add

; 関数オブジェクトの呼び出し
(funcall #'add 2 3)  ; 5

; 関数を引数として受け取る高階関数
(def apply-twice (f x)
  (f (f x)))

(apply-twice #'add-one 5)  ; 7
```

### 再帰関数

DISCOでは再帰関数も簡単に定義できます：

```lisp
; 階乗を計算する再帰関数
(def factorial (n)
  (if (= n 0)
    1
    (* n (factorial (- n 1)))))

(factorial 5)  ; 120
```

## 4. 条件分岐と制御構造

### if式

最も基本的な条件分岐は `if` です：

```lisp
(if (> 5 3)
  "5は3より大きい"
  "5は3以下")  ; "5は3より大きい"
```

### unless式

`if` の否定版である `unless` も利用できます：

```lisp
(unless (> 3 5)
  "3は5より大きくない")  ; "3は5より大きくない"
```

### cond式

複数の条件分岐には `cond` を使用します：

```lisp
(def check-number (n)
  (cond (< n 0)  "負の数"
        (= n 0)  "ゼロ"
        (> n 0)  "正の数"))

(check-number -5)  ; "負の数"
(check-number 0)   ; "ゼロ"
(check-number 10)  ; "正の数"
```

### case式

値に基づく分岐には `case` を使用します：

```lisp
(def day-name (day)
  (case day
    (1 "月曜日")
    (2 "火曜日")
    (3 "水曜日")
    (4 "木曜日")
    (5 "金曜日")
    (6 "土曜日")
    (7 "日曜日")
    ('default "不明な日")))

(day-name 1)  ; "月曜日"
(day-name 9)  ; "不明な日"
```

### 論理演算子を使った制御フロー

```lisp
; andは最初のnilを返すか、最後の式の値を返す
(and 1 2 3)    ; 3
(and 1 nil 3)  ; nil

; orは最初のnilでない値を返す
(or nil 2 3)   ; 2
(or nil nil 5) ; 5
```

## 5. リストの操作

### リストの作成

```lisp
; クォートを使用
'(1 2 3)

; list関数を使用
(list 1 2 3)

; consを使用
(cons 1 (cons 2 (cons 3 nil)))
```

### 基本的なリスト操作

```lisp
; 先頭要素を取得
(car '(1 2 3))  ; 1

; 残りの要素を取得
(cdr '(1 2 3))  ; (2 3)

; 先頭に要素を追加
(cons 0 '(1 2 3))  ; (0 1 2 3)

; リストの連結
(append '(1 2) '(3 4))  ; (1 2 3 4)

; リストの長さを取得
(length '(1 2 3 4))  ; 4

; インデックスで要素を取得
(nth 2 '(a b c d))  ; c（0から始まるインデックス）

; リスト内の要素を検索
(member 'b '(a b c))  ; (b c)（見つかった位置からの残りのリスト）
```

### 入れ子リストのアクセス

```lisp
; car/cdrの組み合わせ
(car (cdr '(a b c)))  ; b（cadrと同じ）

; 短縮形
(cadr '(a b c))       ; b（2番目の要素）
(caddr '(a b c))      ; c（3番目の要素）
(caar '((a b) c))     ; a（入れ子リストの最初の要素）
```

### リスト内包表記（mapcar）

```lisp
; 各要素に1を加える
(mapcar #'(fn (x) (+ x 1)) '(1 2 3))  ; (2 3 4)

; 組み込み関数を使用
(mapcar #'+ '(1 2 3) '(10 20 30))     ; (11 22 33)
```

## 6. 変数とスコープ

### グローバル変数

```lisp
; グローバル変数の定義
(= count 0)

; global特殊形式を使った定義
(global *max-count* 100)
```

### ローカル変数（with）

```lisp
; ローカルスコープでの変数定義
(with (x 10
       y 20)
  (+ x y))  ; 30（スコープ外ではxとyは見えない）
```

### 変数の更新

```lisp
; 変数の更新
(= counter 0)
(= counter (+ counter 1))  ; 1
```

### クロージャ

```lisp
; カウンターを作成する関数
(def make-counter ()
  (with (count 0)
    (fn ()
      (= count (+ count 1)))))

; カウンターを作成
(= my-counter (make-counter))

; カウンターを使用
(funcall my-counter)  ; 1
(funcall my-counter)  ; 2
(funcall my-counter)  ; 3
```

## 7. 高階関数

### 関数を返す関数

```lisp
; 加算器を作成する関数
(def make-adder (n)
  (fn (x) (+ x n)))

; 10を加える関数を作成
(= add10 (make-adder 10))

; 関数を使用
(add10 5)  ; 15
```

### 関数を引数として取る関数

```lisp
; 関数を2回適用する
(def apply-twice (f x)
  (f (f x)))

; 使用例
(apply-twice #'add10 5)  ; 25（5 + 10 + 10）

; 関数合成
(def compose (f g)
  (fn (x) (f (g x))))

; 使用例
(= add10-and-double (compose #'(fn (x) (* x 2)) #'add10))
(add10-and-double 5)  ; 30（(5 + 10) * 2）
```

### 高階関数の実用例

```lisp
; フィルター関数
(def filter (pred lst)
  (if (nil? lst)
    nil
    (if (funcall pred (car lst))
      (cons (car lst) (filter pred (cdr lst)))
      (filter pred (cdr lst)))))

; 使用例：偶数のフィルタリング
(filter #'even? '(1 2 3 4 5 6))  ; (2 4 6)

; 畳み込み関数（reduce）
(def reduce (f init lst)
  (if (nil? lst)
    init
    (reduce f (funcall f init (car lst)) (cdr lst))))

; 使用例：合計計算
(reduce #'+ 0 '(1 2 3 4 5))  ; 15
```

## 8. マクロ

### マクロの基本

マクロはコードを生成するコードです。DISCOでは `mac` を使ってマクロを定義します：

```lisp
; whenマクロの定義（ifの便利ラッパー）
(mac when (test &rest body)
  `(if ,test 
     (progn ,@body)))
```

この例では、以下のクォート記法を使用しています：

- `` `form `` (バッククォート)：式をクォートする
- `,form` (カンマ)：クォート内で式を評価する
- `,@form` (カンマアットマーク)：クォート内でリストをスプライスする

### マクロの使用例

```lisp
; whenマクロの使用
(when (> 5 3)
  (print "5は3より大きい")
  (print "これは真です"))

; 上記は次のコードに展開される：
; (if (> 5 3)
;   (progn 
;     (print "5は3より大きい")
;     (print "これは真です")))
```

### 実用的なマクロ例

```lisp
; aifマクロ（anaphoric if）- テスト結果をitにバインド
(mac aif (test then &optional else)
  `(with (it ,test)
     (if it ,then ,else)))

; 使用例
(aif (/ 10 2)
     (print (format "結果は~aです" it)))  ; "結果は5です"

; nil!マクロ - 変数をnilに設定
(mac nil! (var)
  `(= ,var nil))

; 使用例
(= x 10)
(nil! x)
(print x)  ; nil
```

## 9. データ構造

### ベクター

ベクターは固定長の配列として機能し、インデックスによる高速なアクセスが可能です：

```lisp
; ベクターの作成
(= vec (vector 1 2 3 4 5))

; 要素へのアクセス（0ベース）
(aref 0 vec)  ; 1
(aref 2 vec)  ; 3

; シンタックスシュガー（これも同じ）
(vec 0)  ; 1

; ベクターの長さを取得
(vector-len vec)  ; 5

; 要素の追加（新しいベクターを返す）
(= vec (vector-push vec 6))

; 最後の要素を取り出し（破壊的）
(vector-pop vec)  ; 6
```

### ハッシュテーブル

キーと値のペアを格納するデータ構造：

```lisp
; ハッシュテーブルの作成
(= table (make-hash))

; 値の設定
(sethash 'name table "John")
(sethash 'age table 30)

; シンタックスシュガー（これも同じ）
(table 'name "John")
(table 'age 30)

; 値の取得
(gethash table 'name)  ; "John"

; シンタックスシュガー（これも同じ）
(table 'name)  ; "John"
```

## 10. 入出力と文字列処理

### ファイル操作

```lisp
; ファイルを書き込みモードでオープン
(= file (open "example.txt" "w"))

; ファイルに書き込み
(write file "Hello, World!\nThis is a test.")

; ファイルを読み込みモードでオープン
(= file (open "example.txt" "r"))

; ファイルの内容を読み込み
(= content (read-file file))

; ファイルを閉じる
(close file)
```

### 文字列処理

```lisp
; 文字列のフォーマット
(format "Hello, ~a!" "World")  ; "Hello, World!"
(format "Count: ~d" 42)        ; "Count: 42"

; 文字列の分割
(split " " "Hello World Test") ; ("Hello" "World" "Test")

; 部分文字列の取得（開始位置、長さ）
(subseq "Hello" 1 2)           ; "el"

; 文字列をリストに変換
(str-to-list "abc")            ; ("a" "b" "c")

; 正規表現によるマッチング
(regexp-match "Hello" "H")     ; t

; 正規表現による置換
(regexp-replace "Hello" "H" "J") ; "Jello"
```

### 標準入出力

```lisp
; 値の出力（引用符付き）
(print "Hello")    ; "Hello"

; 値の出力（引用符なし）
(princ "Hello")    ; Hello

; 値の出力（改行なし）- エイリアス：pc
(pc "Hello ")
(pc "World!")      ; Hello World!
```

### コマンド実行

```lisp
; 外部コマンドの実行と結果の取得
(command "ls" "-l")  ; ディレクトリリストの文字列

; 複数の引数を持つコマンド
(command "find" "." "-name" "*.txt")
```

## 11. Webアプリケーション

DISCOは簡易的なWebサーバー機能を内蔵しています：

### シンプルなWebサーバー

```lisp
; ハンドラー関数の定義
(def index-handler ()
  (format "<html><body><h1>Hello, World!</h1></body></html>"))

; サーバーの定義
(defserver 'myserver ":8080")

; ルートへのハンドラー登録
(defhandler "/" #'index-handler)

; サーバーの起動
(run-server 'myserver)
```

### フォーム処理

```lisp
; フォームを表示するハンドラー
(def form-handler ()
  (format "<html><body>
    <form action='/submit' method='get'>
      <input name='name' type='text'>
      <input type='submit' value='Submit'>
    </form>
  </body></html>"))

; フォーム送信を処理するハンドラー
(def submit-handler ()
  (format "<html><body>
    <h1>Hello, ~a!</h1>
  </body></html>" (get-query 'name)))

; ハンドラーの登録
(defhandler "/form" #'form-handler)
(defhandler "/submit" #'submit-handler)
```

### HTTPリクエスト

```lisp
; GETリクエストの作成
(= req (make-request "GET" "https://example.com"))

; ヘッダーの追加
(= req (add-request-header req "User-Agent" "DISCO/1.0"))

; リクエストの送信と結果の取得
(= response (do-request req))
```

## 12. 型推論とエラー処理

### 型推論の理解

DISCOの特徴である型推論システムを理解しましょう：

```lisp
; 関数の定義（引数の型は推論される）
(def add (x y)
  (+ x y))  ; + は数値型を期待するため、xとyは数値型と推論される

; 正しい使用法
(add 1 2)  ; 3

; 誤った使用法（コンパイルエラー）
; (add "a" 2)  ; エラー: "a" is invalid argument type string want int
```

### コンパイル時型チェック

コンパイル時の型チェックを試すには `-t` オプションを使用します：

```bash
./disco myfile.dj -t
```

これにより、実行せずに型チェックのみが行われます。

### 型エラーのデバッグ

型エラーが発生した場合、以下のような情報が表示されます：

```
filename.dj::10::variable x is string want int
```

このエラーは、10行目で変数 `x` が文字列型だが、整数型が期待されていることを示しています。

### 実行時エラー処理

```lisp
; errorを使用してエラーを発生させる
(error "エラーが発生しました")

; エラーハンドリングの例（try/catchに相当する機能は現在ありません）
(def safe-divide (a b)
  (if (zero? b)
    (progn
      (print "ゼロ除算エラー")
      nil)
    (/ a b)))
```

### 型を意識したプログラミング

```lisp
; 型が混在する関数の例
(def process-input (input)
  (if (str? input)
    (length input)        ; 文字列の場合、長さを返す
    (if (list? input)
      (car input)         ; リストの場合、最初の要素を返す
      (+ input 1))))      ; 数値の場合、1を加算

; 使用例
(process-input "hello")  ; 5
(process-input '(a b c)) ; a
(process-input 10)       ; 11
```

## まとめ

このチュートリアルでは、DISCO LISPの基本的な使い方から応用までを学びました。DISCOの特徴である型推論システムを活かすことで、静的型付け言語の安全性と動的言語の柔軟性を兼ね備えたプログラミングが可能です。

さらに詳しく知りたい場合は、公式ドキュメントや実装コードを参照してください。

## 次のステップ

- 独自のプロジェクトでDISCOを使ってみる
- 標準ライブラリの関数を詳しく調べる
- DISCOの拡張や貢献を検討する

Happy Hacking with DISCO!