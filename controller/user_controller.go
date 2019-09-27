package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"park/project/2019_2_Premium-Harbor/component"
	"park/project/2019_2_Premium-Harbor/storage"
	"time"
)

const (
	SessionIDCookieName   = "session_id"
	SessionIDCookieExpire = 10 * time.Hour
)

type UserController struct {
	Controller
	userComponent *component.UserComponent
}

func NewUserController() *UserController {
	return &UserController{
		userComponent: component.NewUserComponent(),
	}
}

type UserToList struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (c UserController) HandleUserList(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	users := c.userComponent.GetAllUsers()
	usersToList := c.convertUsersForList(users)
	c.writeOkWithBody(w, map[string]interface{}{
		"users": usersToList,
	})
}

func (c UserController) convertUsersForList(users []storage.User) []UserToList {
	usersToList := make([]UserToList, 0, len(users))
	for _, user := range users {
		usersToList = append(usersToList, UserToList{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}
	return usersToList
}

type UserToUpdate struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c UserController) HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	decoder := json.NewDecoder(r.Body)
	user := new(UserToUpdate)
	err := decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.userComponent.UpdateUser(user.ID, user.Password, user.Name)
	c.writeOk(w)
}

type UserToRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c UserController) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	decoder := json.NewDecoder(r.Body)
	user := new(UserToRegister)
	err := decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	err = c.userComponent.Register(user.Email, user.Password, user.Name)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOk(w)
}

type UserToLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c UserController) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	decoder := json.NewDecoder(r.Body)
	user := new(UserToLogin)
	err := decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	var sessionID string
	sessionID, err = c.userComponent.Login(user.Email, user.Password)
	if err != nil {
		c.writeError(w, err)
		return
	}
	sessionIDCookie := &http.Cookie{
		Name:    SessionIDCookieName,
		Value:   sessionID,
		Expires: time.Now().Add(SessionIDCookieExpire),
	}
	http.SetCookie(w, sessionIDCookie)
	c.writeOk(w)
}

func (c UserController) HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	sessionIDCookie, err := r.Cookie(SessionIDCookieName)
	if err != nil {
		c.writeError(w, fmt.Errorf("no session cookie"))
		return
	}
	c.deleteCookie(w, sessionIDCookie)
	err = c.userComponent.Logout(sessionIDCookie.Value)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOk(w)
}
