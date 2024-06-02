package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mpodriezov/shuzzles/src/data"
	"github.com/mpodriezov/shuzzles/src/utils"
)

type UsersAdminPage struct {
	Users []*data.UserModel
}

type UserUpdateData struct {
	Id       uint64
	Username string
	Email    string
	Bio      string
	Role     byte
}

type UserAdminPage struct {
	User      *UserUpdateData
	Errors    map[string]string
	UserRoles map[string]bool
}

func CreatePage(errors map[string]string, usr *UserUpdateData) *UserAdminPage {
	return &UserAdminPage{
		User:   usr,
		Errors: errors,
		UserRoles: map[string]bool{
			"user":  usr.Role&data.Role_User != 0,
			"admin": usr.Role&data.Role_Admin != 0,
		},
	}
}

func (u *UserUpdateData) Validate() map[string]string {
	errors := make(map[string]string)
	if len(u.Username) < 3 {
		errors["username"] = "Username must be at least 3 characters"
	}
	if !utils.IsEmailValid(u.Email) {
		errors["email"] = "Invalid email address"
	}
	return errors
}

// [GET] admin/user
func HandleShowUsersAdminView(c echo.Context) error {
	dal := c.Get("dal").(*data.Dal)
	page := UsersAdminPage{
		Users: dal.ListUsers(),
	}
	return c.Render(http.StatusOK, "admin/users.html", page)
}

// [GET] admin/user/:userId
func HandleShowUserAdminView(c echo.Context) error {
	dal := c.Get("dal").(*data.Dal)
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	user := dal.GetUserById(userId)

	c.Logger().Infof("bio: %s", user.Bio)

	page := CreatePage(nil, &UserUpdateData{
		Id:       user.Id,
		Email:    user.Email,
		Bio:      user.Bio,
		Role:     user.Role,
		Username: user.Username,
	})

	return c.Render(http.StatusOK, "admin/user_edit.html", page)
}

func HandlerUserAdminUpdate(c echo.Context) error {

	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	dal := c.Get("dal").(*data.Dal)
	user := dal.GetUserById(userId)

	if user == nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	c.Request().ParseForm()
	userRole, err := utils.ParseBitField(c.Request().Form["role"])

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid role")
	}

	form := c.Request().Form

	usr := UserUpdateData{
		Id:       userId,
		Username: strings.TrimSpace(form.Get("username")),
		Email:    strings.TrimSpace(form.Get("email")),
		Bio:      strings.TrimSpace(form.Get("bio")),
		Role:     userRole,
	}

	errors := usr.Validate()
	if len(errors) != 0 {
		page := CreatePage(errors, &usr)
		return c.Render(http.StatusUnprocessableEntity, "admin/user_form.html", page)
	}

	if user.Email != usr.Email && dal.DoesUserExistWithEmail(usr.Email) {
		errors["email"] = "Email is already taken"
		page := CreatePage(errors, &usr)
		return c.Render(http.StatusUnprocessableEntity, "admin/user_form.html", page)
	}

	if user.Username != usr.Username && dal.DoesUserExistWithUsername(usr.Username) {
		errors["username"] = "Username is already taken"
		page := UserAdminPage{
			User:   &usr,
			Errors: errors,
			UserRoles: map[string]bool{
				"user":  usr.Role&data.Role_User != 0,
				"admin": usr.Role&data.Role_Admin != 0,
			},
		}
		return c.Render(http.StatusUnprocessableEntity, "admin/user_form.html", page)
	}

	dal.UpdateUser(usr.Id, usr.Username, usr.Email, usr.Bio, usr.Role)

	c.Response().Header().Set("HX-Redirect", "/admin/user")
	return c.NoContent(http.StatusFound)
}

func HandleAdminDeleteUserShow(c echo.Context) error {

	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	dal := c.Get("dal").(*data.Dal)
	user := dal.GetUserById(userId)

	if user == nil {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}

	return c.Render(http.StatusOK, "admin/user_delete.html", user)
}

func HandleAdminDeleteUser(c echo.Context) error {
	dal := c.Get("dal").(*data.Dal)
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	user := dal.GetUserById(userId)

	if user == nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(http.StatusNotFound)
	}

	dal.DeleteUser(userId)

	c.Response().Header().Set("HX-Redirect", "/admin/user")
	return c.NoContent(http.StatusFound)
}
