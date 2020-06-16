package combat

import "github.com/utrescu/CombatCavallers.Go/lluitador"

// IResultat retorna el resultat del combat
type IResultat interface {
	GetNom() string
	GetVida() int
	EsKo() bool
	EstaEliminat() bool
	Elimina()
}

// ICombatent definex el que necessita el Ring
type ICombatent interface {
	GetLluitador() lluitador.ILluitador
	TreuVida(vida int) int
}

// Resultat és la implementació de les dues interfícies
type Resultat struct {
	lluitador lluitador.ILluitador
	vida      int
	eliminat  bool
}

// CreateResultat crea un objecte de tipus resultat
func CreateResultat(lluitador1 lluitador.ILluitador, vida int) Resultat {
	return Resultat{lluitador1, vida, false}
}

// GetNom retorna el nom del lluitador
func (r Resultat) GetNom() string {
	return r.lluitador.GetNom()
}

// GetVida retorna la vida que li queda al lluitador
func (r Resultat) GetVida() int {
	return r.vida
}

// EsKo informa de si el lluitador ha quedat Ko
func (r Resultat) EsKo() bool {
	return r.vida == 0
}

// EstaEliminat informa de si el lluitador ha quedat eliminat
func (r Resultat) EstaEliminat() bool {
	return r.eliminat
}

// Elimina elimina el lluitador
func (r *Resultat) Elimina() {
	r.eliminat = true
}

// GetLluitador retorna el lluitador
func (r Resultat) GetLluitador() lluitador.ILluitador {
	return r.lluitador
}

// TreuVida és el mètode per treure vida al lluitador
func (r *Resultat) TreuVida(vida int) int {
	r.vida = r.vida - vida
	return r.vida
}
