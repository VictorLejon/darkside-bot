package discord

type InteractionType int

const (
	InteractionTypePing					InteractionType = 1
	InteractionTypeApplicationCommand	InteractionType = 2
	InteractionTypeMessageComponent		InteractionType = 3
	InteractionTypeChannelMessage		InteractionType = 4
	InteractionTypeModalSubmit			InteractionType = 5
)

const (
    ApplicationCommandTypeChatInput    = 1
    ApplicationCommandTypeUser         = 2
    ApplicationCommandTypeMessage      = 3
    ApplicationCommandTypeAutocomplete = 4
)

const EphemeralMessageFlag = 1 << 6 // 64

type Interaction struct {
    Type InteractionType `json:"type"` 
    Data struct {
		Type    int					   `json:"type,omitempty"`
        Name    string                 `json:"name"`
        Options []InteractionOption    `json:"options,omitempty"`
    } `json:"data,omitempty"`
    Token string `json:"token"`

	MessageComponentData struct {
		CustomID string `json:"custom_id"`
	} `json:"message_component_data,omitempty"`


	ModalSubmitData struct {
		CustomID string `json:"custom_id"`
	} `json:"modal_submit_data,omitempty"`
}

type InteractionOption struct {
    Name  string      `json:"name"`
    Value interface{} `json:"value"`
	Focused bool	  `json:"focused,omitempty"`
}

type InteractionResponse struct {
	Type InteractionType `json:"type"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

type InteractionCallbackData struct {
    Content string `json:"content,omitempty"`
    Flags   int    `json:"flags,omitempty"`
}




