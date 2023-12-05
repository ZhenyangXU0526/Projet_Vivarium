package organisme

import (
	"fmt"
	"vivarium/enums"
	"vivarium/terrain"
)

// Organisme defines the interface that all organisms must implement.
type Organisme interface {
	// SeDeplacer(t *terrain.Terrain)
	Vieillir(t *terrain.Terrain)
	Mourir(t *terrain.Terrain)
	CheckEtat(t *terrain.Terrain) Organisme

	//================================================
	GetID() int
	GetAge() int
	GetPosX() int
	GetPosY() int
	GetRayon() int
	GetEspece() enums.MyEspece
	SetID(newID int)
	GetEtat() bool
}

// BaseOrganisme provides the base implementation of the Organisme interface.
type BaseOrganisme struct {
	OrganismeID          int
	Age                  int
	PositionX            int
	PositionY            int
	Rayon                int
	Espece               enums.MyEspece
	AgeRate              int  // 衰老速度
	MaxAge               int  // 最大寿命
	GrownUpAge           int  // 成年年龄
	TooOldToReproduceAge int  // 老到生不动的年龄
	NbProgeniture        int  // 一次能够生出的后代数量
	Busy                 bool // 是否在进行动作 （相当于一个针对个体生物的锁）
	IsInsecte            bool // 用于判断该生物是否是昆虫
	IsDying              bool // 用于告知ebiten的sprite该昆虫要执行死亡渲染动画
}

// NewBaseOrganisme creates a new BaseOrganisme instance.
func NewBaseOrganisme(id int, age, posX, posY, rayon int, espece enums.MyEspece, ageRate int, maxAge int, grownUpAge, tooOldToRep, nbProgeniture int, isInsecte bool) *BaseOrganisme {
	return &BaseOrganisme{
		OrganismeID:          id,
		Age:                  age,
		PositionX:            posX,
		PositionY:            posY,
		Rayon:                rayon,
		Espece:               espece,
		AgeRate:              ageRate,
		MaxAge:               maxAge,
		GrownUpAge:           grownUpAge,
		TooOldToReproduceAge: tooOldToRep,
		NbProgeniture:        nbProgeniture,
		Busy:                 false,
		IsInsecte:            isInsecte,
		IsDying:              false,
	}
}

// Implementations of the Organisme interface's methods follow:

// // ID returns the organism's ID.
// func (bo *BaseOrganisme) ID() int {
// 	return bo.OrganismeID
// }

// Vieillir simulates the organism aging.
func (bo *BaseOrganisme) Vieillir(t *terrain.Terrain) {
	// 检查是否Busy
	// if bo.Busy {
	// 	fmt.Println("Organisme", bo.GetID(), "is busy, cannot age [他妈的肾上腺素] [老子就是虫王杰森斯坦森]") //他妈的肾上腺素
	// 	return
	// }
	// hotfix-1124: 不用加锁：
	//如果判断可以bang，那个瞬间可以bang就行了，判断过了就相当于成功feconde了
	//如果干到一半被火烧死了，属于庞贝事件，也很合理
	//不可能事件：bang到一半老死了。因为设置了ToooOldToReproduceAge，和MaxAge有一定的差距。

	bo.Busy = true
	defer func() { bo.Busy = false }()

	bo.Age += bo.AgeRate
	if bo.Age >= bo.MaxAge {
		// Reaching the maximum lifespan, the organism should die
		fmt.Println(bo.GetID(), "died because of old age")
		bo.Mourir(t)
	}
}

// Mourir simulates the organism dying.
func (bo *BaseOrganisme) Mourir(t *terrain.Terrain) {
	// Implementation of dying
	// I'm not quite sure if we should clear the memory of the organism (maybe we should)
	/*...*/
	// Remove the organism from the Terrain
	t.RemoveOrganism(bo.OrganismeID, bo.PositionX, bo.PositionY)
	bo.IsDying = true
}

func (bo *BaseOrganisme) GetID() int {
	return bo.OrganismeID
}
func (bo *BaseOrganisme) GetAge() int {
	return bo.Age
}
func (bo *BaseOrganisme) GetPosX() int {
	return bo.PositionX
}
func (bo *BaseOrganisme) GetPosY() int {
	return bo.PositionY
}
func (bo *BaseOrganisme) GetRayon() int {
	return bo.Rayon
}
func (bo *BaseOrganisme) GetEspece() enums.MyEspece {
	return bo.Espece
}
func (bo *BaseOrganisme) GetEtat() bool {
	return bo.IsDying
}
func (bo *BaseOrganisme) SetID(newID int) {
	bo.OrganismeID = newID
}
