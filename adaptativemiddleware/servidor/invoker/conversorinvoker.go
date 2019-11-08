package invoker

import (
	"ConversorMoedasApp/impl"
	"ConversorMoedasApp/util"
	"RAMid/distribution/marshaller"
	"RAMid/distribution/miop"
	"RAMid/infrastructure/srh"
)

type ConversorInvoker struct{}

func NewConversorInvoker() ConversorInvoker {

	p := new(ConversorInvoker)

	return *p
}

func (ConversorInvoker) Invoke() {

	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: util.PORTA}
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	conversorImpl := impl.Conversor{}

	for {
		// receive data
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// demux request
		switch operation {

		case "Converter":
			_p1 := string(miopPacketRequest.Bd.ReqBody.Body[0].(string))
			_p2 := float64(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = conversorImpl.Converter(_p1, _p2)
		}

		// assembly packet
		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send reply
		srhImpl.Send(msgToClientBytes)
	}
}
