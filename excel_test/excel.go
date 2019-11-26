package excel_test

import (
	"fmt"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

type Request struct {
	IdCode string
	Name   string
}

const (
	fileName = "demo1.xlsx"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func NewExcel() {
	filename := "./demo1.xlsx"
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		panic(err)
	}
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "1234567"

	row = sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell()
	cell.Value = "hehe"
	cell = row.AddCell()
	cell.Value = "159263654"

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Println("文件已经存在!")
		return
	}
	err = file.Save(filename)
	if err != nil {
		panic(err)
	}
}

func ReadExcel() []string {
	filePath := "./" + fileName
	xlFile, err := xlsx.OpenFile(filePath)
	checkErr(err) //自己定义的函数
	//获取行数
	length := len(xlFile.Sheets[0].Rows)
	fmt.Println(length)
	//开辟除表头外的行数的数组内存
	resourceArr := make([]string, length-1)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			// if rowIndex == 0 {
			// 	continue
			// }
			//遍历每一个单元
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text == "haha" {
					fmt.Printf("第 %d 行,第 %d 列的值为 %v: \n", rowIndex, cellIndex, text)
					//如果是每一行的第一个单元格
					fmt.Printf(row.Cells[cellIndex].String())
				}
			}
		}
	}
	return resourceArr
}
