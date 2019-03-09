package data

//SpellType Enum
const (
	Attack  SpellType = 0
	Defense SpellType = 1
	Dodge   SpellType = 2

	BuffEfficency  SpellType = 3 //Lowers Manacost for spells
	BuffChanting   SpellType = 4 //Lowers cast time for spells
	BuffDamage     SpellType = 5 //Increases Damage
	BuffEffects    SpellType = 6 //Increases Effect Efficacy
	BuffSpellSkill SpellType = 7 //Increases Chance of successful casting.

	DeBuffEfficency  SpellType = 8
	DeBuffChanting   SpellType = 9
	DeBuffEffects    SpellType = 10
	DeBuffSpellSkill SpellType = 11
	DeBuffDamage     SpellType = 12
)

//SpellType is a wrapper for the SpellType Enum.
type SpellType int
