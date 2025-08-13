package discord

import (
	"sync"
	"strings"
)

type Handler func(Interaction) InteractionResponse

type Router struct {
	mu				sync.RWMutex
	cmd				map[string]Handler
	component		map[string]Handler
	modal			map[string]Handler
}

func NewRouter() *Router {
	return &Router {
		cmd: make(map[string]Handler),
		component: make(map[string]Handler),
		modal: make(map[string]Handler),
	}
}

func (r *Router) lookup(table map[string]Handler, key string) Handler {
	r.mu.Lock()
	h := table[key]
	r.mu.Unlock()
	return h
}

func (r *Router) lookupByPrefix(table map[string]Handler, key string) Handler {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for prefix, h := range table {
		if strings.HasPrefix(key, prefix) {
			return h
		}
	}
	return nil
}

func (r *Router) RegisterCommand(name string, h Handler)        { r.mu.Lock(); r.cmd[name]=h; r.mu.Unlock() }
func (r *Router) RegisterComponent(prefix string, h Handler)    { r.mu.Lock(); r.component[prefix]=h; r.mu.Unlock() }
func (r *Router) RegisterModal(prefix string, h Handler)        { r.mu.Lock(); r.modal[prefix]=h; r.mu.Unlock() }


func (r *Router) Dispatch(i Interaction) InteractionResponse {

	switch i.Type {
	
		case InteractionTypePing:
			return RespPong()
		case InteractionTypeApplicationCommand:
			cmd := i.Data.Name
			if h := r.lookup(r.cmd, cmd); h != nil {
				return h(i)
			} 
			return RespMessage("Unknown command", false)
		
		case InteractionTypeMessageComponent:
			id := i.MessageComponentData.CustomID
			if h := r.lookupByPrefix(r.component, id); h != nil {
				return h(i)
			}
			return RespMessage("Unknown component", true)

		case InteractionTypeModalSubmit:
        	id := i.ModalSubmitData.CustomID
        	if h := r.lookupByPrefix(r.modal, id); h != nil {
            	return h(i)
        	}
        	return RespMessage("Unknown modal", true)

    	default:
        	return RespMessage("Unsupported interaction", true)
	}

}
