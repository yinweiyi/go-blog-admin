package forms

type Comment struct {
	Type      string `form:"type" json:"type"`
	ID        int    `form:"id" json:"id"`
	Email     string `form:"email" json:"email" validate:"c_email"`
	ParentId  int    `form:"parent_id" json:"parent_id"`
	Nickname  string `form:"nickname" json:"nickname"`
	Content   string `form:"content" validate:"required" json:"content"`
	Captcha   string `form:"captcha" validate:"required" json:"captcha"`
	CaptchaId string `form:"captcha_id" validate:"required" json:"captcha_id"`
	IP        string `json:"ip"`
}
