package zkbq

type KillPackage struct {
	Payload struct {
		KillID   int     `json:"killID"`
		Killmail EveKill `json:"killmail"`
	} `json:"package"`
}

type KillItem struct {
	Flag              int         `json:"flag"`
	ItemType          CommonEntry `json:"itemType"`
	NumberOfDestroyed int         `json:"quantityDestroyed"`
	Singleton         int         `json:"singleton"`
}

type CommonEntry struct {
	Link string `json:"href"`
	Icon struct {
		Link string `json:"href"`
	} `json:"icon"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type KillmailAttacker struct {
	Alliance    CommonEntry `json:"alliance"`
	Character   CommonEntry `json:"character"`
	Corporation CommonEntry `json:"corporation"`

	DamageDone     int     `json:"damageDone"`
	FinalBlow      bool    `json:"finalBlow"`
	SecurityStatus float64 `json:"securityStatus"`

	ShipType   CommonEntry `json:"shipType"`
	WeaponType CommonEntry `json:"weaponType"`
}

type KillmailVictim struct {
	Alliance    CommonEntry `json:"alliance"`
	Character   CommonEntry `json:"character"`
	Corporation CommonEntry `json:"corporation"`

	DamageTake int        `json:"damageTaken"`
	Items      []KillItem `json:"items"`
	Position   struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"position"`
	ShipType CommonEntry `json:"shipType"`
}

type EveKill struct {
	NumAttackers int                `json:"attackerCount"`
	KillID       int                `json:"killID"`
	KillTime     string             `json:"killTime"`
	Attackers    []KillmailAttacker `json:"attackers"`
	Victim       KillmailVictim     `json:"victim"`
	War          struct {
		Link string `json:"href"`
		ID   int    `json:"id"`
	} `json:"war,omitempty"`
}
