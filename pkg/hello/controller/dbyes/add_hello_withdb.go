package dbyes

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
)

type BodyAddHello struct {
	Code string `json:"code" validate:"required" example:"H001"`
	Name string `json:"name" validate:"required" example:"Hellow"`
}

type Hello struct {
	Id   int64  `json:"id" example:"10"`
	Code string `json:"code" example:"H001"`
	Name string `json:"name" example:"Hellow"`
}

func AddHello(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyAddHello{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := Hello{}

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" INSERT INTO ", tables.S_HELLO, "( code, name ) ").
			Add(" VALUES ( :code, :name )").
			Add(" RETURNING id, code, name ").
			SetParam("code", body.Code).
			SetParam("name", body.Name), &out)
		if err != nil {
			return err
		}

		return out
	})
}
