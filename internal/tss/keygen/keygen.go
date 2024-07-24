package keygen

import (
	"fmt"
	"sync"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
	"go.uber.org/zap"
)

type Party struct {
	params    *tss.Parameters
	id        string
	out       chan tss.Message
	end       chan keygen.LocalPartySaveData
	party     *keygen.LocalParty
	logger    *zap.SugaredLogger
	publicKey string
	mu        sync.Mutex
}

func NewParty(params *tss.Parameters, id string, logger *zap.SugaredLogger) *Party {
	out := make(chan tss.Message, 1000)
	end := make(chan keygen.LocalPartySaveData, 1)
	return &Party{
		params: params,
		id:     id,
		out:    out,
		end:    end,
		logger: logger,
	}
}

func (p *Party) Start() error {
	p.party = keygen.NewLocalParty(p.params, p.out, p.end)
	go p.handleOutgoingMessages()

	if err := p.party.Start(); err != nil {
		return fmt.Errorf("failed to start party: %v", err)
	}

	saveData := <-p.end
	p.mu.Lock()
	p.publicKey = saveData.ECDSAPub.X().String() + saveData.ECDSAPub.Y().String()
	p.mu.Unlock()

	return nil
}

func (p *Party) handleOutgoingMessages() {
	for msg := range p.out {
		// Here, implement the logic to send the message to other parties
		// This could involve using Kubernetes API or your preferred communication method
		p.logger.Infof("Party %s sending message to %v", p.id, msg.GetTo())
	}
}

func (p *Party) GetPublicKey() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.publicKey
}

func (p *Party) UpdateFromBytes(wireBytes []byte, from *tss.PartyID, isBroadcast bool) (bool, error) {
	return p.party.UpdateFromBytes(wireBytes, from, isBroadcast)
}
