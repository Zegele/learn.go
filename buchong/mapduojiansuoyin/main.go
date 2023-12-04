// go语言map的多键索引——多个数值条件可以同时查询
// www.kancloud.cn/imdszxs/golang/1509741
package main

import "fmt"

//大多数语言，映射容器的键必须以单一值存在
//但随着查询条件越来越复杂，检索也会变得越发困难

// 人员档案：
type Profile struct {
	Name    string // 名字
	Age     int
	Married bool
}

//并且准备好一堆原始数据，需要算法实现构建和查询的过程，如下：

func main() {
	//list := []*Profile{
	//	{Name: "张三", Age: 30, Married: true},
	//	{Name: "王码字", Age: 21},
	//}

	//buildIndex(list)
	//queryData("张三", 30)
}

// 需要用算法实现buildIndex()构建索引函数及queryData()查询数据函数
//查询到结果后将数据打印出来
//下面，分别基于传统的哈希值的多键索引和利用map特性的多键索引进行查询

/*
//一、基于哈希值的多键索引及查询
//传统的数据索引过程是将输入的数据做特征值，这种特征值有几种常见做法:
//1. 将特征使用某种算法转为整数，即哈希值，使用整型值做索引
//2. 将特征转为字符串，使用字符串做索引
//
//数据都基于特征值构建好索引后，就可以进行查询
//查询时，重复这个过程，将查询条件转为特征值，使用特征值进行查询得到结果
//基于哈希的传统多键索引和查询的完整代码位于：//pan.baidu.com/s/1ORFVTOLEYYqDhRzeq0zIiQ    提取密码：hfyf
//1）紫都城转哈希值
//本例中，查询键（classicQueryKey）的特征值需要将查询键中每一个字段转换为整型
//字符串也需要转换为整型值 ，这里使用一种简单算法将字符串转换为需要的哈希值，如下：
func simpleHash(str string) (ret int) {
	// 遍历字符串中的每一个ASCII字符
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i] // c的类型是uint8 也就是byte类型
		// 将字符的ASCII码相加
		ret += int(c)
	}
	return
}

// 哈希算法有很多，这里只是选用一种大家便于理解的算法
//哈希算法的选用标准是尽量减少重复键的发生，俗称“哈希碰撞”
//即同样两个字符的哈希值重复率降低到最低。

//2）查询键
//有了哈希算法函数后，将哈希函数用在查询键结构中。查询键结构如下：
//查询键
type classicQueryKey struct {
	Name string // 要查询的名字
	Age  int    // 要查询的年龄
}

// 计算查询键的哈希值
func (c *classicQueryKey) hash() int {
	//将名字的Hash和年龄哈希合并
	return simpleHash(c.Name) + c.Age*1000000
}

// 获得查询键的哈希值

//3) 构建索引
//本例需要快速查询，因此需要提前对已有的数据构建索引
//前面已经准备好了数据查询键，使用查询键获得哈希即可对数据进行快速索引
//如下：
//创建哈希知道数据的索引关系
var mapper = make(map[int][]*Profile)

//构建数据索引
func buildIndex(list []*Profile) {
	//遍历所有的数据
	for _, profile := range list {
		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}
		// 计算数据的哈希值，取出已经存在的记录
		existValue := mapper[key.hash()] // 使用查询键的哈希方式计算查询键的哈希值，通过这个值，在mapper索引中查找相同哈希值的数据切片集合
		//因为哈希函数不能保证不同数据的哈希值一定完全不同，因此要考虑在发生哈希值重复时的处理办法

		// 将当前数据添加到已经存在的记录切片中
		existValue = append(existValue, profile)
		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

//这种多键的算法就是哈希算法
//map的多个元素对应哈希的“桶”
//哈希函数的选择决定桶的映射好坏，如果哈希冲撞很厉害
//那么就需要将发生冲撞的相同哈希值的元素使用切片保存起来
//
//4）查询逻辑
//从已经构建好索引的数据中查询需要的数据流程序如下：
//1. 给定查询条件（名字、年龄）
//2. 根据查询条件构建查询键
//3. 查询键生成哈希值
//4. 根据哈希值在索引中查找数据集合
//5. 遍历数据集合逐个与条件比对
//6. 获得结果
func queryData(name string, age int) {
	// 根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}
	// 计算查询键哈希值并查询，获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]
	// 遍历结果集合
	for _, result := range resultList {
		// 与查询结果比对，确认找到打印结果
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}
	// 没有查询到， 打印结果
	fmt.Println("no found")
}

*/

// 二、利用map特性的多键性索引及查询
// 使用结构体进行多键索引和查询比传统的写法更为简单
// 最主要的区别时无须准备哈希函数及相应的字段无须做哈希合并，如下：
// 利用map特性的多键索引和查询的代码位于pan.baidu.com/s/1ORFVTOLEYYqDhRzeq0zIiQ    提取密码：hfyf
// 1）构建索引
// 查询键
type queryKey struct {
	Name string
	Age  int
}

// 创建查询键到数据的映射
var mapper = make(map[queryKey]*Profile)

// 注意这里不使用查询键的指针。同时，结果只有*Profile类型，而不是*Profile切片，宝石查到的结果唯一

// 构建查询索引
func buildIndex(list []*Profile) {
	// 遍历所有数据
	for _, profile := range list {
		// 构建查询键
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		//保存查询键
		mapper[key] = profile
	}
}

// 2) 查询逻辑
// 根据条件查询数据
func queryData(name string, age int) {
	//根据查询条件构建查询键
	key := queryKey{name, age}
	// 根据键值查询数据
	result, ok := mapper[key]
	// 找到数据打印出来
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}

//总结：
//基于哈希值的多键索引查询和利用map特性的多键索引查询的代码复杂程度显而易见
//聪明的程序员都会利用go语言的特性进行快速的多键索引查询
//其实，利用map特性的例子中的map类型即便修改为下面的格式，也一样可以获得同样的记过：
//var mapper = make(map[interface{}]*Profile)
//代码量大大减少的关键是：Go语言的底层会为map的键自动构建哈希值，
//能够狗狗缉拿哈希值的类型必须是非动态类型，非指针，函数，闭包
//1.非动态类型：可用数组，不能用切片
//2.非指针：每个指针数值都不同，失去哈希意义
//3. 函数，闭包不能作为map的键
