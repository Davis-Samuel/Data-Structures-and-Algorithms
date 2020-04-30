package _3_动态规划

import "fmt"

//最长公共子串(不是公共子序列)
//如fosh,fish 最长公共子串是sh, 公共子序列是fsh
//先抽象成数据模型. 利用矩阵查找关系
//思路: 先目标字符串抽象成字符行, 匹配字符串抽象成字符列. 行与列就成了矩阵
// 若row为行, col为列,比较之间的关系.
// 如果相等则选择左上方单元格的值加1
// 如果不相等则记作0
func SubString(a, b string) string {
	rChar := []byte(a)
	cChar := []byte(b)
	iRow, jCol := len(rChar), len(cChar)
	var matrix [][]int //申请一个i行j列的矩阵
	matrix = make([][]int, iRow+1)
	//设置一个哨兵
	matrix[0] = make([]int, jCol+1) //从1下标开始.减少判断
	for row := 1; row < iRow; row ++ {
		matrix[row] = make([]int, jCol+1)
		for col := 1; col < jCol; col ++ {
			if rChar[row] == cChar[col] {//i行j列相等时则左上方单格的值加1.
				matrix[row][col] = matrix[row-1][col-1] + 1
			} else {//i行j列不相等时,则选择上方和左方邻居中较大的那个.
				matrix[row][col] = 0 //区别于求公共子序列.
			}
		}
	}
	pArr := make([]byte, 0)//存储公共字符
	count := 0
	for _, rows := range matrix {
		for key, val := range rows {
			if val > count {
				count = val
				pArr = append(pArr, cChar[key])
			}
		}
	}
	//获取矩阵里最大的值.
	printMatrix(rChar, cChar, matrix)
	fmt.Printf("%s与%s的公共字串:%c\n", a, b, pArr)
	return string(pArr)
}
//打印矩阵,显示子串的关系
func printMatrix(rowChar, colChar []byte, mtr [][]int) {
	for r, rows := range mtr {
		if r == 0 {
			fmt.Print("  ")
			for _, rc := range colChar {
				fmt.Print(string(rc), " ")
			}
			fmt.Println()
		}
		for c, col := range rows {
			if c == 0 {
				continue
			}
			if c == 1 {
				fmt.Printf("%c ", rowChar[r])
			}
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
