package main

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	//启动web服务
	router := gin.Default()
	router.GET("/qr", QRCodeHandler)
	router.Run(":8018")

}
func QRCodeHandler(c *gin.Context) {
	// Create the barcode
	qrCode, _ := qr.Encode("https://zhidao.baidu.com/question/1372858917945947699.html?qbl=relate_question_2&word=%B8%C3%B1%CA%BD%BB%D2%D7%D2%EC%B3%A3%2C%C7%EB%C9%D4%BA%F3h%D4%D9%CA%D4", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 64, 64)

	// // create the output file
	file := bytes.NewBuffer(nil)
	// // encode the barcode as png
	fmt.Println(qrCode)
	png.Encode(file, qrCode)
	// dist := make([]byte, 3000)                    //开辟存储空间
	// base64.StdEncoding.Encode(dist, file.Bytes()) //buff转成base6
	fmt.Println(string(file.Bytes()))
	c.String(http.StatusOK, string(file.Bytes()))
}
