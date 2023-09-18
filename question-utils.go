package main

const (
	LTNone   = ""
	LTNormal = "Normal"
	QTChoice = "Feleletvlasztskrds"
)

const (
	QTAssociation         = "Asszocici"
	QTAssociationOptionRE = `^Vlaszthat([A-Z])$` // Ex.: VlaszthatA = option A
	QTAssociationAnswerRE = `^Vlasz([A-Z])$`
)
