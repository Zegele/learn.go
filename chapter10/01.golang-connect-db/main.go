package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //如果go mod tidy后是红的 go mod vender一下应该就好了
	"learn.go/pkg/apis"
	"log"
	//需要的是该包中init 注册mysql驱动器 数据库的驱动
	//func init() {
	//	sql.Register("mysql", &MySQLDriver{})
	//}
)

func main() {
	learnDB, err := getDbConnection()
	defer learnDB.Close()

	err = testDbConnection(err, learnDB)

	//queryAllData(err, learnDB)
	//
	//if err = insertData(learnDB); err != nil {
	//	log.Fatal(err)
	//}
	//
	//queryAllData(err, learnDB)
	queryAllDataHack(err, learnDB) // 查询结果泄露
}
func insertData(learnDB *sql.DB) error {
	_, err := learnDB.Exec(fmt.Sprintf("insert into personal_information(name, sex, tall, weight,age) values('%s','%s',%f,%f,%d)",
		"??",
		"女",
		1.75,
		55.0,
		18,
	))
	if err != nil {
		fmt.Println("新增数据失败：%v", err)
		return err
	}
	return nil
}

func queryAllData(err error, learnDB *sql.DB) {
	rows, err := learnDB.Query("select age,name from personal_information") //sql语句就行
	//rows, err := learnDB.Query("select * from personal_information")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*rows)
	list := &apis.PersonalInfomationList{}
	for rows.Next() { //从这一行，循环下一行。
		var name string
		var age int
		if err := rows.Scan(&age, &name); err != nil { //扫描的元素顺序必须和select元素的顺序保持一致。否则把select出的数据，就扫描进不匹配的变量中了。
			fmt.Println(age, name)
			// 更换列的顺序，因为数据类型不匹配导致失败。 警告：如果数据类型兼容，会引发更大的灾难。
			log.Printf("扫描数据失败，跳过该行：%v", err)
		}
		fmt.Printf("name: %s\tage:%d\n", name, age) // \t 制表符 就是一个tab
		list.Items = append(list.Items, &apis.PersonalInfomation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func queryAllDataHack(err error, learnDB *sql.DB) {
	_sql := fmt.Sprintf(`select name, sex from personal_information where name = '%s' and sex = '%s'`, "222' -- ", "女") //但是查询结果把数据泄露出来了。
	//"222", "女" 本来是不存在该数据的。
	rows, err := learnDB.Query(_sql) //sql语句就行
	//rows, err := learnDB.Query("select * from personal_information")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(*rows)
	list := &apis.PersonalInfomationList{}
	for rows.Next() { //从这一行，循环下一行。
		var name string
		var sex string
		if err := rows.Scan(&name, &sex); err != nil { //扫描的元素顺序必须和select元素的顺序保持一致。否则把select出的数据，就扫描进不匹配的变量中了。
			fmt.Println(name, sex)
			// 更换列的顺序，因为数据类型不匹配导致失败。 警告：如果数据类型兼容，会引发更大的灾难。
			log.Printf("扫描数据失败，跳过该行：%v", err)
		}
		fmt.Printf("name: %s\tsex:%s\n", name, sex) // \t 制表符 就是一个tab
		list.Items = append(list.Items, &apis.PersonalInfomation{
			Name: name,
			Sex:  sex,
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func testDbConnection(err error, learnDB *sql.DB) error {
	if err = learnDB.Ping(); nil != err {
		fmt.Println("DB 测试失败：", err)
	}
	fmt.Println("数据库连接成功，可以执行命令。")
	return err
}

func getDbConnection() (*sql.DB, error) {
	learnDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb") //最后一个是数据库名
	//用户名 ： 密码 @tcp(本机地址：端口)/数据库名
	if err != nil {
		log.Fatalf("链接数据库失败：%v", err)
	}

	return learnDB, err
}
