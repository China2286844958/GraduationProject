package tools

/**
@Title 哈希256密码明文库
@Author 薛智敏
@CreateTime 2022年6月21日19:32:08
*/
const (
	sha256_1234   = "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4"
	sha256_123456 = "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"
)

var shaEnCodeSha256ByDeCode = []string{}
var sha256EncodeMap = make(map[string]string)

//根据明文，获取加密库中，加密处理过的密文，大大节省加密的时间，直接获取明文对应的密文

func GetSha256EnCodeByDeCode(DeCode string) (EnCode string) {
	sha256EnCodeInit()
	return sha256EncodeMap[DeCode]
}

//加密密文放库中，初始化

func sha256EnCodeInit() {
	sha256EncodeMap["1234"] = sha256_1234
	sha256EncodeMap["123456"] = sha256_123456
}
