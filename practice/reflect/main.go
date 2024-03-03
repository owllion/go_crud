package main

import (
	"fmt"
	"reflect"
)

type Savings interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}
type Account struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	topSecret string  `json:"topSecret"`
}

var a int
var b string
var m (map[string]string)
var account = Account{"1", "Ruel", 8335, "nothing"}

func main() {
	// printBasicType()
	printFieldInfo()
	// printStructFieldValue()
	// getStructValByElem()
	// deepEqual()
	// printByVFormat()
	// printEnum()
	// typeAssert()
	// testEnum()
	// printTypeFromFormat()
	// dynamicTypeDeclare()
}
func dynamicTypeDeclare() {
	var x interface{}
	x = 5
	fmt.Printf("print x's type, %T\n", x) //int
	x = "RR"
	fmt.Printf("print x's type-2, %T\n", x) //string
}
func printTypeFromFormat() {
	var a, b, c = 3, 4, "Foo"
	fmt.Printf("print a's type: %T\n", a) //int
	fmt.Printf("print b's type: %T\n", b) //int
	fmt.Printf("print c's type: %T\n", c) //string
}
func testEnum() {
	const (
		i = 7
		j
		k
	)
	fmt.Println(i, j, k)
}

func printBasicType() {
	typeOf_A := reflect.TypeOf(a)
	typeOf_B := reflect.TypeOf(b)
	typeOf_M := reflect.TypeOf(m)
	t := reflect.TypeOf(account)

	fmt.Println("typeff_A:", typeOf_A)
	fmt.Println("typeOf_B:", typeOf_B)
	fmt.Println("typeOf_M:", typeOf_M)
	fmt.Println("t:", t)
}

// 取struct的type、name、tag(下面兩種寫法都依樣，但第2種比較簡潔)
func printFieldInfo() {
	t := reflect.TypeOf(account)
	// paths := t.PkgPath()
	// fmt.Println("@@@@@@@@@paths@@@@@@", paths)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		path := t.Field(i).PkgPath
		fmt.Println("@@@print path@@@@@", path)
		fmt.Println("print Field Tag**---", tag) //json:"balance"
		name := t.Field(i).Name
		fmt.Println("priint name----", name)
		//NOTE: FieldName
		s, _ := t.FieldByName("Balance")
		fmt.Println("print type----", s.Type)
		fmt.Println("print name by FieldByName----", s.Name)
		fmt.Println("print entire tag----------", s.Tag) //json:"balance"
		//NOTE: Get tag value
		fmt.Println("print tag value----------", s.Tag.Get("json")) //balance
		//NOTE: lowercase member's info can also be accessed.
		private, _ := t.FieldByName("topSecret")
		fmt.Println("print private member's name----------", private.Name)
		fmt.Println("print private member's tag----------", private.Tag.Get("json"))

	}

}

// 要取值要用ValueOf
func printStructFieldValue() {
	valueOfAccount := reflect.ValueOf(account)
	fmt.Println("print all---", valueOfAccount)
	fmt.Println("account id---", valueOfAccount.FieldByName("Id"))
	fmt.Println("account name---", valueOfAccount.FieldByName("Name"))
	fmt.Println("account balance---", valueOfAccount.FieldByName("Balance"))
}

func getStructValByElem() {
	account := &Account{"2", "Justin", 1245, "nothing"}
	var a = 8
	basicType := reflect.TypeOf(account).Kind()
	basicAType := reflect.TypeOf(a).Kind()

	fmt.Println("print account's basicType(kind)----", basicType) //ptr
	//NOTE: 看來只有pointer類型的會有差，其餘基本類別加不加Kind都一樣的樣子
	fmt.Println("print account's basicAType(kind)----", basicAType)  //int
	fmt.Println("print account's basicAType----", reflect.TypeOf(a)) //int

	if basicType.String() == "ptr" {
		fmt.Println("印出basicType---", basicType.String())
	}
	fmt.Println("print account's type----", reflect.TypeOf(account))

	t := reflect.ValueOf(account).Elem()                          //如果你需要從pointer讀取值就要用Elem，不然一般實體就是ValueOf()即可
	fmt.Println("print elem value----", t.FieldByName("Balance")) //1245
}

func deepEqual() {

	a1 := Account{"1", "Ruel", 123, "nothing"}
	// a2 := Account{"1", "Ruel", 132} //Not Equal
	a2 := Account{"1", "Ruel", 123, "nothing"} // Equal

	isEqual := reflect.DeepEqual(a1, a2)

	if isEqual {
		fmt.Println("Is Equal.")
		return
	}
	fmt.Println("Is Not Equal.")
}

func printByVFormat() {
	aVar := 5
	account := Account{"5", "Juan", 123, "nothing"}

	fmt.Printf("Use %%v for account: %v\n", account)
	fmt.Printf("Use %%v for aVar: %v\n", aVar)
	fmt.Printf("Use %%+v for account: %+v\n", account)
	fmt.Printf("Use %%+v for aVar: %+v\n", aVar)
	fmt.Printf("Use %%#v for account: %#v\n", account)
	fmt.Printf("Use %%#v for aVar: %#v\n", aVar)
}

// 定义枚举类型
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

func (d Weekday) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	default:
		return "懶得打剩下的了ㄏ"
	}
}
func printEnum() {
	// day := Sunday
	day := Wednesday
	//兩個結果一樣，應該是因為都是要轉成字串
	fmt.Printf("Today is %s", day)
	fmt.Println("Today is", day)
}

func typeAssert() {
	m_ := map[string]interface{}{
		"name": "Ruel",
		"age":  23,
	}
	if val, ok := m_["name"].(string); ok {
		fmt.Println("name val---", val)
		fmt.Println("print ok*---", ok)
	}
}
