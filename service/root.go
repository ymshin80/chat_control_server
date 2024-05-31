package service

import (
	"chat_golang_control/repository"
	"chat_golang_control/types/table"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Service struct {
	repository  *repository.Repository
	//ip주소,사용여부
	AvgServerList map[string]bool

}

func NewService(repository  *repository.Repository) *Service{
	s := &Service{repository: repository, AvgServerList: make(map[string]bool)}

	s.setServerInfo()

	if err := s.repository.Kafka.RegisterSubTopic("test-topic"); err != nil {
		panic(err)
	} else {
		go s.loopSubKafka()
	}

	return s
}

func (s *Service) loopSubKafka() {
	for {
		ev := s.repository.Kafka.Pool((100))

		switch event := ev.(type) {
		case *kafka.Message :
			fmt.Println(event)
		case *kafka.Error:
			fmt.Println("failed to pooling event", event.Error())
		}
	}
}

func (s *Service) setServerInfo() {
	if serveList, err := s.GetAvailableServerList(); err != nil {
		panic(err)
	} else {
		for _, server := range serveList{
			s.AvgServerList[server.IP] = true
		}
	}
}
func (s *Service) GetAvailableList() []string {
	var res []string
	for ip, flg := range s.AvgServerList {
		if flg {
			res = append(res,ip)
		}
	}

	return res
}

func (s *Service) GetAvailableServerList()([]*table.ServerInfo, error) {
	return s.repository.GetAvailableServerList()
}