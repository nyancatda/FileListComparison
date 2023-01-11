/*
 * @Author: NyanCatda
 * @Date: 2023-01-12 02:26:20
 * @LastEditTime: 2023-01-12 02:40:17
 * @LastEditors: NyanCatda
 * @Description: 对比数组
 * @FilePath: \FileListComparison\internal\ArrayComparison\ArrayComparison.go
 */
package ArrayComparison

/**
 * @description: 对比数组
 * @param {[]string} Src 源数组
 * @param {[]string} Dest 目标数组
 * @return {[]string} 源数组与目标数组新增的部分
 * @return {[]string} 源数组与目标数组缺少的部分
 */
func Comparison(Src []string, Dest []string) ([]string, []string) {
	// 建立索引
	SrcMap := make(map[string]byte)
	AllMap := make(map[string]byte)

	var Set []string // 交集

	// 源数组建立map
	for _, v := range Src {
		SrcMap[v] = 0
		AllMap[v] = 0
	}

	// 尝试存入目标目标数组，如果无法存入的则为重复
	for _, Value := range Dest {
		AllMapLen := len(AllMap)
		AllMap[Value] = 1
		if AllMapLen == len(AllMap) {
			// 重复部分进入交集
			Set = append(Set, Value)
		}
	}

	// 遍历交集，删除并集中的元素，计算出补集
	for _, Value := range Set {
		delete(AllMap, Value)
	}

	// 此时AllMap为补集，通过对比源数组，如果存在则是缺少的，不存在则是新增的
	var NewElement, MissingElement []string
	for v := range AllMap {
		_, exist := SrcMap[v]
		if exist {
			MissingElement = append(MissingElement, v)
		} else {
			NewElement = append(NewElement, v)
		}
	}

	return NewElement, MissingElement
}
