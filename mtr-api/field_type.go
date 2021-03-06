package main

import (
	"bytes"
	"database/sql"
	"github.com/GeoNet/mtr/mtrpb"
	"github.com/GeoNet/weft"
	"github.com/golang/protobuf/proto"
	"net/http"
)

type fieldType struct {
}

func (f fieldType) proto(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}

	var err error
	var rows *sql.Rows

	if rows, err = dbR.Query(`SELECT typeID FROM field.type ORDER BY typeID ASC`); err != nil {
		return weft.InternalServerError(err)
	}
	defer rows.Close()

	var ftr mtrpb.FieldTypeResult

	for rows.Next() {
		var ft mtrpb.FieldType

		if err = rows.Scan(&ft.TypeID); err != nil {
			return weft.InternalServerError(err)
		}

		ftr.Result = append(ftr.Result, &ft)
	}

	var by []byte
	if by, err = proto.Marshal(&ftr); err != nil {
		return weft.InternalServerError(err)
	}

	b.Write(by)

	h.Set("Content-Type", "application/x-protobuf")

	return &weft.StatusOK
}
