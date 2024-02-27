package mock_main

import (
	"errors"
	"fmt"
	"practice/test/db/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	//NOTE: 創建了一個新的 gomock 控制器。這個控制器將用於創建模擬對象並管理對這些模擬對象的期望設置
	ctrl := gomock.NewController(t)

	//NOTE: checks to see if all the methods(就是mock的interface裡面的方法) that were expected to be called were called
	//NOTE:但現在已經改版成NewController裡面傳入 *testing.T就不用call這個了
	// defer ctrl.Finish() //斷言DB.Get() 方法是否被調用

	m := mock.NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(6541210, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Tesla")).Return(3000, nil)
	// o2 := m.EXPECT().Get(gomock.Any()).Return(630, nil)
	// o3 := m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	// m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))
	//NOTE: InOrder表示這兩次 Get 方法呼叫的順序
	//NOTE: 在這邊Get首先被呼叫兩次，其參數分別為 "Tom" 和 "Tesla"，然後會檢查這兩次呼叫的順序是否正確
	gomock.InOrder(o1, o2)
	tom_tes := GetFromDB(m, "Tom")
	fmt.Println("print tom_tes----", tom_tes) //-1(因為o1有mock回傳錯誤，所以val一定是-1)

	if tes_res := GetFromDB(m, "Tesla"); tes_res != -1 {
		fmt.Println("print tes_res----------", tes_res) //3000

		//Fatal會直接停止執行，Error不會
		t.Fatalf("expected -1, but got %v", tes_res)
	}

}
