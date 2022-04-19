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

package user_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/activity/streams"
	"github.com/superseriousbusiness/activity/streams/vocab"
	"github.com/superseriousbusiness/gotosocial/internal/api/s2s/user"
	"github.com/superseriousbusiness/gotosocial/testrig"
)

type RepliesGetTestSuite struct {
	UserStandardTestSuite
}

func (suite *RepliesGetTestSuite) TestGetReplies() {
	// the dereference we're gonna use
	derefRequests := testrig.NewTestDereferenceRequests(suite.testAccounts, suite.db)
	signedRequest := derefRequests["foss_satan_dereference_local_account_1_status_1_replies"]
	targetAccount := suite.testAccounts["local_account_1"]
	targetStatus := suite.testStatuses["local_account_1_status_1"]

	tc := testrig.NewTestTransportController(testrig.NewMockHTTPClient(nil), suite.db)
	federator := testrig.NewTestFederator(suite.db, tc, suite.storage, suite.mediaManager)
	emailSender := testrig.NewEmailSender("../../../../web/template/", nil)
	processor := testrig.NewTestProcessor(suite.db, suite.storage, federator, emailSender, suite.mediaManager)
	userModule := user.New(processor).(*user.Module)

	// setup request
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, targetStatus.URI+"/replies", nil) // the endpoint we're hitting
	ctx.Request.Header.Set("accept", "application/activity+json")
	ctx.Request.Header.Set("Signature", signedRequest.SignatureHeader)
	ctx.Request.Header.Set("Date", signedRequest.DateHeader)

	// we need to pass the context through signature check first to set appropriate values on it
	suite.securityModule.SignatureCheck(ctx)

	// normally the router would populate these params from the path values,
	// but because we're calling the function directly, we need to set them manually.
	ctx.Params = gin.Params{
		gin.Param{
			Key:   user.UsernameKey,
			Value: targetAccount.Username,
		},
		gin.Param{
			Key:   user.StatusIDKey,
			Value: targetStatus.ID,
		},
	}

	// trigger the function being tested
	userModule.StatusRepliesGETHandler(ctx)

	// check response
	suite.EqualValues(http.StatusOK, recorder.Code)

	result := recorder.Result()
	defer result.Body.Close()
	b, err := ioutil.ReadAll(result.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), `{"@context":"https://www.w3.org/ns/activitystreams","first":{"id":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?page=true","next":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?only_other_accounts=false\u0026page=true","partOf":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies","type":"CollectionPage"},"id":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies","type":"Collection"}`, string(b))

	// should be a Collection
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	assert.NoError(suite.T(), err)

	t, err := streams.ToType(context.Background(), m)
	assert.NoError(suite.T(), err)

	_, ok := t.(vocab.ActivityStreamsCollection)
	assert.True(suite.T(), ok)
}

func (suite *RepliesGetTestSuite) TestGetRepliesNext() {
	// the dereference we're gonna use
	derefRequests := testrig.NewTestDereferenceRequests(suite.testAccounts, suite.db)
	signedRequest := derefRequests["foss_satan_dereference_local_account_1_status_1_replies_next"]
	targetAccount := suite.testAccounts["local_account_1"]
	targetStatus := suite.testStatuses["local_account_1_status_1"]

	tc := testrig.NewTestTransportController(testrig.NewMockHTTPClient(nil), suite.db)
	federator := testrig.NewTestFederator(suite.db, tc, suite.storage, suite.mediaManager)
	emailSender := testrig.NewEmailSender("../../../../web/template/", nil)
	processor := testrig.NewTestProcessor(suite.db, suite.storage, federator, emailSender, suite.mediaManager)
	userModule := user.New(processor).(*user.Module)

	// setup request
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, targetStatus.URI+"/replies?only_other_accounts=false&page=true", nil) // the endpoint we're hitting
	ctx.Request.Header.Set("accept", "application/activity+json")
	ctx.Request.Header.Set("Signature", signedRequest.SignatureHeader)
	ctx.Request.Header.Set("Date", signedRequest.DateHeader)

	// we need to pass the context through signature check first to set appropriate values on it
	suite.securityModule.SignatureCheck(ctx)

	// normally the router would populate these params from the path values,
	// but because we're calling the function directly, we need to set them manually.
	ctx.Params = gin.Params{
		gin.Param{
			Key:   user.UsernameKey,
			Value: targetAccount.Username,
		},
		gin.Param{
			Key:   user.StatusIDKey,
			Value: targetStatus.ID,
		},
	}

	// trigger the function being tested
	userModule.StatusRepliesGETHandler(ctx)

	// check response
	suite.EqualValues(http.StatusOK, recorder.Code)

	result := recorder.Result()
	defer result.Body.Close()
	b, err := ioutil.ReadAll(result.Body)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), `{"@context":"https://www.w3.org/ns/activitystreams","id":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?page=true\u0026only_other_accounts=false","items":"http://localhost:8080/users/admin/statuses/01FF25D5Q0DH7CHD57CTRS6WK0","next":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?only_other_accounts=false\u0026page=true\u0026min_id=01FF25D5Q0DH7CHD57CTRS6WK0","partOf":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies","type":"CollectionPage"}`, string(b))

	// should be a Collection
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	assert.NoError(suite.T(), err)

	t, err := streams.ToType(context.Background(), m)
	assert.NoError(suite.T(), err)

	page, ok := t.(vocab.ActivityStreamsCollectionPage)
	assert.True(suite.T(), ok)

	assert.Equal(suite.T(), page.GetActivityStreamsItems().Len(), 1)
}

func (suite *RepliesGetTestSuite) TestGetRepliesLast() {
	// the dereference we're gonna use
	derefRequests := testrig.NewTestDereferenceRequests(suite.testAccounts, suite.db)
	signedRequest := derefRequests["foss_satan_dereference_local_account_1_status_1_replies_last"]
	targetAccount := suite.testAccounts["local_account_1"]
	targetStatus := suite.testStatuses["local_account_1_status_1"]

	tc := testrig.NewTestTransportController(testrig.NewMockHTTPClient(nil), suite.db)
	federator := testrig.NewTestFederator(suite.db, tc, suite.storage, suite.mediaManager)
	emailSender := testrig.NewEmailSender("../../../../web/template/", nil)
	processor := testrig.NewTestProcessor(suite.db, suite.storage, federator, emailSender, suite.mediaManager)
	userModule := user.New(processor).(*user.Module)

	// setup request
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, targetStatus.URI+"/replies?only_other_accounts=false&page=true&min_id=01FF25D5Q0DH7CHD57CTRS6WK0", nil) // the endpoint we're hitting
	ctx.Request.Header.Set("accept", "application/activity+json")
	ctx.Request.Header.Set("Signature", signedRequest.SignatureHeader)
	ctx.Request.Header.Set("Date", signedRequest.DateHeader)

	// we need to pass the context through signature check first to set appropriate values on it
	suite.securityModule.SignatureCheck(ctx)

	// normally the router would populate these params from the path values,
	// but because we're calling the function directly, we need to set them manually.
	ctx.Params = gin.Params{
		gin.Param{
			Key:   user.UsernameKey,
			Value: targetAccount.Username,
		},
		gin.Param{
			Key:   user.StatusIDKey,
			Value: targetStatus.ID,
		},
	}

	// trigger the function being tested
	userModule.StatusRepliesGETHandler(ctx)

	// check response
	suite.EqualValues(http.StatusOK, recorder.Code)

	result := recorder.Result()
	defer result.Body.Close()
	b, err := ioutil.ReadAll(result.Body)
	assert.NoError(suite.T(), err)

	fmt.Println(string(b))
	assert.Equal(suite.T(), `{"@context":"https://www.w3.org/ns/activitystreams","id":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?page=true\u0026only_other_accounts=false\u0026min_id=01FF25D5Q0DH7CHD57CTRS6WK0","items":[],"next":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?only_other_accounts=false\u0026page=true","partOf":"http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies","type":"CollectionPage"}`, string(b))

	// should be a Collection
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	assert.NoError(suite.T(), err)

	t, err := streams.ToType(context.Background(), m)
	assert.NoError(suite.T(), err)

	page, ok := t.(vocab.ActivityStreamsCollectionPage)
	assert.True(suite.T(), ok)

	assert.Equal(suite.T(), page.GetActivityStreamsItems().Len(), 0)
}

func TestRepliesGetTestSuite(t *testing.T) {
	suite.Run(t, new(RepliesGetTestSuite))
}
