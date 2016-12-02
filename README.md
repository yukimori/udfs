# memo
sensorbeeのプラグインを作成する

UDFとしてmy_incを登録する。指定された数に対して+1する関数
udf.MustRegisterGlobalUDF("my_inc", udf.MustConvertGeneric(udfs.Inc))

プラグインとして以下が必要
- inc.go
- plugin/plugin.go

goのパッケージとしてダウンロードし、
config/build.yamlを指定して（もしくは同ディレクトリで）build_sensorbeeを実行することで、
プラグインを組み込んだsensorbeeバイナリが作成できる
