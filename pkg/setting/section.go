package setting

type AppSettingS struct {
	RunMode string
}

type EmailSettingS struct {
	// ServerHost 邮箱服务器地址，如腾讯邮箱为smtp.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式）
	FromPasswd string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
