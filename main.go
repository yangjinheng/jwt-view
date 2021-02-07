package main

import (
	"encoding/json"
	"log"
	"strings"

	jwt "github.com/cristalhq/jwt/v3"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var (
		mainWindow *walk.MainWindow
		searchBox  *walk.LineEdit
		header     *walk.TextEdit
		payload    *walk.TextEdit
		signature  *walk.TextEdit
	)

	var Font Font = Font{
		Family:    "Consolas",
		PointSize: 10,
	}

	err := MainWindow{
		AssignTo: &mainWindow,
		Title:    "JWT Token View",
		Size:     Size{Width: 700, Height: 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					LineEdit{
						AssignTo: &searchBox,
						Text:     "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJJeXhNSThiVXIwcFlUSEkzbXdUS2lXRkU3QmtjVWlyNkdwbWRtVzhIVU9RIn0.eyJleHAiOjE1ODU5MjUyMzYsImlhdCI6MTU4NTkyNTE3NiwiYXV0aF90aW1lIjoxNTg1OTI0ODMzLCJqdGkiOiIyNDE0ZTZhOC1hODgyLTQ4ZGYtOWE0My1kZWI3ZWRkMzY5ZjgiLCJpc3MiOiJodHRwOi8vMTcyLjE2LjEwMC4xMDo4MDgwL2F1dGgvcmVhbG1zL21hc3RlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIzOTg4MDZiNS1hY2ZlLTQyNjAtOTE4Mi1iZDc5NTgxMjJhNTYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJrb25nIiwic2Vzc2lvbl9zdGF0ZSI6IjZlZDAxMmZmLWU2MzAtNDBjZS04NmU3LTI0ZDNhNTBiYTI2MSIsImFjciI6IjAiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiamluaGVuZyJ9.Zo0ZXymHbCGuv_IjGe32rVzkTm5fAYNn5eVONc_Pot1S6fMd-AASn-1w4hyOYSLWrGQGzR5Okoc-TEloVvDVsax-K29_VQsoq2UcKK9His4Cnth1nWO4NK7oTvOUHM3axFr48Fo-mjmW-mqMBk4krGGEbGDq9saXlrgWgc-n-Zn_91qvi5ncSf33oy-fnOFY-_YWZxykFySEXZ1bvPlzkut7gVAPF7f4IAtuwoPy2RFHTy0QKUUWUufoLXnvu1XZwFJTt4enuCliY_mD9kV6nghdfu9tii4TJHcv1bc_7Z4FX_ywvGkgb4aUx76697Z9SlOVqMtED8JziSC1Na0-qg",
						Font:     Font,
					},
					PushButton{
						Text: "查看",
						OnClicked: func() {
							input := searchBox.Text()
							token, err := jwt.ParseString(input)
							if err != nil {
								return
							}
							data, err := json.MarshalIndent(token.Header(), "", "  ")
							if err != nil {
								return
							}
							header.SetText(strings.Replace(string(data), "\n", "\r\n", -1))
							pack := map[string]interface{}{}
							err = json.Unmarshal(token.RawClaims(), &pack)
							if err != nil {
								return
							}
							result, _ := json.MarshalIndent(pack, "", "  ")
							payload.SetText(strings.Replace(string(result), "\n", "\r\n", -1))
							signature.SetText("未开发")
						},
					},
				},
			},

			TextEdit{
				AssignTo: &header,
				Font:     Font,
				MinSize:  Size{Height: 100},
				VScroll:  true,
			},
			TextEdit{
				AssignTo: &payload,
				Font:     Font,
				MinSize:  Size{Height: 500},
				VScroll:  true,
			},
			TextEdit{
				AssignTo: &signature,
				Font:     Font,
				VScroll:  true,
			},
			// 	},
			// },
		},
	}.Create()

	if err != nil {
		log.Fatal(err)
	}

	mainWindow.Run()
}
