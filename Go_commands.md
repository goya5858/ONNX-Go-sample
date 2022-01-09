```sh
# モジュールの初期化
>> go mod init Go-img-ONNX

# 外部ライブラリのinstall
>> go get github.com/owulveryck/onnx-go
>> go get github.com/owulveryck/onnx-go/backend/x/gorgonnx

### tensorを手動でインストールしなおすとgorgoniaとバージョン合わなくなるので注意
### これダメ=> >>go get gorgonia.org/tensor
```

"Go-img-ONNX/images" はエラーでgithub.com/owulveryck/onnx-go/internal/x/imagesを使用することができないためソースコードをコピペしてる