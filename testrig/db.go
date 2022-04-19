/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package testrig

import (
	"context"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/superseriousbusiness/gotosocial/internal/config"
	"github.com/superseriousbusiness/gotosocial/internal/db"
	"github.com/superseriousbusiness/gotosocial/internal/db/bundb"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
)

var testModels = []interface{}{
	&gtsmodel.Account{},
	&gtsmodel.AdminAccountAction{},
	&gtsmodel.Application{},
	&gtsmodel.Block{},
	&gtsmodel.DomainBlock{},
	&gtsmodel.EmailDomainBlock{},
	&gtsmodel.Follow{},
	&gtsmodel.FollowRequest{},
	&gtsmodel.MediaAttachment{},
	&gtsmodel.Mention{},
	&gtsmodel.Status{},
	&gtsmodel.StatusToEmoji{},
	&gtsmodel.StatusToTag{},
	&gtsmodel.StatusFave{},
	&gtsmodel.StatusBookmark{},
	&gtsmodel.StatusMute{},
	&gtsmodel.Tag{},
	&gtsmodel.User{},
	&gtsmodel.Emoji{},
	&gtsmodel.Instance{},
	&gtsmodel.Notification{},
	&gtsmodel.RouterSession{},
	&gtsmodel.Token{},
	&gtsmodel.Client{},
}

// NewTestDB returns a new initialized, empty database for testing.
//
// If the environment variable GTS_DB_ADDRESS is set, it will take that
// value as the database address instead.
//
// If the environment variable GTS_DB_TYPE is set, it will take that
// value as the database type instead.
//
// If the environment variable GTS_DB_PORT is set, it will take that
// value as the port instead.
func NewTestDB() db.DB {
	if alternateAddress := os.Getenv("GTS_DB_ADDRESS"); alternateAddress != "" {
		viper.Set(config.Keys.DbAddress, alternateAddress)
	}

	if alternateDBType := os.Getenv("GTS_DB_TYPE"); alternateDBType != "" {
		viper.Set(config.Keys.DbType, alternateDBType)
	}

	if alternateDBPort := os.Getenv("GTS_DB_PORT"); alternateDBPort != "" {
		port, err := strconv.ParseInt(alternateDBPort, 10, 64)
		if err != nil {
			panic(err)
		}
		viper.Set(config.Keys.DbPort, port)
	}

	testDB, err := bundb.NewBunDBService(context.Background())
	if err != nil {
		logrus.Panic(err)
	}
	return testDB
}

// StandardDBSetup populates a given db with all the necessary tables/models for perfoming tests.
//
// The accounts parameter is provided in case the db should be populated with a certain set of accounts.
// If accounts is nil, then the standard test accounts will be used.
//
// When testing http signatures, you should pass into this function the same accounts map that you generated
// signatures with, otherwise this function will randomly generate new keys for accounts and signature
// verification will fail.
func StandardDBSetup(db db.DB, accounts map[string]*gtsmodel.Account) {
	if db == nil {
		logrus.Panic("db setup: db was nil")
	}

	ctx := context.Background()

	for _, v := range NewTestTokens() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestClients() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestApplications() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestBlocks() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestDomainBlocks() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestUsers() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	if accounts == nil {
		for _, v := range NewTestAccounts() {
			if err := db.Put(ctx, v); err != nil {
				logrus.Panic(err)
			}
		}
	} else {
		for _, v := range accounts {
			if err := db.Put(ctx, v); err != nil {
				logrus.Panic(err)
			}
		}
	}

	for _, v := range NewTestAttachments() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestStatuses() {
		if err := db.PutStatus(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestEmojis() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestTags() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestMentions() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestFaves() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestFollows() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	for _, v := range NewTestNotifications() {
		if err := db.Put(ctx, v); err != nil {
			logrus.Panic(err)
		}
	}

	if err := db.CreateInstanceAccount(ctx); err != nil {
		logrus.Panic(err)
	}

	if err := db.CreateInstanceInstance(ctx); err != nil {
		logrus.Panic(err)
	}

	logrus.Debug("testing db setup complete")
}

// StandardDBTeardown drops all the standard testing tables/models from the database to ensure it's clean for the next test.
func StandardDBTeardown(db db.DB) {
	ctx := context.Background()
	if db == nil {
		logrus.Panic("db teardown: db was nil")
	}

	for _, m := range testModels {
		if err := db.DropTable(ctx, m); err != nil {
			logrus.Panic(err)
		}
	}

	if err := db.DropMigrationTables(ctx); err != nil {
		logrus.Panic(err)
	}
}
