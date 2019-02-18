package producer

import (
	"github.com/nsqio/go-nsq"
	"sync"
)

var fileProducer *nsq.Producer

type FileProducer struct {
	Producer *nsq.Producer
}

var instance *FileProducer
var once sync.Once

func GetInstance() *FileProducer {
	once.Do(func() {
		instance = &FileProducer{}
	})
	return instance
}

func (p *FileProducer) InitProducer(url string) error {
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		return err
	}

	fileProducer = producer
	return nil
}

func (p *FileProducer) GetFileProducer() *nsq.Producer {
	return fileProducer
}

func (p *FileProducer) Publish(topicName string, data []byte) error {
	return fileProducer.Publish(topicName, data)
}
