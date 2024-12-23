package pogo

import (
	"encoding/xml"
	"test/mediators"
	"test/synapsecontext"
)

type Sequence struct {
	MediatorList []Mediator
}

const log = "log"
const variable = "variable"
const respond = "respond"

func (v *Sequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var mediatorList []Mediator

	for {
		t, err := d.Token()
		if err != nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			var mediator Mediator
			switch se.Name.Local {
			case log:
				var logMediator mediators.LogMediator
				if err := d.DecodeElement(&logMediator, &se); err != nil {
					return err
				}
				mediator = logMediator
			case variable:
				var variableMediator mediators.VariableMediator
				if err := d.DecodeElement(&variableMediator, &se); err != nil {
					return err
				}
				mediator = variableMediator
			case respond:
				var respondMediator mediators.RespondMediator
				if err := d.DecodeElement(&respondMediator, &se); err != nil {
					return err
				}
				mediator = respondMediator
			}

			if mediator != nil {
				mediatorList = append(mediatorList, mediator)
			}
		}
	}
	v.MediatorList = mediatorList
	return nil
}

func (v *Sequence) Execute(context *synapsecontext.SynapseContext) bool {
	for _, mediator := range v.MediatorList {
		result := mediator.Execute(context)
		if !result {
			return false
		}
	}
	return true
}
