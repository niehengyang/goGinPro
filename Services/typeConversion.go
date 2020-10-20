package Services

/**
  变量类型转换
*/

import (
    "fmt"
    "reflect"
    "strconv"
)

// func main() {
//     v := "hello world"
//     fmt.Println(typeofFmt(v))
//     fmt.Println(typeofReflect(v))

//     str := "1245"
//     fmt.Println(stringToInt(str))
//     fmt.Println(stringToInt64(str))

//     a := 3545;
//     var b int64;
//     b = 98;
//     fmt.Println(int64ToString(b))
//     fmt.Println(intToString(a))

//     var f float64
//     f = 3.45
//     fmt.Println(floatToString(f))

//     str = "1.24"
//     fmt.Println(stringToFloat(str))

//     //int到int64
//     fmt.Println(int64(1234))

//     //int 转化为 float
//     score := 100
//     fmt.Println(float64(score))
// }

// 利用fmt
func TypeofFmt(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

// 利用reflect
func TypeofReflect(v interface{}) string {
    return reflect.TypeOf(v).String()
}

// string转int
func StringToInt(a string) int {
    d, _ := strconv.Atoi(a)
    return d
}

//Atoi是ParseInt(s, 10, 0)的简写。

// string转int64
func StringToInt64(a string) int64 {
    d, _ := strconv.ParseInt(a, 10, 64)
    return d
}

/**
func ParseInt(s string, base int, bitSize int) (i int64, err error)
返回字符串表示的整数值，接受正负号。
base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
*/

// int转string
func IntToString(a int) string {
    str := strconv.Itoa(a)
    return str
}

// int64转string
func Int64ToString(a int64) string {
    str := strconv.FormatInt(a, 10)
    return str
}

/**
func FormatInt(i int64, base int) string
返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
*/

// float转string
func FloatToString(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 32)
}

/**
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
*/

// string转float
func StringToFloat(s string) float64 {
    f, _ := strconv.ParseFloat(s, 64)
    return f
}

/**
func ParseFloat(s string, bitSize int) (f float64, err error)
解析一个表示浮点数的字符串并返回其值。

如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
*/
