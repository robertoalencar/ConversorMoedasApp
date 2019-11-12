package invoker

import (
	"ConversorMoedasApp/impl"
	"ConversorMoedasApp/util"
	"RAMid/distribution/miop"
	"RAMid/plugins"
	"plugin"
)

type ConversorInvoker struct{}

func NewConversorInvoker() ConversorInvoker {

	p := new(ConversorInvoker)

	return *p
}

func (ConversorInvoker) Invoke() {

	manager := plugins.Manager{}

	srhInst, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_SRH))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcReceive, err := srhInst.Lookup("Receive")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Receive := funcReceive.(func(chan [3]interface{}))

	funcSend, err := srhInst.Lookup("Send")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Send := funcSend.(func(chan []byte))

	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	marshallerInst, err := plugin.Open(manager.ObterComponente(util.ID_MARSHALLER))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcUnmarshall, err := marshallerInst.Lookup("Unmarshall")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Unmarshall := funcUnmarshall.(func(chan interface{}))

	conversorImpl := impl.Conversor{}

	for {

		// receive data
		chReceiveSrh := make(chan [3]interface{})
		go Receive(chReceiveSrh)

		var parametros [3]interface{}
		parametros[0] = util.SERVER_HOST
		parametros[1] = util.SERVER_PORT

		// send request packet and receive reply packet
		chReceiveSrh <- parametros
		retornoReceive := <-chReceiveSrh

		rcvMsgBytes := retornoReceive[2].([]byte)

		// unmarshall
		chUnmarshall := make(chan interface{})
		go Unmarshall(chUnmarshall)

		chUnmarshall <- rcvMsgBytes
		retornoUnmarshall := <-chUnmarshall
		miopPacketRequest := retornoUnmarshall.(miop.Packet)

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
		funcMarshall, err := marshallerInst.Lookup("Marshall")
		util.ChecaErro(err, "Falha ao carregar a função do componente")

		Marshall := funcMarshall.(func(chan interface{}))

		chMarshaller := make(chan interface{})
		go Marshall(chMarshaller)

		// serialise request packet
		chMarshaller <- miopPacketReply
		retornoMarshall := <-chMarshaller
		msgToClientBytes := retornoMarshall.([]byte)

		// send reply
		chSendSrh := make(chan []byte)
		go Send(chSendSrh)
		chSendSrh <- msgToClientBytes
	}
}
