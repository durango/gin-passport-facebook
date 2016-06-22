# Gin Passport Facebook

Gin middleware for obtaining common Facebook profile information. I don't personally use all of the permission attributes but feel free to open an issue and I can take a look into it (or open a pull request).

## Example

```go
package main

import (
  "github.com/gin-gonic/gin"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/facebook"
  "github.com/durango/gin-passport-facebook"
)

func main() {
  opts := &oauth2.Config{
    RedirectURL:  "http://localhost:8080/auth/facebook/callback",
    ClientID:     "CLIENT_ID",
    ClientSecret: "CLIENT_SECRET",
    Scopes:       []string{"email", "public_profile"},
    Endpoint:     facebook.Endpoint,
  }

  router := gin.Default()

  auth := router.Group("/auth/facebook")

  // setup the configuration and mount the "/login" route
  GinPassportFacebook.Routes(opts, auth)

  // setup a customized callback url...
  auth.GET("/callback", GinPassportFacebook.Middleware(), func(c *gin.Context) {
    user, err := GinPassportFacebook.GetProfile(c)
    if user == nil || err != nil {
      c.AbortWithStatus(500)
      return
    }

    c.String(200, "Got it!")
  })

  router.Run()
}
```
