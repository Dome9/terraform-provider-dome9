package iplist

import (
	"fmt"
	"net/http"
)

const (
	ipListResourcePath = "iplist"
)

type IpList struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Items       []Item `json:"items,omitempty"`
}

type Item struct {
	Ip      string `json:"ip,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (service *Service) Get(ipListId int64) (*IpList, *http.Response, error) {
	v := new(IpList)
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]IpList, *http.Response, error) {
	v := new([]IpList)
	resp, err := service.Client.NewRequestDo("GET", ipListResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(ipList *IpList) (*IpList, *http.Response, error) {
	v := new(IpList)
	resp, err := service.Client.NewRequestDo("POST", ipListResourcePath, nil, ipList, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(ipListId int64, ipList *IpList) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, ipList, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(ipListId int64) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
