package storage

import (
	"context"
	"sync"
)

type Symbol struct {
	symbols map[string]map[string]struct{}
	mu      sync.Mutex
}

func NewSymbol() *Symbol {
	return &Symbol{symbols: make(map[string]map[string]struct{})}
}

func (s *Symbol) SaveSymbols(ctx context.Context, exchange string, symbols []string) error {
	exchangeSymbols := make(map[string]struct{})
	for _, symbol := range symbols {
		exchangeSymbols[symbol] = struct{}{}
	}
	s.mu.Lock()
	s.symbols[exchange] = exchangeSymbols
	s.mu.Unlock()
	return nil
}

func (s *Symbol) ExchangeSymbols(exchange string) []string {
	symbols := make([]string, 0, len(s.symbols))
	for key := range s.symbols[exchange] {
		symbols = append(symbols, key)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return symbols
}

func (s *Symbol) HasSymbol(exchange, symbol string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.symbols[exchange] == nil {
		return false
	}
	_, has := s.symbols[exchange][symbol]
	return has
}

func (s *Symbol) HasExchange(exchange string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.symbols[exchange] == nil {
		return false
	}
	return true
}
