package invoker

import (
	"ConversorMoedasApp/impl"
	"ConversorMoedasApp/util"
	"RAMid/distribution/miop"
	"RAMid/plugins"
	"fmt"
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
	Receive := funcReceive.(func(chan [3]interface{}))

	funcSend, err := srhInst.Lookup("Send")
	Send := funcSend.(func(chan [2]interface{}))

	marshallerInst, err := plugin.Open(manager.ObterComponente(util.ID_MARSHALLER))

	funcUnmarshall, err := marshallerInst.Lookup("Unmarshall")
	Unmarshall := funcUnmarshall.(func(chan interface{}))

	funcMarshall, err := marshallerInst.Lookup("Marshall")
	Marshall := funcMarshall.(func(chan interface{}))

	conversorImpl := impl.Conversor{}

	chReceiveSrh := make(chan [3]interface{})
	chUnmarshall := make(chan interface{})
	chMarshaller := make(chan interface{})
	chSendSrh := make(chan [2]interface{})

	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	for {

		// receive data
		go Receive(chReceiveSrh)

		var parametros [3]interface{}
		parametros[0] = util.SERVER_HOST
		parametros[1] = util.SERVER_PORT

		// send request packet and receive reply packet
		chReceiveSrh <- parametros
		retornoReceive := <-chReceiveSrh

		rcvMsgBytes := retornoReceive[2].([]byte)

		// unmarshall
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
		go Marshall(chMarshaller)

		// serialise request packet
		chMarshaller <- miopPacketReply
		retornoMarshall := <-chMarshaller
		msgToClientBytes := retornoMarshall.([]byte)

		// send reply
		go Send(chSendSrh)

		var parametrosSend [2]interface{}
		parametrosSend[0] = msgToClientBytes

		// send request packet and receive reply packet
		chSendSrh <- parametrosSend

		retornoSend := <-chSendSrh

		mensagemEnviada := retornoSend[1].(bool)

		if mensagemEnviada {
			fmt.Println("Mensagem enviada")
		}
	}
}
