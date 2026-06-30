package infrastructure

import "sync"

type Subscriber struct {
    Send chan []byte
}

type Hub struct {
    mu          sync.RWMutex
    subscribers map[int32]map[*Subscriber]bool
}

func NewHub() *Hub {
    return &Hub{
        subscribers: make(map[int32]map[*Subscriber]bool),
    }
}

func (h *Hub) Subscribe(rideID int32, sub *Subscriber) {
    h.mu.Lock()
    defer h.mu.Unlock()
    if h.subscribers[rideID] == nil {
        h.subscribers[rideID] = make(map[*Subscriber]bool)
    }
    h.subscribers[rideID][sub] = true
}

func (h *Hub) Unsubscribe(rideID int32, sub *Subscriber) {
    h.mu.Lock()
    defer h.mu.Unlock()
    if subs, ok := h.subscribers[rideID]; ok {
        if _, exists := subs[sub]; exists {
            delete(subs, sub)
            close(sub.Send)
            if len(subs) == 0 {
                delete(h.subscribers, rideID)
            }
        }
    }
}

func (h *Hub) Publish(rideID int32, message []byte) {
    h.mu.RLock()
    defer h.mu.RUnlock()
    for sub := range h.subscribers[rideID] {
        select {
        case sub.Send <- message:
        default:
        }
    }
}

// CloseAll notifica a todos los suscriptores del viaje y cierra sus canales.
// Usar cuando el conductor se desconecta para que los clientes sepan que el stream terminó.
func (h *Hub) CloseAll(rideID int32) {
    h.mu.Lock()
    defer h.mu.Unlock()
    if subs, ok := h.subscribers[rideID]; ok {
        for sub := range subs {
            close(sub.Send)
        }
        delete(h.subscribers, rideID)
    }
}