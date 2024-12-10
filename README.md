# 使用技術
Go, echo, testing, testfy

# 構成
fibo_api/router：routerの設定を記述する
fibo_api/handler：ルーティングされた後の動きを記述する
fibo_api/utils：汎用的なロジックを記述する

# エンドポイント
http://ec2-43-207-57-8.ap-northeast-1.compute.amazonaws.com/fib

# その他
fibo_api/handler/fobo_test.goでは、外部依存の関数をモック化しようとしたが、間に合わなかったので外部関数を直接使用してテストした。
