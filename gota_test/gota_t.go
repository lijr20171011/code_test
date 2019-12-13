package gota_t

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

//type matrix struct {
//	dataframe.DataFrame
//}
//
//func (m matrix) At(i, j int) float64 {
//	return m.Elem(i, j).Float()
//}
//
//func (m matrix) T() mat.Matrix {
//	return mat.Transpose{m}
//}

func T1() {
	fmt.Println("==========")
	// 新建二维矩阵 每行记录列数必须相等
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)

	//df = dataframe.LoadRecords(
	//	[][]string{
	//		[]string{"A", "B", "C", "D"}, // 表头
	//		[]string{"a", "4", "5.1", "true"},
	//		[]string{"k", "5", "7.0", "true"},
	//		[]string{"k", "4", "6.0", "true"},
	//		[]string{"a", "2", "7.1", "false"},
	//	},
	//	dataframe.DetectTypes(false),        // 不自动检测类型
	//	dataframe.DefaultType(series.Float), // 设置默认类型
	//	dataframe.WithTypes(map[string]series.Type{ // 单独设置列类型
	//		"A": series.String,
	//		"D": series.Bool,
	//	}),
	//)

	//// 根据下标选择行
	//sub := df.Subset([]int{0, 2})
	//fmt.Printf("sub: %v\n", sub)
	//fmt.Println("==========")
	//
	//// 使用索引选择指定列
	//sel := df.Select([]int{0,2})
	//fmt.Printf("sel1: %v\n", sel)
	//fmt.Println("==========")
	//// 使用列名选择指定列
	//sel = df.Select([]string{"A","D"})
	//fmt.Printf("sel2: %v\n", sel)
	//fmt.Println("==========")

	//// 更新下标为0,2行的数据
	//df = df.Set(
	//	[]int{0, 2},
	//	dataframe.LoadRecords([][]string{
	//		[]string{"A", "B", "C", "D"},
	//		[]string{"b", "4", "6.0", "true"},
	//		[]string{"c", "3", "6.0", "false"},
	//	}),
	//)

	//fmt.Printf("%v\n", df)
	//fmt.Println("==========")

	//// 条件查询
	//// 获取A="a"或B>4的记录
	//df = df.Filter(
	//	dataframe.F{"A", series.Eq, "a"},
	//	dataframe.F{"B", series.Greater, 4},
	//)
	//// 并且 D=true 的记录
	//df = df.Filter(
	//	dataframe.F{"D", series.Eq, true},
	//)

	//// 排序 按照 A升序 B降序 排序
	//df = df.Arrange(dataframe.Sort("A"),dataframe.RevSort("B"))

	//// 修改列
	//df = df.Mutate(series.New([]string{"a", "b", "c", "d"}, series.String, "C"))
	//// 增加列
	//df = df.Mutate(series.New([]string{"a", "b", "c", "d"}, series.String, "E"))

	//df := dataframe.LoadRecords(
	//	[][]string{
	//		[]string{"A", "B", "C", "D"},
	//		[]string{"a", "1", "5.1", "true"},
	//		[]string{"k", "5", "7.0", "true"},
	//		[]string{"k", "4", "6.0", "true"},
	//		[]string{"a", "2", "7.1", "false"},
	//	},
	//)
	//df2 := dataframe.LoadRecords(
	//	[][]string{
	//		[]string{"A", "F", "D"},
	//		[]string{"a", "1", "true"},
	//		[]string{"c", "2", "false"},
	//		[]string{"c", "8", "false"},
	//		[]string{"c", "9", "false"},
	//	},
	//)
	//// 内联
	//join := df.InnerJoin(df2, "A")
	//// 左联
	//join = df.LeftJoin(df2, "A")
	//// 右联
	//join = df.RightJoin(df2, "A")

	fmt.Printf("%v\n", df)
	fmt.Println("==========")

	//// 取平均值
	//mean := func(s series.Series) series.Series {
	//	// 将数据转换为float类型(不能转换的位NaN)
	//	floats := s.Float()
	//	sum := 0.0
	//	for _, f := range floats {
	//		sum += f
	//	}
	//	return series.Floats(sum / float64(len(floats)))
	//}
	// 对列进行统计
	//df = df.Capply(mean)
	// 对行进行统计
	//df = df.Rapply(mean)

	fmt.Printf("%v\n", df)
	fmt.Println("==========")
	fmt.Printf("%v\n", df.Err)

	fmt.Println(df.Elem(0, 1))

}
