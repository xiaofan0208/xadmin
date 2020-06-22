package controllers

//BaseAdminController 后台控制器
type BaseAdminController struct {
	BaseController
}

// Prepare  prepare
func (ctl *BaseAdminController) Prepare() {
	ctl.BaseController.Prepare()

}

// SetTpl 设置布局文件
func (ctl *BaseAdminController) SetTpl(tpl ...string) {
	baselayout := "admin/base/layout.html"
	switch len(tpl) {
	case 1:
		ctl.TplName = tpl[0]
		ctl.Layout = baselayout
		break
	case 2:
		ctl.TplName = tpl[0]
		ctl.Layout = tpl[1]
		break
	}
}
