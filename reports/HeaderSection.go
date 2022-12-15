package reports

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func HeaderSection(f *excelize.File){
	err := f.AddPicture("Sheet1", "A1", "./files/assets/GAVIOTA-min.png", `{
        "x_scale": 0.25,
        "y_scale": 0.25
    }`)
	if err != nil {
		fmt.Println(err)
	}
	ApplyFontStyleTitle("J2","LANCHA DE PASAJEROS L/P GAVIOTA",f)
	err = f.AddPicture("Sheet1", "T1", "./files/assets/LANCHA_FOTO.png", `{
        "x_scale": 0.10,
        "y_scale": 0.10
    }`)
	if err != nil {
		fmt.Println(err)
	}
}
