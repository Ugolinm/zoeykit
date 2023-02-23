package easykit

import (
	"fmt"

	"github.com/google/uuid"
)

func RandomUUID() {

	uuid := uuid.New()
	key := uuid.String()
	fmt.Println("key", key)
	// return key
}

// 待加密字符串
// s := "4738c1d9-04ce-46a5-a49f-7cc2b9f061e2"

// // 进行md5加密，因为Sum函数接受的是字节数组，因此需要注意类型转换
// srcCode := md5.Sum([]byte(s))

// // md5.Sum函数加密后返回的是字节数组，需要转换成16进制形式
// code := fmt.Sprintf("%x", srcCode)

// fmt.Println(string(code))
