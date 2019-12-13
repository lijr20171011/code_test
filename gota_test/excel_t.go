package gota_t

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/tealeg/xlsx"
	"log"
	"math"
	"time"
)

func T_Excel() {
	fileName := "测试文件.xlsx"
	filePath := `D:\code\gowork\src\mytest\draft_pro\gota_t\` + fileName
	newFilePath := `D:\code\gowork\src\mytest\draft_pro\gota_t\new` + fileName

	startTime := time.Now()
	defer func(start time.Time) {
		fmt.Printf("任务结束,共消耗时间 %s\n", time.Now().Sub(start).String())
	}(startTime)
	// 打开xlsx文件
	f, err := xlsx.OpenFile(filePath)

	if err != nil {
		log.Fatalf("打开xlsx文件异常: %v", err)
	}
	// 获取工作表数据
	sheets, err := f.ToSlice()
	if err != nil {
		log.Fatalf("转换为数组异常: %v", err)
	}
	if len(sheets) == 0 {
		log.Fatalf("获取工作表异常")
	}
	sheet := sheets[0]

	// 加载excel数据生成dataframe
	df := dataframe.LoadRecords(sheet)
	nRow, nCol := df.Dims()
	fmt.Printf("共 [%v:%v]\n", nRow, nCol)
	trueRecords := []int{}  // 全部正确的记录
	falseRecords := []int{} // 出现不匹配的记录

	log.Println(df.Select("大病补贴企缴金额"))

	// 遍历行
	for i := 0; i < nRow; i++ {
		if i%1000 == 0 {
			fmt.Printf("正在比较第[%v]条数据\n", i+1)
		}
		//if i != 1 {
		//	continue
		//}
		match := true
		// 对比每项
		for j := 1; j+2 < nCol; j += 4 {
			if ElemEq(df.Elem(i, j), df.Elem(i, j+1)) {
				df.Elem(i, j+2).Set(1)
			} else {
				match = false
				df.Elem(i, j+2).Set(0)
				//log.Printf("匹配失败 行[%v] 列[%s] 值1[%v(%v)] 值2[%v(%v)] ",
				//	i, df.Names()[j], df.Elem(i, j).Val(), df.Elem(i, j).Type(), df.Elem(i, j+1).Val(), df.Elem(i, j+1).Type())
			}
		}
		if match {
			trueRecords = append(trueRecords, i)
		} else {
			falseRecords = append(falseRecords, i)
		}
	}
	//fmt.Println("===============")
	//fmt.Println(df)

	SliceToExcel(f, "Sheet1", 1, 1, df, newFilePath)
	fmt.Println(df)
	fmt.Println(df.Elem(1, 0).Val())
}

func ElemEq(e1, e2 series.Element) bool {
	// 空数据直接返回
	if e1.Val() == nil && e2.Val() == nil {
		return true
	}
	// 类型一致直接判断
	if e1.Type() == e2.Type() {
		return e1.Eq(e2)
	}
	// 尝试转为float64
	ef1 := e1.Float()
	isE1Float := false
	if e1.Type() != series.String || ef1 != math.NaN() || e1.String() == "" {
		isE1Float = true
	}
	ef2 := e2.Float()
	isE2Float := false
	if e2.Type() != series.String || ef2 != math.NaN() || e2.String() == "" {
		isE2Float = true
	}

	// 同为float判断
	if isE1Float == isE2Float == true {
		return (ef1 == math.NaN() && ef2 == math.NaN()) || ef1 == ef2
	}

	return e1.String() == e2.String()
}

// sRow,sCol 起始行列(从1开始) eg. start(1,1)  df[2,2]  -> end(3,2) 第一行是表头
func SliceToExcel(f *xlsx.File, sheetName string, sRow, sCol int, df dataframe.DataFrame, newFilePath string) {
	log.Printf("起始行列[%v,%v]\n", sRow, sCol)
	dfRow, dfCol := df.Dims()
	head := df.Records()[0]
	//log.Println("表头: ", head)
	// 起始列占一个数据位 需要-1 (1,0) 第一行是表头
	sCol -= 1
	// 计算填充数据需要的最小行列数 (3,2)
	minRow := sRow + dfRow
	minCol := sCol + dfCol
	// 获取工作表
	sheet := f.Sheet[sheetName]
	// 遍历每行
	for i := sRow; i < minRow; i++ {
		if i%1000 == 0 {
			fmt.Printf("正在写入第[%v]条数据\n", i+1)
		}
		// 补充行数
		for sheet.MaxRow <= i {
			sheet.AddRow()
		}
		row := sheet.Rows[i]

		// 遍历每个单元格设置数据
		for j := sCol; j < minCol; j++ {
			// 补充单元格
			for len(row.Cells) <= j {
				row.AddCell()
			}
			if i-sRow == 0 { // 表头行
				row.Cells[j].SetValue(head[j-sCol])
			}
			if df.Elem(i-sRow, j-sCol).Val() == math.NaN() {
				row.Cells[j].SetValue("")
			} else {
				row.Cells[j].SetValue(df.Elem(i-sRow, j-sCol).Val())
			}

		}
	}
	fmt.Println("正在保存文件...")
	err := f.Save(newFilePath)
	if err != nil {
		log.Fatalf("保存文件异常: %v", err)
	}
	fmt.Println("保存完成")
}
