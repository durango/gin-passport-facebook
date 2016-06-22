package GinPassportFacebook

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"

  "github.com/gin-gonic/gin"
  "golang.org/x/oauth2"
)

const KeyNamespace string = "gin-passport-facebook-profile"

const ProfileUrl string = "https://graph.facebook.com/v2.2/me?fields=id,name,email,picture,first_name,last_name"

var passportOauth *oauth2.Config

func Routes(oauth *oauth2.Config, r *gin.RouterGroup) {
  passportOauth = oauth

  r.GET("/login", func(c *gin.Context) {
    Login(oauth, c)
  })
}

func Login(oauth *oauth2.Config, c *gin.Context) {
  url := oauth.AuthCodeURL("")
  c.Redirect(http.StatusFound, url)
}

func Middleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    getProfile(c)
  }
}

func GetProfile(c *gin.Context) (*Profile, error) {
  user, exists := c.Get(KeyNamespace)
  if !exists {
    return nil, errors.New("GinPassportFacebook namespace key doesn't exist")
  }

  return user.(*Profile), nil
}

func getProfile(c *gin.Context) {
  c.Request.ParseForm()

  opts := passportOauth
  code := c.Request.Form.Get("code")

  t, err := opts.Exchange(c, code)

  // most likely already authenticated / all errors will return `t` as nil
  if t == nil {
    c.Redirect(301, "/")
    return
  } else if err != nil {
    c.AbortWithError(http.StatusInternalServerError, err)
    return
  }

  client := opts.Client(c, t)

  resp, err := client.Get(ProfileUrl)
  if err != nil {
    c.AbortWithError(http.StatusInternalServerError, err)
    return
  }

  defer resp.Body.Close()
  contents, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    c.AbortWithError(http.StatusInternalServerError, err)
    return
  }

  var userInformation Profile
  err = json.Unmarshal(contents, &userInformation)
  if err != nil {
    c.AbortWithError(http.StatusInternalServerError, err)
    return
  }

  c.Set(KeyNamespace, &userInformation)
  c.Next()
}
