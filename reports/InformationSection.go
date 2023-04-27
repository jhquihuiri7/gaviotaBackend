package reports

import "github.com/xuri/excelize/v2"

func InformationSection(f *excelize.File, sheet, date, time string) {
	var pZarpe, pArribo, hZarpe, hArribo string

	switch time {
	case "Am":
		pZarpe = "Pto. Baq. Moreno"
		pArribo = "Pto. Ayora"
		hZarpe = "07:00"
		hArribo = "09:00"
	case "Pm":
		pZarpe = "Pto. Ayora"
		pArribo = "Pto. Baq. Moreno"
		hZarpe = "15:00"
		hArribo = "17:00"
	}

	//Linea 1
	ApplyFontStyleSubtitle("A4", "I. INFORMACIÓN DEL VIAJE", f)
	ApplyFontStyleSubtitle("P4", "II. INFORMACIÓN DE LA NAVE", f)
	SetCellStyleSubtitle("A4", "O4", f)
	SetCellStyleSubtitle("P4", "W4", f)
	//Linea 2
	ApplyFontStyleSubtitle("A5", "Fecha:", f)
	ApplyFontStyleText("C5", date[:10], f)
	ApplyFontStyleSubtitle("E5", "Puerto zarpe:", f)
	ApplyFontStyleText("H5", pZarpe, f)
	ApplyFontStyleSubtitle("K5", "Hora zarpe:", f)
	ApplyFontStyleText("N5", hZarpe, f)
	SetCellStyleText("A5", "B5", f)
	SetCellStyleText("C5", "D5", f)
	SetCellStyleText("E5", "G5", f)
	SetCellStyleText("H5", "J5", f)
	SetCellStyleText("K5", "M5", f)
	SetCellStyleText("N5", "O5", f)
	ApplyFontStyleSubtitle("A6", "Actividad:", f)
	ApplyFontStyleText("C6", "Cabotaje", f)
	ApplyFontStyleSubtitle("E6", "Puerto arribo:", f)
	ApplyFontStyleText("H6", pArribo, f)
	ApplyFontStyleSubtitle("K6", "Hora arribo:", f)
	ApplyFontStyleText("N6", hArribo, f)
	SetCellStyleText("A6", "B6", f)
	SetCellStyleText("C6", "D6", f)
	SetCellStyleText("E6", "G6", f)
	SetCellStyleText("H6", "J6", f)
	SetCellStyleText("K6", "M6", f)
	SetCellStyleText("N6", "O6", f)

	ApplyFontStyleSubtitle("P5", "Nombre:", f)
	ApplyFontStyleText("R5", "Gaviota", f)
	ApplyFontStyleSubtitle("T5", "Cap. Tripulantes:", f)
	ApplyFontStyleText("W5", "3", f)
	SetCellStyleText("P5", "Q5", f)
	SetCellStyleText("R5", "S5", f)
	SetCellStyleText("T5", "V5", f)
	SetCellStyleText("W5", "W5", f)

	ApplyFontStyleSubtitle("P6", "Matrícula:", f)
	ApplyFontStyleText("R6", "TN-01-01070", f)
	ApplyFontStyleSubtitle("T6", "Cap. Pasajeros:", f)
	ApplyFontStyleText("W6", "38", f)
	SetCellStyleText("P6", "Q6", f)
	SetCellStyleText("R6", "S6", f)
	SetCellStyleText("T6", "V6", f)
	SetCellStyleText("W6", "W6", f)

	//Linea 4
	ApplyFontStyleSubtitle("A7", "III. INFORMACIÓN ARMADOR", f)
	ApplyFontStyleSubtitle("H7", "IV. INFORMACIÓN RESPONSABLE DEL EMBARQUE", f)
	ApplyFontStyleSubtitle("P7", "V. INFORMACIÓN TRIPULANTES", f)
	SetCellStyleSubtitle("A7", "G7", f)
	SetCellStyleSubtitle("H7", "O7", f)
	SetCellStyleSubtitle("P7", "W7", f)

	//Linea 5
	ApplyFontStyleSubtitle("A8", "Nombre:", f)
	ApplyFontStyleText("C8", "Darwin Ernesto Freire Escarabay", f)
	ApplyFontStyleSubtitle("H8", "Nombre:", f)
	ApplyFontStyleText("J8", "Carlos Zambrano Macías", f)
	ApplyFontStyleSubtitle("P8", "Capitán:", f)
	ApplyFontStyleText("R8", "Carlos Zambrano", f)
	ApplyFontStyleSubtitle("T8", "Cédula:", f)
	ApplyFontStyleText("V8", "2000079711", f)
	SetCellStyleText("A8", "B8", f)
	SetCellStyleText("C8", "G8", f)
	SetCellStyleText("H8", "I8", f)
	SetCellStyleText("J8", "O8", f)
	SetCellStyleText("P8", "Q8", f)
	SetCellStyleText("R8", "S8", f)
	SetCellStyleText("T8", "U8", f)
	SetCellStyleText("V8", "W8", f)

	//Linea 6
	ApplyFontStyleSubtitle("A9", "RUC", f)
	ApplyFontStyleText("C9", "2000026241001", f)
	ApplyFontStyleSubtitle("E9", "Teléf:", f)
	ApplyFontStyleText("F9", "0981915924", f)
	ApplyFontStyleSubtitle("H9", "Cédula:", f)
	ApplyFontStyleText("J9", "2000079711", f)
	ApplyFontStyleSubtitle("L9", "Teléfono:", f)
	ApplyFontStyleText("N9", "0969348203", f)
	ApplyFontStyleSubtitle("P9", "Marinero 1:", f)
	ApplyFontStyleText("R9", "Jeferson Guerrero", f)
	ApplyFontStyleSubtitle("T9", "Cédula:", f)
	ApplyFontStyleText("V9", "2050000864", f)
	SetCellStyleText("A9", "B9", f)
	SetCellStyleText("C9", "D9", f)
	SetCellStyleText("E9", "E9", f)
	SetCellStyleText("F9", "G9", f)
	SetCellStyleText("H9", "I9", f)
	SetCellStyleText("J9", "K9", f)
	SetCellStyleText("L9", "M9", f)
	SetCellStyleText("N9", "O9", f)
	SetCellStyleText("P9", "Q9", f)
	SetCellStyleText("R9", "S9", f)
	SetCellStyleText("T9", "U9", f)
	SetCellStyleText("V9", "W9", f)
	//Linea 7
	ApplyFontStyleSubtitle("A10", "e-mail:", f)
	ApplyFontStyleText("C10", "gaviota.ferry@gmail.com", f)
	ApplyFontStyleSubtitle("H10", "e-mail:", f)
	ApplyFontStyleText("J10", "", f)
	ApplyFontStyleSubtitle("P10", "", f)
	ApplyFontStyleText("R10", "", f)
	ApplyFontStyleSubtitle("T10", "", f)
	ApplyFontStyleText("V10", "", f)
	SetCellStyleText("A10", "B10", f)
	SetCellStyleText("C10", "G10", f)
	SetCellStyleText("H10", "I10", f)
	SetCellStyleText("J10", "O10", f)
	SetCellStyleText("P10", "Q10", f)
	SetCellStyleText("R10", "S10", f)
	SetCellStyleText("T10", "U10", f)
	SetCellStyleText("V10", "W10", f)

}
