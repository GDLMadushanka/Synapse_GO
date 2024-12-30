package artifacts

import (
	"encoding/xml"
	"synapse/synapsecontext"
)

type Sequence struct {
	MediatorList []Mediator
	LineNo       int
	FileName     string
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

func (v *Sequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v.FileName = ""
	var mediatorList []Mediator

	for {
		t, err := d.Token()
		if err != nil {
			break
		}
		line, _ := d.InputPos()
		switch se := t.(type) {
		case xml.StartElement:
			var mediator Mediator
			switch se.Name.Local {
			case "log":
				logMediator := &LogMediator{}
				if err := d.DecodeElement(logMediator, &se); err != nil {
					return err
				}

				logMediator.LineNo = line
				mediator = logMediator
			case "variable":
				variableMediator := &VariableMediator{}
				if err := d.DecodeElement(variableMediator, &se); err != nil {
					return err
				}
				variableMediator.LineNo = line
				mediator = variableMediator
			case "respond":
				respondMediator := &RespondMediator{}
				if err := d.DecodeElement(respondMediator, &se); err != nil {
					return err
				}
				respondMediator.LineNo = line
				mediator = respondMediator
			case "payloadFactory":
				payloadMediator := &PayloadMediator{}
				if err := d.DecodeElement(payloadMediator, &se); err != nil {
					return err
				}
				payloadMediator.LineNo = line
				mediator = payloadMediator
			case "call":
				callMediator := &CallMediator{}
				if err := d.DecodeElement(callMediator, &se); err != nil {
					return err
				}
				callMediator.LineNo = line
				mediator = callMediator
			}

			if mediator != nil {
				mediatorList = append(mediatorList, mediator)
			}
		}
	}
	v.MediatorList = mediatorList
	return nil
}

func (v *Sequence) SetFileName(fileName string) {
	v.FileName = fileName
	for _, mediator := range v.MediatorList {
		mediator.SetFileName(fileName)
	}
}
