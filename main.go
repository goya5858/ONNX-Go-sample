package main

import (
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"

	//"github.com/owulveryck/onnx-go/internal/x/images"
	"Go-img-ONNX/images"

	"gorgonia.org/tensor"
)

func main() {
	// I. Inputの作成
	file, fileOpenErr := os.Open("sample.png")
	defer file.Close()
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	}

	img, format, decodeErr := image.Decode(file)
	fmt.Println(format) // PNG
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	height, width := img.Bounds().Dy(), img.Bounds().Dx()
	input := tensor.New(tensor.WithShape(1, 3, height, width), tensor.Of(tensor.Float32))
	convertErr := images.ImageToBCHW(img, input)
	if convertErr != nil {
		fmt.Println("ImageToBCHW error:", convertErr)
	}
	fmt.Println("input size:", input.Size())

	// II. Modelの作成
	//backend := gorgonnx.NewGraph()
	//model := onnx.NewModel(backend)
	//
	//byte_model, _ := os.ReadFile("./SampleModel.onnx")
	//ReadModelErr := model.UnmarshalBinary(byte_model)
	//if ReadModelErr != nil {
	//	fmt.Println(ReadModelErr)
	//}
	//fmt.Println("model loaded")
	//
	//// III. Inference
	//model.SetInput(0, input)
	//fmt.Println("Run model")
	//runErr := backend.Run()
	//if runErr != nil {
	//	log.Fatal(runErr)
	//}
	//fmt.Println("Inference done")
	//output, _ := model.GetOutputTensors()
	//

	//outimg, convertErr2 := images.TensorToImg(output[0])
	outimg, convertErr2 := images.TensorToImg(input)
	if convertErr2 != nil {
		log.Fatal(convertErr2)
	}
	fmt.Println("outIMG:", outimg)

	// IV. 予測の保存
	file_out, createErr := os.Create("./decode.png")
	defer file_out.Close()
	if createErr != nil {
		log.Fatal(createErr)
	}
	png.Encode(file_out, outimg)
}
