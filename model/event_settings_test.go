// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventSettingsReadWrite(t *testing.T) {
	db := setupTestDb(t)
	defer db.Close()

	eventSettings, err := db.GetEventSettings()
	assert.Nil(t, err)
	assert.Equal(
		t,
		EventSettings{
			Id:                              1,
			Name:                            "Untitled Event",
			ElimType:                        "single",
			NumElimAlliances:                8,
			SelectionRound2Order:            "L",
			SelectionRound3Order:            "",
			TbaDownloadEnabled:              true,
			ApTeamChannel:                   157,
			ApAdminChannel:                  0,
			ApAdminWpaKey:                   "1234Five",
			WarmupDurationSec:               0,
			AutoDurationSec:                 15,
			PauseDurationSec:                3,
			TeleopDurationSec:               135,
			WarningRemainingDurationSec:     30,
			ChargeStationElectronicsEnabled: true,
			SustainabilityBonusLinkThresholdWithoutCoop: 7,
			SustainabilityBonusLinkThresholdWithCoop:    6,
			ActivationBonusPointThreshold:               26,
		},
		*eventSettings,
	)

	eventSettings.Name = "Chezy Champs"
	eventSettings.NumElimAlliances = 6
	eventSettings.SelectionRound2Order = "F"
	eventSettings.SelectionRound3Order = "L"
	err = db.UpdateEventSettings(eventSettings)
	assert.Nil(t, err)
	eventSettings2, err := db.GetEventSettings()
	assert.Nil(t, err)
	assert.Equal(t, eventSettings, eventSettings2)
}
