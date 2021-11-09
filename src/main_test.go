package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	myMock "testTrain2/mocks"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	result := Summer(5 ,6)
	if result != 11 {
		t.Fatal()
	}

}

var testCases = []struct {
	x  int
	y int
	out int
}{
	{1,5 ,6},
	{-5,5,10},
	{6,12,18},
	{-9,-8,17},
}

func TestSolution2(t *testing.T) {
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("inputs %d , %d",tc.x , tc.y), func(t *testing.T) {
			t.Parallel() // must be in first line
			rnd := rand.Intn(10)
			fmt.Println("****" ,rnd , "***")
			time.Sleep(time.Duration(10 -rnd) * time.Second)
			result := Summer(tc.x ,tc.y)
			//if result != tc.out	 {
			//	t.Fatal()
			//}
			assert.Equal(t, tc.out , result)
		})
	}
}

func TestFlagParser(t *testing.T) {
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("inputs %d , %d",tt.x , tt.y) ,func(t *testing.T) {
			time.Sleep(1 * time.Second)
			s := Summer(tt.x,tt.y)
			if s != tt.out {
				t.Errorf("be fana raftim")
			}
		})
	}
}

var testSub = []struct {
	name  string
	income int
	out bool
}{
	{"hasan",5 ,true},
	{"hasan",5,true},
	{"ali",12,false},
	{"ali",-8,false},
}


func TestMyThing(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//mockObj := mym.NewMockMyInterface(mockCtrl)
	mockObj := myMock.NewMockmyDB(mockCtrl)
	mockObj.EXPECT().SubMoney("hasan",5442 ).Return(true).AnyTimes() //stubs
	mockObj.EXPECT().AddMoney(gomock.Any(),gomock.Any()).DoAndReturn(func(name string,j int) interface{} {
		if name == "hasan" {
			return true
		}else{
			return false
		}
	}).AnyTimes() // if u dont set any time u can , its work just one

	fmt.Println("*****",mockObj.AddMoney("hasan" , 5485),"******")


	for _, tc := range testSub {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("inputs %s , %d",tc.name , tc.income), func(t *testing.T) {
			t.Parallel() // must be in first line
			result := mockObj.AddMoney(tc.name ,tc.income)
			//if result != tc.out	 {
			//	t.Fatal()
			//}
			assert.Equal(t, tc.out , result)
		})
	}



	//assert.Equal(t, true,b)
	// pass mockObj to a real object and play with it.
}


func BenchmarkMyDB(b *testing.B) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)
	}
	assert.Equal(b, nil,nil)
}

func BenchmarkFib(b *testing.B) {
	b.N = 5 //number of repeat

	fmt.Println(Fib(b.N))

}



