package server

import (
	"errors"

	rpc "github.com/pravahio/go-mesh/rpc"
)

// Subscribe serves a subscribe request over RPC
func (s *Server) Subscribe(info *rpc.PeerTopicInfo, stream rpc.Mesh_SubscribeServer) error {
	if s.m.SubService == nil {
		return errors.New("RPC is not running as a subscriber")
	}

	log.Info("[RPC] Subscribing to ", info.GetTopics())

	//ctx := stream.Context()

	msgChan, err := s.m.SubService.SubscribeToTopics(info.GetTopics())
	if err != nil {
		log.Error(err)
		return err
	}

	for {
		// TODO: What id the chan is closed.
		data := <-msgChan

		err := stream.Send(&rpc.Data{
			Topic: data.GetTopicIDs(),
			Raw:   data.GetData(),
		})

		if err != nil {
			log.Warning(err)
			return err
		}
		log.Info("RPC Data Sent: ", err)
	}

}
