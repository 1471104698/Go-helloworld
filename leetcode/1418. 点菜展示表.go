package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	strs := [][]string{
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
	}
	fmt.Println(displayTable(strs))
}

func displayTable(orders [][]string) [][]string {
	/*
	   需要存储每个桌子每个菜品点的份数
	   使用两层嵌套 map
	   order[0] 是人名
	   order[1] 是桌号
	   order[2] 是菜品名
	*/
	foodMap := map[string]interface{}{}
	table := map[string]map[string]int{}
	for _, order := range orders {
		// 初始化 order[1]
		if _, exist := table[order[1]]; !exist {
			table[order[1]] = map[string]int{}
		}
		food := table[order[1]]
		food[order[2]]++
		if foodMap[order[2]] == nil {
			foodMap[order[2]] = struct{}{}
		}
	}
	foods := []string{}
	for food, _ := range foodMap {
		foods = append(foods, food)
	}
	// 将菜品名根据名字进行排序
	sort.Strings(foods)

	res := [][]string{}
	for tableID, foodMap := range table {
		curTable := []string{tableID}
		for _, food := range foods {
			curTable = append(curTable, strconv.Itoa(foodMap[food]))
		}
		res = append(res, curTable)
	}
	sort.Sort(Ss(res))
	firstStr := []string{"Table"}
	for _, food := range foods {
		firstStr = append(firstStr, food)
	}
	res = append([][]string{firstStr}, res...)

	return res
}

type Ss [][]string

func (ss Ss) Len() int {
	return len(ss)
}

func (ss Ss) Less(i, j int) bool {
	ii, _ := strconv.Atoi(ss[i][0])
	jj, _ := strconv.Atoi(ss[j][0])
	return ii < jj
}

func (ss Ss) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
