// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model and datastore read/write methods for event-level configuration.

package model

import "github.com/Team254/cheesy-arena/game"

type EventSettings struct {
	Id                                          int `db:"id"`
	Name                                        string
	ElimType                                    string
	NumElimAlliances                            int
	SelectionRound2Order                        string
	SelectionRound3Order                        string
	TbaDownloadEnabled                          bool
	TbaPublishingEnabled                        bool
	TbaEventCode                                string
	TbaSecretId                                 string
	TbaSecret                                   string
	NetworkSecurityEnabled                      bool
	UseMultiConnectionAPConfiguration           bool
	ApAddress                                   string
	ApUsername                                  string
	ApPassword                                  string
	ApTeamChannel                               int
	ApAdminChannel                              int
	ApAdminWpaKey                               string
	Ap2Address                                  string
	Ap2Username                                 string
	Ap2Password                                 string
	Ap2TeamChannel                              int
	SwitchAddress                               string
	SwitchPassword                              string
	PlcAddress                                  string
	AdminPassword                               string
	WarmupDurationSec                           int
	AutoDurationSec                             int
	PauseDurationSec                            int
	TeleopDurationSec                           int
	WarningRemainingDurationSec                 int
	ChargeStationElectronicsEnabled             bool
	SustainabilityBonusLinkThresholdWithoutCoop int
	SustainabilityBonusLinkThresholdWithCoop    int
	ActivationBonusPointThreshold               int
}

func (database *Database) GetEventSettings() (*EventSettings, error) {
	allEventSettings, err := database.eventSettingsTable.getAll()
	if err != nil {
		return nil, err
	}
	if len(allEventSettings) == 1 {
		return &allEventSettings[0], nil
	}

	// Database record doesn't exist yet; create it now.
	eventSettings := EventSettings{
		Name:                            "Untitled Event",
		ElimType:                        "single",
		NumElimAlliances:                8,
		SelectionRound2Order:            "L",
		SelectionRound3Order:            "",
		TbaDownloadEnabled:              true,
		ApTeamChannel:                   157,
		ApAdminChannel:                  0,
		ApAdminWpaKey:                   "1234Five",
		Ap2TeamChannel:                  0,
		WarmupDurationSec:               game.MatchTiming.WarmupDurationSec,
		AutoDurationSec:                 game.MatchTiming.AutoDurationSec,
		PauseDurationSec:                game.MatchTiming.PauseDurationSec,
		TeleopDurationSec:               game.MatchTiming.TeleopDurationSec,
		WarningRemainingDurationSec:     game.MatchTiming.WarningRemainingDurationSec,
		ChargeStationElectronicsEnabled: true,
		SustainabilityBonusLinkThresholdWithoutCoop: game.SustainabilityBonusLinkThresholdWithoutCoop,
		SustainabilityBonusLinkThresholdWithCoop:    game.SustainabilityBonusLinkThresholdWithCoop,
		ActivationBonusPointThreshold:               game.ActivationBonusPointThreshold,
	}

	if err := database.eventSettingsTable.create(&eventSettings); err != nil {
		return nil, err
	}
	return &eventSettings, nil
}

func (database *Database) UpdateEventSettings(eventSettings *EventSettings) error {
	return database.eventSettingsTable.update(eventSettings)
}
