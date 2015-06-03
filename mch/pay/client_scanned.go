// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"github.com/xinhuang327/wechat/mch"
)

// 提交被扫支付.
func MicroPay(proxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return proxy.PostXML("https://api.mch.weixin.qq.com/pay/micropay", req)
}

// 撤销支付.
//  NOTE: 请求需要双向证书.
func Reverse(proxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return proxy.PostXML("https://api.mch.weixin.qq.com/secapi/pay/reverse", req)
}
