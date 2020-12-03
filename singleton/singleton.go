package singleton

import (
	"log"
	"net/http"
	"net/url"
	"sync"
)

type Service struct {
	Url    string
	client http.Client
}

type ImageService struct {
	Service
}

func (s ImageService) getSaveImageUrl() string {
	return s.Url + "/image"
}

func (s ImageService) saveImage(image string) *http.Response {
	resp, err := s.client.PostForm(s.getSaveImageUrl(), url.Values{"image": {image}})
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func New() *ImageService {
	return &ImageService{
		Service{
			"http://localhost:8000/image",
			http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       10,
			},
		},
	}
}

var once sync.Once
var Image *ImageService

func GetInstance(channel chan *ImageService, wg *sync.WaitGroup) {
	defer wg.Done()
	if Image == nil {
		once.Do(func() {
			Image = New()
		})
		channel <- Image
	} else {
		channel <- Image
	}
}
