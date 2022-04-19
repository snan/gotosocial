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

package trans_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
	"github.com/superseriousbusiness/gotosocial/internal/trans"
	"github.com/superseriousbusiness/gotosocial/testrig"
)

type ImportMinimalTestSuite struct {
	TransTestSuite
}

func (suite *ImportMinimalTestSuite) TestImportMinimalOK() {
	ctx := context.Background()

	// use a temporary file path
	tempFilePath := fmt.Sprintf("%s/%s", suite.T().TempDir(), uuid.NewString())

	// export to the tempFilePath
	exporter := trans.NewExporter(suite.db)
	err := exporter.ExportMinimal(ctx, tempFilePath)
	suite.NoError(err)

	// we should have some bytes in that file now
	b, err := os.ReadFile(tempFilePath)
	suite.NoError(err)
	suite.NotEmpty(b)
	fmt.Println(string(b))

	// create a new database with just the tables created, no entries
	testrig.StandardDBTeardown(suite.db)
	newDB := testrig.NewTestDB()

	importer := trans.NewImporter(newDB)
	err = importer.Import(ctx, tempFilePath)
	suite.NoError(err)

	// we should have some accounts in the database
	accounts := []*gtsmodel.Account{}
	err = newDB.GetAll(ctx, &accounts)
	suite.NoError(err)
	suite.NotEmpty(accounts)

	// we should have some blocks in the database
	blocks := []*gtsmodel.Block{}
	err = newDB.GetAll(ctx, &blocks)
	suite.NoError(err)
	suite.NotEmpty(blocks)

	// we should have some follows in the database
	follows := []*gtsmodel.Follow{}
	err = newDB.GetAll(ctx, &follows)
	suite.NoError(err)
	suite.NotEmpty(follows)

	// we should have some domain blocks in the database
	domainBlocks := []*gtsmodel.DomainBlock{}
	err = newDB.GetAll(ctx, &domainBlocks)
	suite.NoError(err)
	suite.NotEmpty(domainBlocks)
}

func TestImportMinimalTestSuite(t *testing.T) {
	suite.Run(t, &ImportMinimalTestSuite{})
}
