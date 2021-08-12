package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/api/v1/home").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
