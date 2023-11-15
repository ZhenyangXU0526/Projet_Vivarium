package enums

// Meteo represents the different weather conditions in the simulation.
type Meteo int

// Enumeration of Meteo
const (
	Pluie       Meteo = iota // Rain
	Brouillard               // Fog
	SaisonSeche              // Dry Season
	Incendie                 // Fire
	Tonnerre                 // Thunder
	// ... add other weather conditions as needed
)

// String returns the string representation of the Meteo
func (m Meteo) String() string {
	return [...]string{"Pluie", "Brouillard", "Saison Seche", "Incendie", "Tonnerre"}[m]
}