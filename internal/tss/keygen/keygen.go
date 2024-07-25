package keygen

import (
	"context"
	"fmt"
	"sync"

	"tss_project/proto"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedTSSServiceServer
	parties map[string]*Party
	mu      sync.Mutex
}

type Party struct {
	ID       string
	IP       string
	PartyID  *tss.PartyID
	Out      chan tss.Message
	In       chan tss.Message
	SaveData *keygen.LocalPartySaveData
	Logger   *zap.SugaredLogger
	clients  map[string]proto.TSSServiceClient
	mu       sync.Mutex
}

func NewParty(id string, ip string, logger *zap.SugaredLogger) *Party {
	return &Party{
		ID:      id,
		IP:      ip,
		Out:     make(chan tss.Message, 100),
		In:      make(chan tss.Message, 100),
		Logger:  logger,
		clients: make(map[string]proto.TSSServiceClient),
	}
}

func (s *Server) ExchangeMessage(ctx context.Context, req *proto.MessageWrapper) (*proto.MessageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fromParty, exists := s.parties[req.From.Id]
	if !exists {
		return nil, fmt.Errorf("party %s not found", req.From.Id)
	}

	msg, err := tss.ParseWireMessage(req.Content.Value, fromParty.PartyID, req.IsBroadcast)
	if err != nil {
		return nil, fmt.Errorf("failed to parse message: %v", err)
	}

	fromParty.In <- msg
	return &proto.MessageResponse{Success: true}, nil
}

func (p *Party) Start(ctx context.Context, params *tss.Parameters) error {
	p.PartyID = params.PartyID()
	party := keygen.NewLocalParty(params, p.Out, nil)
	go p.handleOutgoingMessages()

	if err := party.Start(); err != nil {
		return fmt.Errorf("failed to start party: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-p.In:
			if _, err := party.UpdateFromBytes(msg.WireBytes(), msg.GetFrom(), msg.IsBroadcast()); err != nil {
				p.Logger.Errorf("Failed to update party: %v", err)
			}
		case save := <-party.EndCh():
			p.SaveData = save
			return nil
		}
	}
}

func (p *Party) handleOutgoingMessages() {
	for msg := range p.Out {
		go p.sendMessage(msg)
	}
}

func (p *Party) sendMessage(msg tss.Message) {
	data, _ := msg.WireBytes()
	if msg.IsBroadcast() {
		// Broadcast to all other parties
		for _, client := range p.clients {
			go client.ExchangeMessage(context.Background(), &proto.MessageWrapper{
				From: &proto.PartyID{
					Id: p.ID,
				},
				Content: &any.Any{
					Value: data,
				},
				IsBroadcast: true,
			})
		}
	} else {
		// Send to specific party
		to := msg.GetTo()[0]
		client, exists := p.clients[to.Id]
		if exists {
			go client.ExchangeMessage(context.Background(), &proto.MessageWrapper{
				From: &proto.PartyID{
					Id: p.ID,
				},
				Content: &any.Any{
					Value: data,
				},
				IsBroadcast: false,
			})
		}
	}
}

func (p *Party) AddClient(id string, client proto.TSSServiceClient) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clients[id] = client
}

func RunKeygen(ctx context.Context, parties []*Party, threshold int) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(parties))

	for _, p := range parties {
		wg.Add(1)
		go func(party *Party) {
			defer wg.Done()
			params := tss.NewParameters(tss.Edwards(), tss.NewPeerContext(tss.SortPartyIDs(getPartyIDs(parties))), party.PartyID, len(parties), threshold)
			if err := party.Start(ctx, params); err != nil {
				errCh <- fmt.Errorf("party %s failed: %v", party.ID, err)
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func getPartyIDs(parties []*Party) []*tss.PartyID {
	ids := make([]*tss.PartyID, len(parties))
	for i, p := range parties {
		ids[i] = p.PartyID
	}
	return ids
}
