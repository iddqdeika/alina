package requester

import (
	"alina/alina"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrCfg        = errors.New("invalid config")
	ErrLogger     = errors.New("invalid logger")
	ErrDispatcher = errors.New("invalid dispatcher")
)

func New(config alina.Config, logger alina.Logger, dispatcher alina.Dispatcher) (alina.Requester, error) {
	r := &requester{}

	if config == nil || len(config.GetAccessToken()) == 0 {
		return nil, ErrCfg
	}
	r.config = config

	if logger == nil {
		return nil, ErrLogger
	}
	r.logger = logger

	if dispatcher == nil {
		return nil, ErrDispatcher
	}
	r.dispatcher = dispatcher

	r.stop = make(chan struct{})
	return r, nil
}

type requester struct {
	config      alina.Config
	logger      alina.Logger
	dispatcher  alina.Dispatcher
	key         string
	server      string
	ts          string
	stop        chan struct{}
	lastRequest time.Time
}

func (r *requester) Init() error {
	err := r.getLongPollServer()
	if err != nil {
		return fmt.Errorf("error durint requester initialisation: %v", err)
	}
	return nil
}

func (r *requester) Run() {
	r.run()
}

func (r *requester) run() {
	dataChan := make(chan *updateResponse)
	go r.getUpdates(dataChan)
	for {
		select {
		case <-r.stop:
			return
		case ud := <-dataChan:
			r.handleUpdateResponse(ud)
			go r.getUpdates(dataChan)
		}
	}
}

func (r *requester) GetServer() (string, error) {
	panic("implement me")
}

func (r *requester) GetKey() (string, error) {
	panic("implement me")
}

func (r *requester) handleUpdateResponse(ud *updateResponse) {
	if ud.err != nil {
		r.logger.Error(ud.err)
	}
	if ud.data != nil {
		err := r.handleUpdateData(ud.data)
		if err != nil {
			r.logger.Error("error during handling of updates response")
			return
		}
	}
}

func (r *requester) handleUpdateData(data []byte) error {
	ur := &updatesResponseBody{}
	err := json.Unmarshal(data, ur)
	if err != nil {
		return fmt.Errorf("error during update response data parsing: %v", err)
	}

	for _, v := range ur.Updates {
		r.dispatcher.Handle(v)
	}

	r.ts = ur.Ts
	return nil
}

func (r *requester) getUpdates(ch chan *updateResponse) {
	r.logger.Info("getting updates...")
	data, err := r.getUpdateData()
	ub := &updateResponse{
		data: data,
		err:  err,
	}
	r.logger.Info("updates got")
	ch <- ub
}

func (r *requester) getUpdateData() ([]byte, error) {
	now := time.Now()

	target := r.lastRequest.Add(r.config.GetLongPollInterval())
	if now.Before(target) {
		<-time.Tick(r.config.GetLongPollInterval())
	}
	r.lastRequest = time.Now()
	url := r.getLongPollUpdateUrl()
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error during longpoll update request: %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code from longpoll update request")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error during response body reading from longpoll update request: %v", err)
	}
	return data, nil
}

func (r *requester) getLongPollUpdateUrl() string {
	return fmt.Sprintf("%v?act=a_check&key=%v&ts=%v&wait=25", r.server, r.key, r.ts)
}

func (r *requester) getLongPollServer() error {
	url := r.getLongPollServerUrl()
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("status code = %v from url \"%v\"", res.StatusCode, url)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	rs := &getLongPollServerStruct{}
	err = json.Unmarshal(data, rs)
	if err != nil {
		return err
	}

	r.server = rs.Response.Server
	r.key = rs.Response.Key
	r.ts = rs.Response.Ts
	return nil
}

func (r *requester) getLongPollServerUrl() string {
	return "https://api.vk.com/method/groups.getLongPollServer?group_id=" + r.config.GetGroupId() + "&access_token=" + r.config.GetAccessToken() + "&v=" + r.config.GetVersion()
}

func (r *requester) SendGet(methodName string, paramMap map[string]string) ([]byte, error) {
	var u *url.URL
	u, err := url.Parse("https://api.vk.com")
	if err != nil {
		return nil, fmt.Errorf("error during url parsing: &v", err)
	}
	u.Path += fmt.Sprintf("/method/%v", methodName)
	params := url.Values{}
	params.Add("access_token", r.config.GetAccessToken())
	params.Add("v", r.config.GetVersion())

	//var parameters string
	for key, value := range paramMap {
		//if len(parameters) > 0 {
		//	parameters += "&"
		//}
		//parameters += key + "=" + value
		params.Add(key, value)
	}

	u.RawQuery = params.Encode()
	//url := fmt.Sprintf("https://api.vk.com/method/%v?%v&access_token=%v&v=%v", methodName, parameters, r.config.GetAccessToken(), r.config.GetVersion())

	url := u.String()
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error during sending request for method %v, err:%v", methodName, err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error during sending request for methid %v, err:%v", err)
	}
	if resp.StatusCode != 200 {
		return data, fmt.Errorf("error during sending request for method %v, code:%v", methodName, resp.StatusCode)
	}

	return data, nil
}
