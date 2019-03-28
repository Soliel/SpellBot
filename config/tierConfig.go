package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	//TierMap is the global map for the collection of settings, searchable by tier.
	TierMap map[int]tierSettings
	//TierMax explains itself.
	TierMax int
)

type tierSettings struct {
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
		MaxEfficency int `json:"max_Efficency"`
		MinCasttime  int `json:"min_cast_time"`
		MinCooldown  int `json:"min_cool_down"`
	} `json:"max_settings"`
}

func loadTierSettings(tierPath string) (map[int]tierSettings, error) {
	var loadTier []tierSettings
	loadBody, err := ioutil.ReadFile(tierPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(loadBody, &loadTier)
	if err != nil {
		return nil, err
	}

	tierMap := make(map[int]tierSettings)
	for _, value := range loadTier {
		if tierMap[value.TierLevel] != (tierSettings{}) {
			tierMap[value.TierLevel] = value
		}
	}

	return tierMap, nil
}

func getMaximumTier(tierMap map[int]tierSettings) int {
	max := 0
	for key := range tierMap {
		if key > max {
			max = key
		}
	}

	return max
}
