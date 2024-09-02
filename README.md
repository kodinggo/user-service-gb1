# user-service-gb1

User service kodinggo golang bootcamp 1

## Installation
```
make migrate-up
make tidy
make run
```

## JWT Middleware gRPC Implementation
1. Clone repo auth-service ini, dan pastikan berjalan dengan baik.
        req := &pb.ValidateTokenRequest{Token: token} 
2. Buat middleware sendiri pada setiap service. Hal ini agar setiap middleware dapat di custom pada setiap service. Kemudian didalamnya, call gRPC client `ValidateToken()`
    ```go
        req := &pb.ValidateTokenRequest{Token: token} 
		res, err := m.authClient.ValidateToken(ctx, req)
		if err != nil || res == nil || !res.Valid {
			logrus.Error(err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}
    ```
3. Lalu gunakan pada echo middleware
    ```go
        protected := e.Group("/products")
	    protected.Use(JWTMiddleware)
    ```
4. Buat helper untuk get user session. (Gunakan pada level handler/delivery)
    ```go
        func GetUserSession(c echo.Context) *model.User {
            u := c.Get("user")
            if u == nil {
                return nil
            }

            user, ok := u.(*model.User)
            if !ok {
                return nil
            }

            return user
        }
    ```

    Implementasi
    ```go
        userSession := helper.GetUserSession(c)
        if userSession == nil {
            return c.JSON(http.StatusUnauthorized, response{
                Status:  http.StatusUnauthorized,
                Message: "unauthorized",
            })
        }
    ```

Contoh lebih lanjut ada pada folder `middleware-example`.

