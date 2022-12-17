package tools

import "os"

// RemoveLogo
//
//	@Description: 删除头像资源
//	@param logoUrl 删除的头像地址
//	@return error  DD
func RemoveLogo(logoUrl string) error {
	err := os.Remove("./" + logoUrl)
	return err
}
