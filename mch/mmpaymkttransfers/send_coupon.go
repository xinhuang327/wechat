// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mmpaymkttransfers

import (
	"github.com/xinhuang327/wechat/mch"
)

// 发放代金券.
func SendCoupon(proxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return proxy.PostXML("https://api.mch.weixin.qq.com/mmpaymkttransfers/send_coupon", req)
}
