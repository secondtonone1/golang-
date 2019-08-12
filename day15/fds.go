package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

func eocgongzhen(n float64) float64 {
	var res float64 = 0.945 * (40000*20 - n*n - (float64)(n/3)*(n/3) - (float64)(n/5)*(n/5) -
		(n/8)*(n/8) - (250000-(n-499)*(n-499))*2.2)
	return res
}

func fdsgongzhen(n float64) float64 {
	var res float64 = 1.08722 * (70000*20 - n*n - (float64)(n/3)*(n/3) - (float64)(n/5)*(n/5) -
		(n/8)*(n/8) - (250000-(n-499)*(n-499))*2.2)
	return res
}

func price(i int, count float64) float64 {

	if i < 50 {
		return (14000 / count) * 6.7
	} else if 50 <= i && i < 100 {
		return (16000 / count) * 6.7
	} else if 100 <= i && i < 150 {
		return (16000 / count) * 6.7
	} else if 150 <= i && i < 200 {
		return (16000 / count) * 6.7
	} else if 200 <= i && i < 250 {
		return (16000 / count) * 6.7
	} else if 250 <= i && i <= 300 {
		return (16000 / count) * 6.7
	} else {
		return (15000 / count) * 6.7
	}

	//return (13000 / count) * 6.7

}

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "周期"
	cell = row.AddCell()
	cell.Value = "单价"
	cell = row.AddCell()
	cell.Value = "本期发行量"
	cell = row.AddCell()
	cell.Value = "当前总发行量"
	cell = row.AddCell()
	cell.Value = "本期获利"
	cell = row.AddCell()
	cell.Value = "当前总获利"
	/*
		row1 = sheet.AddRow()
		row1.SetHeightCM(1)
		cell = row1.AddCell()
		cell.Value = "狗子"
		cell = row1.AddCell()
		cell.Value = "18"

		row2 = sheet.AddRow()
		row2.SetHeightCM(1)
		cell = row2.AddCell()
		cell.Value = "蛋子"
		cell = row2.AddCell()
		cell.Value = "28"
	*/

	var sumc float64
	var sump float64
	for i := 1; i <= 350; i++ {
		res := eocgongzhen(float64(i))
		sumc += res
		aa := res * price(i, res)
		sump += aa
		priced := price(i, res)
		fmt.Println(i, priced, res, sumc, aa, sump)
		row = sheet.AddRow()
		row.SetHeightCM(1)

		//周期
		cell = row.AddCell()
		data := strconv.Itoa(i)
		cell.Value = data

		//单价
		cell = row.AddCell()
		data = strconv.FormatFloat(priced, 'E', -1, 64)
		cell.Value = data

		//本期发行数
		cell = row.AddCell()
		data = strconv.FormatFloat(res, 'E', -1, 64)
		cell.Value = data

		//总发行数
		cell = row.AddCell()
		data = strconv.FormatFloat(sumc, 'E', -1, 64)
		cell.Value = data

		//本期获利
		cell = row.AddCell()
		data = strconv.FormatFloat(aa, 'E', -1, 64)
		cell.Value = data

		//历史总获利
		cell = row.AddCell()
		data = strconv.FormatFloat(sump, 'E', -1, 64)
		cell.Value = data
	}

	err = file.Save("gongzhen.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
