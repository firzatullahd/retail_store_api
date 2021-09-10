package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"retailStore/config"
	"retailStore/lib/db"
	"retailStore/lib/seeders"
	"retailStore/models"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

func createModelCartLists() []models.ShoppingCartList {
	cartLists := make([]models.ShoppingCartList, 1)
	cartLists[0] = models.ShoppingCartList{
		ItemID:   1,
		Quantity: 3,
	}
	return cartLists
}

func TestPostItemToShoppingCart(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	seeders.ItemSeed()
	userModel, err := seeders.UserSeed()

	config.DB.Create(&userModel)
	userModel.ShoppingCart.UserID = userModel.ID
	config.DB.Create(&userModel.ShoppingCart)

	user := models.User{
		Username: userModel.Username,
		Password: userModel.Password,
	}
	db.LoginUser(&user)
	cartLists := createModelCartLists()
	cartLists[0].ShoppingCartID = userModel.ShoppingCart.ID
	//items := createModelItem()
	itemsJSON, _ := json.Marshal(cartLists[0])

	e := echo.New()
	h := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     jwt.MapClaims{},
		SigningKey: []byte("legal"),
	})(PostItemToShoppingCartController)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(itemsJSON)))
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+user.Token)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//Assertions
	// test insert item to shopping cart
	if assert.NoError(t, h(c)) && err == nil {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Condition(t, func() bool {
			var dat models.ShoppingCartAPI
			var b []byte = rec.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.ShoppingCartList[0].ItemID == cartLists[0].ItemID && dat.Data.ShoppingCartList[0].Quantity == cartLists[0].Quantity && dat.Data.ShoppingCartList[0].ShoppingCartID == cartLists[0].ShoppingCartID {
				return true
			}

			return false
		}, rec.Body.String())
	}

	//test getshoppingcart
	middlewareHandler := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     jwt.MapClaims{},
		SigningKey: []byte("legal"),
	})(GetShoppingCartController)
	req2 := httptest.NewRequest(http.MethodPost, "/", nil)
	req2.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+user.Token)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	if assert.NoError(t, middlewareHandler(c2)) && err == nil {
		assert.Equal(t, http.StatusOK, rec2.Code)
		assert.Condition(t, func() bool {
			var dat models.ShoppingCartAPI
			var b []byte = rec2.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.ShoppingCartList[0].ItemID == cartLists[0].ItemID && dat.Data.ShoppingCartList[0].Quantity == cartLists[0].Quantity && dat.Data.ShoppingCartList[0].ShoppingCartID == cartLists[0].ShoppingCartID && dat.Data.ID == userModel.ShoppingCart.ID {
				return true
			}

			return false
		}, rec2.Body.String())
	}
	//checkout shoppingcart
	modelOrder := models.Order{
		CourierID:        1,
		PaymentServiceID: 2,
		AddressID:        1,
	}
	jsonModelOrder, _ := json.Marshal(modelOrder)
	middlewareHandler3 := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     jwt.MapClaims{},
		SigningKey: []byte("legal"),
	})(ShoppingCartCheckoutController)
	req4 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonModelOrder)))
	req4.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+user.Token)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)

	if assert.NoError(t, middlewareHandler3(c4)) && err == nil {
		assert.Equal(t, http.StatusOK, rec4.Code)
		assert.Condition(t, func() bool {
			var dat models.ResponseOrder
			var b []byte = rec4.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.AddressID == modelOrder.AddressID && dat.Data.CourierID == modelOrder.CourierID && dat.Data.PaymentServiceID == modelOrder.PaymentServiceID {
				return true
			}

			return false
		}, rec4.Body.String())
	}

	//test delete item from shoppingcart
	JsonArrCartList, _ := json.Marshal(cartLists)
	middlewareHandler2 := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     jwt.MapClaims{},
		SigningKey: []byte("legal"),
	})(DeleteItemFromShoppingCartController)
	req3 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(JsonArrCartList)))
	req3.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+user.Token)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)

	if assert.NoError(t, middlewareHandler2(c3)) && err == nil {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Condition(t, func() bool {
			var dat models.ShoppingCartAPI
			var b []byte = rec3.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if len(dat.Data.ShoppingCartList) == 0 {
				return true
			}

			return false
		}, rec3.Body.String())
	}

}