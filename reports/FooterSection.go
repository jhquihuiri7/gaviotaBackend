package reports

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func FooterSection(cellStart int, sheet string,f *excelize.File){
	f.SetRowHeight("Sheet1", cellStart+1, 22)
	for i:=cellStart+2; i<=cellStart+8;i++{
		f.SetRowHeight("Sheet1", i, 15)
	}
	//f.SetRowHeight("Sheet1", cellStart+2, 20)
	//f.SetRowHeight("Sheet1", cellStart+3, 20)
	//f.SetRowHeight("Sheet1", cellStart+4, 20)
	//f.SetRowHeight("Sheet1", cellStart+5, 20)
	//f.SetRowHeight("Sheet1", cellStart+6, 20)
	//f.SetRowHeight("Sheet1", cellStart+7, 20)
	//f.SetRowHeight("Sheet1", cellStart+8, 20)
	//f.SetRowHeight("Sheet1", cellStart+9, 20)
	//f.SetRowHeight("Sheet1", cellStart+10, 20)
	//Linea 11
	ApplyFontStyleSubtitle(fmt.Sprintf("A%d",cellStart+1),"VII. RESPONSABLE DEL REGISTRO DE PASAJEROS",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+1),"VIII. CAPITANÍA DEL PUERTO (Zarpe)",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+1),"IX. GAD MUNICIPAL (Recepción)",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+1),"X. CAPITANIA DE PUERTO (Control)",f)
	SetCellStyleSubtitle(fmt.Sprintf("A%d",cellStart+1),fmt.Sprintf("K%d",cellStart+1),f)
	SetCellStyleSubtitle(fmt.Sprintf("L%d",cellStart+1),fmt.Sprintf("O%d",cellStart+1),f)
	SetCellStyleSubtitle(fmt.Sprintf("P%d",cellStart+1),fmt.Sprintf("S%d",cellStart+1),f)
	SetCellStyleSubtitle(fmt.Sprintf("T%d",cellStart+1),fmt.Sprintf("W%d",cellStart+1),f)
	//Linea 12
	ApplyFontStyleMini(fmt.Sprintf("A%d",cellStart+2),"Declaración de responsabilidad: ","El Armador asume toda responsabilidad legal sobre los actos relacionados con la operación de la embarcación, incluido el registro de pasajeros. Asimismo, como persona responsable del registro de pasajeros ","DECLARO "," que la información detallada en el presente formulario es verídica en su totalidad, asimismo, conozco que puede estar sujeto al análisis que en derecho corresponda y que es de mi entera responsabilidad cualquier tipo de falsificación, destrucción, adulteración, modificación u omisión en la información proporcionada a las Autoridades competentes.",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+2),"Nombre:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+2),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+2),"Nombre:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+2),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+2),"Nombre:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+2),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+2),"Nombre:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+2),"",f)
	SetCellStyleMini(fmt.Sprintf("A%d",cellStart+2),fmt.Sprintf("G%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+2),fmt.Sprintf("I%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+2),fmt.Sprintf("K%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+2),fmt.Sprintf("M%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+2),fmt.Sprintf("O%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+2),fmt.Sprintf("Q%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+2),fmt.Sprintf("S%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+2),fmt.Sprintf("U%d",cellStart+2),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+2),fmt.Sprintf("W%d",cellStart+2),f)


	//Linea 13
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+3),"Cédula:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+3),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+3),"Cédula:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+3),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+3),"Cédula:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+3),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+3),"Cédula:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+3),"",f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+3),fmt.Sprintf("I%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+3),fmt.Sprintf("K%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+3),fmt.Sprintf("M%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+3),fmt.Sprintf("O%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+3),fmt.Sprintf("Q%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+3),fmt.Sprintf("S%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+3),fmt.Sprintf("U%d",cellStart+3),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+3),fmt.Sprintf("W%d",cellStart+3),f)

	//Linea 14
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+4),"Cargo:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+4),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+4),"Cargo:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+4),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+4),"Cargo:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+4),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+4),"Cargo:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+4),"",f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+4),fmt.Sprintf("I%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+4),fmt.Sprintf("K%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+4),fmt.Sprintf("M%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+4),fmt.Sprintf("O%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+4),fmt.Sprintf("Q%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+4),fmt.Sprintf("S%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+4),fmt.Sprintf("U%d",cellStart+4),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+4),fmt.Sprintf("W%d",cellStart+4),f)

	//Linea 15
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+5),"Fecha:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+5),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+5),"Fecha:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+5),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+5),"Fecha:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+5),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+5),"Fecha:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+5),"",f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+5),fmt.Sprintf("I%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+5),fmt.Sprintf("K%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+5),fmt.Sprintf("M%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+5),fmt.Sprintf("O%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+5),fmt.Sprintf("Q%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+5),fmt.Sprintf("S%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+5),fmt.Sprintf("U%d",cellStart+5),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+5),fmt.Sprintf("W%d",cellStart+5),f)

	//Linea 16
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+6),"Teléfono:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+6),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+6),"Teléfono:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+6),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+6),"Teléfono:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+6),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+6),"Teléfono:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+6),"",f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+6),fmt.Sprintf("I%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+6),fmt.Sprintf("K%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+6),fmt.Sprintf("M%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+6),fmt.Sprintf("O%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+6),fmt.Sprintf("Q%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+6),fmt.Sprintf("S%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+6),fmt.Sprintf("U%d",cellStart+6),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+6),fmt.Sprintf("W%d",cellStart+6),f)

	//Linea 17
	ApplyFontStyleSubtitle(fmt.Sprintf("H%d",cellStart+7),"Firma y sello:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("J%d",cellStart+7),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("L%d",cellStart+7),"Firma y sello:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("N%d",cellStart+7),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("P%d",cellStart+7),"Firma y sello:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("R%d",cellStart+7),"",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("T%d",cellStart+7),"Firma y sello:",f)
	ApplyFontStyleSubtitle(fmt.Sprintf("V%d",cellStart+7),"",f)
	SetCellStyleText(fmt.Sprintf("H%d",cellStart+7),fmt.Sprintf("I%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("J%d",cellStart+7),fmt.Sprintf("K%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("L%d",cellStart+7),fmt.Sprintf("M%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("N%d",cellStart+7),fmt.Sprintf("O%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("P%d",cellStart+7),fmt.Sprintf("Q%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("R%d",cellStart+7),fmt.Sprintf("S%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("T%d",cellStart+7),fmt.Sprintf("U%d",cellStart+8),f)
	SetCellStyleText(fmt.Sprintf("V%d",cellStart+7),fmt.Sprintf("W%d",cellStart+8),f)
}

