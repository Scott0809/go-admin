package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

func (h *Handler) Plugins(ctx *context.Context) {

	var (
		size = types.Size(12, 6, 4)
		rows = template.HTML("")
		list = plugins.Get()
	)

	for i := 0; i < len(list); i += 3 {

		updated, _ := list[i].CheckUpdate()
		skip, _ := list[i].GetInstallationPage()

		box1 := aBox().
			SetBody(h.pluginBox(list[i].GetInfo(), list[i].IsInstalled(), updated, skip, list[i].Name())).
			GetContent()
		col1 := aCol().SetSize(size).SetContent(box1).GetContent()
		box2, col2, box3, col3 := template.HTML(""), template.HTML(""), template.HTML(""), template.HTML("")
		if i+1 < len(list) {
			updated, _ := list[i+1].CheckUpdate()
			skip, _ := list[i+1].GetInstallationPage()
			box2 = aBox().
				SetBody(h.pluginBox(list[i+1].GetInfo(), list[i+1].IsInstalled(), updated, skip, list[i+1].Name())).
				GetContent()
			col2 = aCol().SetSize(size).SetContent(box2).GetContent()
			if i+2 < len(list) {
				updated, _ := list[i+2].CheckUpdate()
				skip, _ := list[i+2].GetInstallationPage()
				box3 = aBox().
					SetBody(h.pluginBox(list[i+2].GetInfo(), list[i+2].IsInstalled(), updated, skip, list[i+2].Name())).
					GetContent()
				col3 = aCol().SetSize(size).SetContent(box3).GetContent()
			}
		}
		rows += aRow().SetContent(col1 + col2 + col3).GetContent()
	}

	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content:     rows,
		Description: language.GetFromHtml("plugins"),
		Title:       language.GetFromHtml("plugins"),
	})
}

func (h *Handler) pluginBox(info plugins.Info, install, upgrade, skip bool, name string) template.HTML {
	col1 := template.HTML(`<div class="plugin-item-img">`) + aImage().
		SetSrc("http://localhost:9033/admin/assets/dist/img/avatar04.png").
		SetHeight("110px").
		SetWidth("110px").
		GetContent() + template.HTML(`</div>`)
	uid := utils.Uuid(18)
	footer := template.HTML(`<button data-toggle="modal" data-target="#`+uid+`"  class="btn btn-primary plugin-info">`) +
		plugWordHTML("info") + template.HTML("</button>")
	popupModel := template2.Default().Popup().SetID(uid).
		SetTitle(plugWordHTML("plugin detail")).
		SetWidth("730px").
		SetHeight("400px").
		SetBody(template.HTML(`
<div class="plugin-detail">
<div class="plugin-detail-head">
<div class="plugin-detail-head-logo">
	<img src="http://localhost:9033/admin/assets/dist/img/avatar04.png" width="110px" height="110px">
</div>
<div class="plugin-detail-head-title">
	<div class="plugin-detail-title">` + info.Title + `</div>
	<div class="plugin-detail-provider">` +
			fmt.Sprintf(plugWord("provided by %s"), info.Author) + `</div>
</div>
</div>
<div class="plugin-detail-info">
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("introduction") + `</div>
		<div class="plugin-detail-info-item-content">` + info.Description + `</div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("website") + `</div>
		<div class="plugin-detail-info-item-content">` + info.Website + `</div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("version") + `</div>
		<div class="plugin-detail-info-item-content">` + info.Version + `</div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("created at") + `</div>
		<div class="plugin-detail-info-item-content">` + info.CreatedAt.Format("2006-01-02") + `</div>
	</div>
	<div class="plugin-detail-info-item">
		<div class="plugin-detail-info-item-head">` + plugWord("updated at") + `</div>
		<div class="plugin-detail-info-item-content">` + info.UpdatedAt.Format("2006-01-02") + `</div>
	</div>
</div>
</div>`))
	if install {
		if upgrade {
			footer += template.HTML(`<button class="btn btn-primary installation">`) + plugWordHTML("upgrade") + template.HTML("</button>")
		}
	} else {
		if skip {
			footer += template.HTML(`<button class="btn btn-primary installation">`) + plugWordHTML("install") + template.HTML("</button>")
		} else {
			btn := template2.HTML(`<button class="btn btn-primary" onclick="plugin`+name+`Install()">`) + plugWordHTML("install") + template.HTML("</button>")
			footer += template2.HTML(`<a href="`+h.config.Prefix()+`/info/plugin_`+name+`/new"><button class="btn btn-primary installation" >`) +
				plugWordHTML("install") + template.HTML("</button></a>")
			popupModel = popupModel.SetFooterHTML(btn)
		}
	}
	popup := popupModel.GetContent()
	col2 := template.HTML(`<div class="plugin-item-content">`) +
		template.HTML(`<div class="plugin-item-content-title">`+info.Title+`</div>
	<div class="plugin-item-content-description">`+info.Description+"</div>") +
		footer + template.HTML(`</div>`)

	return `<div class="clear:both;">` + col1 + col2 + `</div>` + pluginCSS + popup + template2.HTML(
		`<script>function plugin`+name+`Install(){location.href="`+h.config.Prefix()+`/info/plugin_`+name) +
		`/new"}</script>`
}

func plugWord(word string) string {
	return language.GetWithScope(word, "plugin")
}

func plugWordHTML(word template.HTML) template.HTML {
	return language.GetFromHtml(word, "plugin")
}

var pluginCSS = template.HTML(`
<style>
	.plugin-item-content {
		margin-left: 15px;
	}
	.plugin-item-content-title {
		font-size: 15px;
		margin-bottom: 10px;
		font-weight: bold;
	}
	.plugin-item-content {
		position: absolute;
		margin-left: 121px;
		padding-right: 10px;
		top: 7px;
	}
	.plugin-item-content-description {
		overflow: hidden;
		text-overflow: ellipsis;
		word-break: break-all;
		display: -webkit-box;
		font-size: 15px;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		height: 42px;
	}
	.installation {
		float: right;
		margin-top: 10px;
	}
	.plugin-info {
		float: right;
		margin-top: 10px;
		margin-left: 10px;
	}
	.plugin-detail {
		padding: 10px;
	}
	.plugin-detail-head {
		clear: both;
		height: 112px;
    	margin-bottom: 33px;
	}
	.plugin-detail-title {
		font-size: 30px;
	}
	.plugin-detail-provider {
    	font-size: 15px;
    	margin-top: 4px;
	}
	.plugin-detail-head-logo {
		width: 120px;
		float: left;
	}
	.plugin-detail-head-title {
		float: left;
		margin-left: 10px;
	}
	.plugin-detail-info-item {
		clear: both;
		height: 15px;
		margin-bottom: 17px;
	}
	.plugin-detail-info-item-head {
		width: 80px;
		float: left;
		font-weight: bold;
	}
	.plugin-detail-info-item-content {
		float: left;
		margin-left: 10px;
	}
</style>`)
