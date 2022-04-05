package rankserver

import (
	"context"
	"io"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/apis"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/frinterface"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &RankServer{}

type RankServer struct {
	sync.Mutex                            //为了map的线程安全，放了个锁
	RankS      frinterface.ServeInterface //这是个接口
	PersonCh   chan *apis.PersonalInformation
}

func (r *RankServer) Update(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	r.regPerson(information)
	return r.RankS.UpdatePersonalInformation(information)
}

func (r *RankServer) GetFatRate(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalRank, error) {
	return r.RankS.GetFatRate(information.Name)
}

func (r *RankServer) GetTop(ctx context.Context, null *apis.Null) (*apis.PersonalRanks, error) {
	top, err := r.RankS.GetTop()
	if err != nil {
		log.Println("获取榜单时出错：", err)
		return nil, err
	}
	return &apis.PersonalRanks{PersonalRanks: top}, nil
}

///----------------
//单管道案例
func (r *RankServer) regPerson(pi *apis.PersonalInformation) {
	r.RankS.RegisterPersonalInformation(pi) // todo handle error
	r.PersonCh <- pi
}

func (r *RankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.PersonCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束：", err)
			return err
		}
	}
	return nil
}

func (r *RankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv() // 客户端发送很多，服务端要不停接收。直到发完了，给我一个答复
		if err == io.EOF {       //这里的io.EOF 表示消息已经发完了，正常结束
			break
		}
		if err != nil && err != io.EOF { //err不为空，且不是io.EOF，说明是真的错误。
			log.Printf("WARNING: 获取人员注册时失败：%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)

		r.regPerson(pi)
	}

	log.Println("得到连续注册清单：", pis.String())
	//接收完消息，发送，并且关闭
	return server.SendAndClose(pis) //
}

func (r *RankServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人： %s\n", information.String()) //information.String() 把information转成string类型
	return information, nil
}
