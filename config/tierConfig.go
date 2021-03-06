package config

import (
	"encoding/json"
)

var (
	//TierMap is the global map for the collection of settings, searchable by tier.
	TierMap map[int]TierSettings
	//TierMax explains itself.
	TierMax int
)

type TierSettings struct {
	TierLevel        int `json:"tier"`
	TierBaseSettings struct {
		Damage            int `json:"base_damage"`
		DamageVariance    int `json:"damage_variance"`
		Efficency         int `json:"base_efficency"`
		EfficencyVariance int `json:"efficency_varience"`
		CastTime          int `json:"base_cast_time"`
		CastTimeVariance  int `json:"cast_time_varience"`
		Cooldown          int `json:"base_cooldown"`
		CooldownVarience  int `json:"cooldown_varience"`
	} `json:"base_settings"`
	TierMaxSettings struct {
		MaxDamage    int `json:"max_damage"`
		MaxEfficency int `json:"max_efficency"`
		MinCasttime  int `json:"min_cast_time"`
		MinCooldown  int `json:"min_cooldown"`
	} `json:"max_settings"`
}

//LoadTierSettings helps load in our base values and variances. Max values are to be used to cap improvement.
func LoadTierSettings(tierJSON []byte) error {
	var loadTier []TierSettings

	err := json.Unmarshal(tierJSON, &loadTier)
	if err != nil {
		return err
	}

	tierMap := make(map[int]TierSettings)
	for _, value := range loadTier {
		if tierMap[value.TierLevel] == (TierSettings{}) {
			tierMap[value.TierLevel] = value
		}
	}

	TierMap = tierMap
	TierMax = getMaximumTier(TierMap)

	return nil
}

func getMaximumTier(tierMap map[int]TierSettings) int {
	max := 0
	for key := range tierMap {
		if key > max {
			max = key
		}
	}

	return max
}
