// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

import (
	"net/http"

	"github.com/xinhuang327/wechat/corp"
)

type Client struct {
	*corp.Client
}

// 兼容保留, 建議實際項目全局維護一個 *corp.Client
func NewClient(srv corp.AccessTokenServer, clt *http.Client) Client {
	return Client{
		Client: corp.NewClient(srv, clt),
	}
}
