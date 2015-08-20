# Gin Passport Facebook

Gin middleware for obtaining common Facebook profile information. I don't personally use all of the permission attributes but feel free to open an issue and I can take a look into it (or open a pull request).

## Example

```go
import (
  "github.com/gin-gonic/gin"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"

  "github.com/durango/gin-passport-facebook"
)

func main() {
  opts, _ := oauth2.New(
    oauth2.Client("ClientId", "YourSecretKey"),
    oauth2.RedirectURL("Your redirect URL"),
    oauth2.Scope("public_profile", "email"),
    oauth2.Endpoint(
      "https://www.facebook.com/dialog/oauth",
      "https://graph.facebook.com/oauth/access_token",
    ),
  )

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
}
```
