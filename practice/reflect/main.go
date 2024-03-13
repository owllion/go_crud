package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type Savings interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}
type Account struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	TopSecret string  `json:"topSecret"`
}

var a int
var b string
var m (map[string]string)
var account = Account{"1", "Ruel", 8335, "nothing"}

// 幫Account家個方法
func (a *Account) String() string {
	//NOTE:補充1
	//String()這種方法，在你取用enum值的時候就會直接幫你取道該ENUM變-數對應的(你設定的)值
	//NOTE:補充2
	//承上，其實String()貌似有特殊意義欸，目前發現只要和fmt.印出系列有關的行為，印出的結果都會是我這邊回傳的!!!!像是我原本取 account的值，印出的是" {1 Ruel 8335 nothing}"，後來改成取&指針的，印出結果就是下面這行，然後methodByName.Call的結果也是依樣!!!

	return fmt.Sprintf("ID: %s, Name: %s, Balance: %.2f, Remark: %s", a.Id, a.Name, a.Balance, a.TopSecret)

}

func (a *Account) SG(singerNames ...string) string {
	return fmt.Sprintf("SG's singers consist of %s,%s,%s and %s", singerNames[0], singerNames[1], singerNames[2], singerNames[3])
}

func main() {
	// printBasicType()
	// printFieldInfo()
	printStructFieldValue()
	// getStructValByElem()
	// deepEqual()
	// printByVFormat()
	// printEnum()
	// typeAssert()
	// printTypeFromFormat()
	// dynamicTypeDeclare()
	// printConst()
	go func() {
		fmt.Println("print Hello")
	}()
	fmt.Println("numOfGoroutine", runtime.NumGoroutine()) //2
}

type Test int

func printConst() {
	numOfCPU := runtime.NumCPU()
	maxProcs := runtime.GOMAXPROCS(20)
	fmt.Println("numOfCPU:", numOfCPU) //20
	fmt.Println("maxProcs:", maxProcs) //20
	//NOTE:補充
	//CPU核心分成虛擬&實體，以我的12核20緒來說，12是實體數量，20緒是虛擬的核心(又稱作"邏輯"處理器)，
	//那golang上面兩個方法都是拿虛擬的20，為何?不是明擺著是12嗎?
	//這是因為比較好的CPU會有超頻功能(具體不清楚，已涉及到硬體)，總之這個功能就是可以讓你的虛擬核心被當成實體的拿去做使用，以我的電腦來說，就是相當於我有20科核心，一次最多可以開20條thread，而不只12條。

	const (
		i Test = iota
		j
		k
		// // i = 7
		// j
		// k
	)
	fmt.Println("print i j k", i, j, k) //0 1 2
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
	for i := 0; i < t.NumField(); i++ {
		entireTag := t.Field(i).Tag
		fmt.Println("print entireTagg----", entireTag) //json:"id"

		specificTagVal := t.Field(i).Tag.Get("json")
		fmt.Println("print specificTagVal----", specificTagVal) //id

		name := t.Field(i).Name
		s, _ := t.FieldByName(name) //不知為啥突然用reflect.FieldByName會報錯...
		fmt.Println("print json tag value*---", s.Tag.Get("json"))
	}

}

// 要取值要用ValueOf
func printStructFieldValue() {
	//NOTE: 這邊可以加上pointer，但只限於單純印出valueOfAccount，如果想要用這結果取到每個欄位結果，會報錯(也就是說，取記憶體位置的value，就不能用FieldByName這種方式去取直)
	ptrVal := reflect.ValueOf(&account)
	instanceval := reflect.ValueOf(account)
	fmt.Println("print ptrVal---", ptrVal)
	fmt.Println("print instanceval---", instanceval)

	resOfCallString := ptrVal.MethodByName("String").Call([]reflect.Value{})
	fmt.Println("print result of calling ptr method----", resOfCallString)

	resOfCallSG := ptrVal.MethodByName("SG").Call(
		[]reflect.Value{
			reflect.ValueOf("Megan Thee Stallion"),
			reflect.ValueOf("Ozuna"),
			reflect.ValueOf("LISA"),
			reflect.ValueOf("DJ Snake"),
		})
	fmt.Println("print res of calling SG---", resOfCallSG)

	fmt.Println("account id---", instanceval.FieldByName("Id"))
	fmt.Println("account name---", instanceval.FieldByName("Name"))
	fmt.Println("account balance---", instanceval.FieldByName("Balance"))
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
