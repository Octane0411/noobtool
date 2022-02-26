package ipgw

import (
	"github.com/neucn/neugo"
	"log"
	"net/http"
	"net/url"
	"noobtool/global"
	"noobtool/pkg/utils"
)

type IpgwHandler struct {
	client *http.Client
}

func (h *IpgwHandler) Login() {
	_, err := h.login()
	if err != nil {
		log.Printf("login err: %v", err)
	}
}

func (h *IpgwHandler) login() (string, error) {
	if err := h.NEUAuth(); err != nil {
		return "", err
	}
	return h.requestLoginApi()
}

func NewIpgwHandler() *IpgwHandler {
	return &IpgwHandler{
		neugo.NewSession(),
	}
}
func (h *IpgwHandler) NEUAuth() error {
	return neugo.Use(h.client).WithAuth(global.IpgwSetting.Username, global.IpgwSetting.Password).Login(neugo.CAS)
}

func (h *IpgwHandler) loginCookie(cookie string) (string, error) {
	h.client.Jar.SetCookies(&url.URL{
		Scheme: "https",
		Host:   "ipgw.neu.edu.cn",
	}, []*http.Cookie{{
		Name:   "session_for%3Asrun_cas_php",
		Value:  cookie,
		Domain: "ipgw.neu.edu.cn",
	}})
	return h.requestLoginApi()
}

func (h *IpgwHandler) requestLoginApi() (string, error) {
	// 获取当前网络下对应网关url的query参数
	resp, err := h.client.Get("https://ipgw.neu.edu.cn/")
	if err != nil {
		return "", err
	}
	// 统一认证拿到ticket
	resp, err = h.client.Get("https://pass.neu.edu.cn/tpass/login?service=http://ipgw.neu.edu.cn/srun_portal_sso?" + resp.Request.URL.RawQuery)
	if err != nil {
		return "", err
	}
	// 使用ticket调用api登录
	req, _ := http.NewRequest("GET", "https://ipgw.neu.edu.cn/v1"+resp.Request.URL.RequestURI(), nil)
	resp, err = h.client.Do(req)
	if err != nil {
		return "", err
	}
	return utils.ReadBody(resp), nil
}
