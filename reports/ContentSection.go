package reports

import (
	"gaviotaBackend/variables"
	"github.com/xuri/excelize/v2"
)

func ContentSection(f *excelize.File, sheet string, reserves []variables.Reserve) {
	//Linea 8
	ApplyFontStyleSubtitle("A11", "Check", f)
	SetCellStyleSubtitle("A11", "B12", f)
	ApplyFontStyleSubtitle("C11", "VI. REGISTRO DE PASAJEROS", f)
	SetCellStyleSubtitle("C11", "W11", f)
	//Linea 9
	ApplyFontStyleText("A13", "Si", f)
	ApplyFontStyleText("B13", "No", f)
	SetCellStyleText("A13", "A13", f)
	SetCellStyleText("B13", "B13", f)
	ApplyFontStyleSubtitle("C12", "Nro", f)
	SetCellStyleSubtitle("C12", "C13", f)
	ApplyFontStyleSubtitle("D12", "Apellidos y Nombres (completos)", f)
	SetCellStyleSubtitle("D12", "G13", f)
	ApplyFontStyleSubtitle("H12", "Cédula/Pasaporte", f)
	SetCellStyleSubtitle("H12", "J13", f)
	ApplyFontStyleSubtitle("K12", "Nacionalidad", f)
	SetCellStyleSubtitle("K12", "M13", f)
	ApplyFontStyleSubtitle("N12", "Fecha de Nacimiento", f)
	SetCellStyleSubtitle("N12", "O13", f)
	ApplyFontStyleSubtitle("P12", "Estatus", f)
	SetCellStyleSubtitle("P12", "R12", f)
	ApplyFontStyleSubtitle("P13", "Res", f)
	SetCellStyleSubtitle("P13", "P13", f)
	ApplyFontStyleSubtitle("Q13", "Tra", f)
	SetCellStyleSubtitle("Q13", "Q13", f)
	ApplyFontStyleSubtitle("R13", "Tur", f)
	SetCellStyleSubtitle("R13", "R13", f)
	ApplyFontStyleSubtitle("S12", "Teléfono emergencia", f)
	SetCellStyleSubtitle("S12", "T13", f)
	ApplyFontStyleSubtitle("U12", "Observaciones", f)
	SetCellStyleSubtitle("U12", "W13", f)

	GenerateContent(reserves, f)
}
