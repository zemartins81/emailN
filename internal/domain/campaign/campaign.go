package campaign

import "time"

type Contact struct {
	Name  string
	Email string
}

type Campaign struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedOn time.Time `json:"created_on"`
	Content   string    `json:"content,omitempty"`
	Contacts  []Contact `json:"contacts,omitempty"`
}

// NewCampaign atua como um construtor para a estrutura de Campanha (Campaign).
func NewCampaign(name string, content string, emails []string) *Campaign {

	// Criando um slice de contatos do mesmo tamanho do slice de emails passado como parâmetro.
	contacts := make([]Contact, len(emails))

	// Looping através de cada email, criando um contato para cada um.
	// O nome do contato é iniciado como uma string vazia e o email é o que passado no slice de emails.
	for index, email := range emails {
		contacts[index].Name = ""
		contacts[index].Email = email
	}

	// Retorna uma nova instância da estrutura de Campanha (Campaign), preenchida com os dados fornecidos.
	// A ID é configurada como 1, o nome e o conteúdo são os que foram passados como parâmetros,
	// a hora de criação é a hora atual e os contatos são os que acabamos de criar.
	// Note que a função retorna um ponteiro para a estrutura de Campaign.
	return &Campaign{
		ID:        1,
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
