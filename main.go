package main

import (
	"encoding/json"
	"fmt"
	"strings"

	jwt "github.com/cristalhq/jwt/v3"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	header    *walk.TextEdit
	payload   *walk.TextEdit
	signature *walk.TextEdit
}

func main() {
	mw := &MyMainWindow{}
	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "JWT Token View",
		MinSize:  Size{640, 320},
		Size:     Size{640, 320},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LineEdit{
						AssignTo: &mw.searchBox,
						Text:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
						Row:      5,
					},
					PushButton{
						Text:      "查看",
						OnClicked: mw.clicked,
					},
				},
			},
			GroupBox{
				Layout: VBox{},
				Children: []Widget{
					TextEdit{
						AssignTo: &mw.header,
						MinSize:  Size{320, 80},
						Row:      2,
					},
					TextEdit{
						AssignTo: &mw.payload,
						MinSize:  Size{320, 320},
						Row:      10,
					},
					TextEdit{
						AssignTo: &mw.signature,
					},
				},
			},
		},
	}.Run()
}

func (mw *MyMainWindow) clicked() {
	input := mw.searchBox.Text()
	token, err := jwt.ParseString(input)
	if err != nil {
		return
	}
	header, err := json.MarshalIndent(token.Header(), "", "    ")
	if err != nil {
		return
	}
	mw.header.SetText(strings.Replace(string(header), "\n", "\r\n", -1))
	fmt.Println(string(header))
	pack := map[string]interface{}{}
	err = json.Unmarshal(token.RawClaims(), &pack)
	if err != nil {
		return
	}
	payload, _ := json.MarshalIndent(pack, "", "    ")
	mw.payload.SetText(strings.Replace(string(payload), "\n", "\r\n", -1))
	mw.signature.SetText("未开发")
}
