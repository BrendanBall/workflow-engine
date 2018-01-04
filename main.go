package main

import "fmt"

func main() {
	buyTokenStage.Run(&Request{})
}

type Request struct {
	Data map[string]string
}

type Result struct {
	Success bool
	Request *Request
}

type Stage interface {
	Run(request *Request) Result
}

type Pipeline struct {
	Name   string
	Stages []Stage
}

func (p *Pipeline) Run(request *Request) Result {
	fmt.Printf("Start %s\n", p.Name)
	defer fmt.Printf("End %s\n", p.Name)
	for _, s := range p.Stages {
		result := s.Run(request)
		if !result.Success {
			return result
		}
	}
	return Result{true, request}
}

type ReserveMoneyStage struct {
}

func (p *ReserveMoneyStage) Run(request *Request) Result {
	fmt.Println("reserving money")
	return Result{true, request}
}

type CommitMoneyStage struct {
}

func (p *CommitMoneyStage) Run(request *Request) Result {
	fmt.Println("committing money")
	return Result{true, request}
}

type GetTokenStage struct {
}

func (p *GetTokenStage) Run(request *Request) Result {
	fmt.Println("getting token")
	return Result{true, request}
}

type SmsStage struct {
}

func (p *SmsStage) Run(request *Request) Result {
	fmt.Println("sending sms")
	return Result{true, request}
}

var (
	reserveMoneyStage = &ReserveMoneyStage{}
	commitMoneyStage  = &CommitMoneyStage{}
	authStage         = &Pipeline{"Auth Pipeline", []Stage{reserveMoneyStage}}
	getTokenStage     = &GetTokenStage{}
	smsStage          = &SmsStage{}
	notificationStage = &Pipeline{"Notification Pipeline", []Stage{smsStage}}
	buyTokenStage     = Pipeline{"Buy Token Pipeline", []Stage{authStage, getTokenStage, commitMoneyStage, notificationStage}}
)
