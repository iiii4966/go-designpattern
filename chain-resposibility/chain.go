package chain_resposibility

import "fmt"

type Image struct {
	Size int
}

type ImageStorage interface {
	Save(image *Image) (*Image, error)
	SetNext(storage ImageStorage)
}

type S3 struct {
	next ImageStorage
	limit int
}

func NewS3() *S3 {
	return &S3{limit: 10000}
}

func (s S3) Save(image *Image) (*Image, error){
	if image.Size > s.limit && s.next != nil {
		return s.next.Save(image)
	}
	// save image
	fmt.Println("S3 save image")
	return image, nil
}

func (s *S3) SetNext(storage ImageStorage){
	s.next = storage
}

type Local struct {
	next ImageStorage
	limit int
}

func NewLocal() *Local {
	return &Local{limit: 20000}
}

func (l Local) Save(image *Image) (*Image, error){
	if image.Size > l.limit && l.next != nil {
		return l.next.Save(image)
	}
	// save image
	fmt.Println("Local save image")
	return image, nil
}

func (l *Local) SetNext(storage ImageStorage){
	l.next = storage
}

type Other struct {
	next ImageStorage
	limit int
}

func NewOther() *Other {
	return &Other{limit: 30000}
}

func (o Other) Save(image *Image) (*Image, error){
	if image.Size > o.limit && o.next != nil {
		return o.next.Save(image)
	}
	// save image
	fmt.Println("Other save image")
	return image, nil
}

func (o *Other) SetNext(storage ImageStorage){
	o.next = storage
}


