package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type status int

const (
	TODO status = iota
	DOING
	DONE
)

type workStatus struct {
	status status
	conns  []*websocket.Conn

	mu *sync.Mutex
}

type Server struct {
	// workId と workStatus
	works map[string]*workStatus
	// conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		works: make(map[string]*workStatus),
		// conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleActions(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client: ", ws.RemoteAddr())

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		// client から閉じられた時。
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read error: ", err)
			break
		}

		msg := buf[:n]

		client := &ClientMsg{}
		if err := json.Unmarshal(msg, client); err != nil {
			fmt.Println("failed to unmarshal client message to ClientMsg: ", err)
			continue
		}

		if client.WorkId == "" {
			fmt.Println("workID is empty")
			continue
		}

		// initialize work if not found
		if _, found := s.works[client.WorkId]; !found {
			s.works[client.WorkId] = &workStatus{
				status: TODO,
				conns:  []*websocket.Conn{},

				mu: &sync.Mutex{},
			}
		}

		// append ws to conns
		s.appendConn(ws, client.WorkId)

		// s.broadcast(msg)
		fmt.Println(string(msg))

		s.doWorkAndBroadcast(client.WorkId)

		// ws.Write([]byte("Thank you for the message!!!"))
	}
}

// func (s *Server)

func (s *Server) appendConn(ws *websocket.Conn, clientId string) {
	for _, conn := range s.works[clientId].conns {
		if ws == conn {
			return
		}
	}

	fmt.Printf("append conn %v to %s\n", ws, clientId)

	s.works[clientId].conns = append(s.works[clientId].conns, ws)
}

func (s *Server) action(id string, w workStatus) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO であれば、処理を開始する！
	if w.status == TODO {
		w.status = DOING

		s.doWorkAndBroadcast(id)
	}
}

func (s *Server) doWorkAndBroadcast(workId string) {
	work, ok := s.works[workId]
	if !ok {
		fmt.Println("failed to get workId: ", workId)

		return
	}

	workInfo, ok := Works[workId]
	if !ok {
		fmt.Println("failed to get workInfo from Works: ", workId)
	}

	for step := 0; step < workInfo.MaxSteps; step++ {
		// 仕事したつもり
		time.Sleep(3 * time.Second)

		msg := broadcastMsg{
			Type:          1,
			WorkId:        workId,
			StepNumber:    step,
			StepCompleted: true,
			Done:          false,
		}

		s.broadcast(msg, work.conns)
	}

	// 完了通知
	msg := broadcastMsg{
		Type:          1,
		WorkId:        workId,
		StepNumber:    workInfo.MaxSteps,
		StepCompleted: true,
		Done:          true,
	}

	s.broadcast(msg, work.conns)

	// 最後まで終わったらコネクションを閉じる。
	for _, ws := range work.conns {
		go func(ws *websocket.Conn) {
			if err := ws.Close(); err != nil {
				fmt.Println("failed to ws.Close: ", err)
			}
		}(ws)
	}
	work.conns = []*websocket.Conn{}
}

func (s *Server) HandleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client to orderbook feed: ", ws.RemoteAddr())

	maxAge := 10
	cnt := 0
	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())

		ws.Write([]byte(payload))
		time.Sleep(2 * time.Second)

		cnt++
		if cnt > maxAge {
			ws.Write([]byte("Finished!!"))
			break
		}
	}
}

func (s *Server) broadcast(msg broadcastMsg, conns []*websocket.Conn) {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("failed to marshal broadcastMsg: ", err)

		return
	}

	// fmt.Println("broadcast: ", string(msgJson))

	for i, ws := range conns {
		// fmt.Println("conns: ", i)
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msgJson); err != nil {
				fmt.Println("failed to ws.Write: ", err)

				// 書き込み失敗してる時は何かがおかしいので閉じる？？
				ws.Close()
				if conns[i] == ws {
					conns = append(conns[:i], conns[i+1:]...)
				}
			}
		}(ws)
	}
}
