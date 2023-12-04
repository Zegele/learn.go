package main

// gorm演示连接数据库
import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn.go/chapter10/02.orm/types"
	"log"
)

func main() {
	conn := connectDb()
	fmt.Println(conn)

	//查找
	//ormSelect(conn, "yy") //精确查找

	//ormSelectWithInaccurateQuery(conn)
	//ormSelectWithInaccurateQueryHack(conn) // 查询结果不会出现泄露

	//增加
	//if err := creatNewPerson(conn); err != nil {
	//	fmt.Println(err)
	//}

	//更新
	//全覆盖更新
	//undateExistingPerson(conn)
	// 更新几个值
	undateExistingPersonSelectFields(conn)

	//删除
	//deletePerson(conn)

}
func connectDb() *gorm.DB { //连接数据库 返回一个数据库
	conn, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

// 查找
// 精确查找
func ormSelect(conn *gorm.DB, name string) {
	output := make([]*types.PersonalInformation, 0)
	resp := conn.Where(&types.PersonalInformation{Name: name}).Find(&output) // 只能精确查找
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}
	fmt.Printf("查询结果：%+v\n", output) //输出的是地址，看不懂
	data, _ := json.Marshal(output)

	fmt.Printf("查询结果：%+v\n", string(data))
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) {
	output := make([]*types.PersonalInformation, 0)
	resp := conn.Where("tall > ?", 1.71).Find(&output) // 这个大于1.71 怎么查询结果包含1.71呢？
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
}
func ormSelectWithInaccurateQueryHack(conn *gorm.DB) {
	output := make([]*types.PersonalInformation, 0)
	resp := conn.Where("name = ? and sex = ?", "222' -- ", "女").Find(&output) // 这个大于1.71 怎么查询结果包含1.71呢？
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
}

// 增加
func creatNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&types.PersonalInformation{ // 注意这里有返回错误，但是被隐藏了。一定要有东西接住，查看是否有错。
		//有了types.Table函数，这里没有ID这项，也可以直接操作。和直接操作go的结构体是一样的。不需要sql语句了。
		Name:   "222",
		Sex:    "男",
		Tall:   1.80,
		Weight: 65.0,
		Age:    25,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建xxx人时失败：", err)
		return err
	}
	fmt.Println("创建xxx人成功！")
	return nil
}

// 更新
// 全覆盖模式更新
func undateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&types.PersonalInformation{
		//没有 primary key 就不知道你要更新的数据是给哪条数据做更新的。
		ID:     6,
		Name:   "y",
		Sex:    "男",
		Tall:   2,
		Weight: 71,
		Age:    30,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新***人时失败：", err)
		return err
	}
	fmt.Println("更新***成功")
	return nil
}

// 只更新某个字段
func undateExistingPersonSelectFields(conn *gorm.DB) error {
	p := &types.PersonalInformation{
		ID:   7,
		Name: "xiao",
		Sex:  "男",
		Tall: 2.0, // 结构体中的数据是可以 缺的。
		//Weight: 65,
		//	Age:    30,
	}
	resp := conn.Model(p).Select("name", "tall").Updates(p) // 表示 把p结构体中的，name和tall，更新回p
	//Model(p)指定更新p，Select("name", "tall")是选择p的name和tall字段，Updates(p)是更新回p
	if err := resp.Error; err != nil {
		fmt.Println("更新***人时失败：", err)
		return err
	}
	fmt.Println("更新***成功")
	return nil
}

// 删除
func deletePerson(conn *gorm.DB) error {
	p := types.PersonalInformation{ID: 6} // 通过primary key 删除
	resp := conn.Delete(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除**人时失败：", err)
		return err
	}
	fmt.Println("删除**成功。")
	return nil
}
