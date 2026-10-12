## Defining "Grant" in OAuth 2.0  

An Authorization Grant is defined in the OAuth 2.0 specification (RFC 6749) as:  
Authorization Grant: A credential representing the resource owner's authorization (permission) used by the client to obtain an access token.  

To understand why a Grant is distinct from an Access Token, consider the lifecycle of an OAuth flow:  

```
[ User Permission ] ──> [ Authorization Grant ] ──> [ Access Token ] ──> [ Resource API ]                             (Exchange Phase)           (Consumption Phase)
```


1. User Permission: The user agrees to give the app access.  
2. Authorization Grant: An intermediate claim ticket issued to the client as proof that the user said "yes" (e.g., an Authorization Code or a JWT Assertion).  
3. Access Token: The final bearer key issued by the Authorization Server (AS) after consuming the Grant. The client uses this token to call protected APIs.  

Because the Grant is consumed to _obtain_ the Access Token, the Access Token itself cannot be the Grant.

The name of the grant type—such as `authorization_code`, `client_credentials`, or `urn:ietf:params:oauth:grant-type:jwt-bearer`—describes how the client obtained permission (the input credential/mechanism presented to the Token Endpoint).

### Example 

#### Step 1: User Permission  
The user clicks "Login with OAuth Server" on a client app, and the app redirects the user's browser to the Authorization Server (AS) to ask for consent.  

```http
GET /authorize?response_type=code
&client_id=my_photo_app_123
&redirect_uri=https%3A%2F%2Fphotoapp.com%2Fcallback  &scope=photos.read  
&state=xyz123 HTTP/1.1
Host: auth.example.com
```

- What happens: The user sees a consent page asking: _"Do you give PhotoApp permission to access your photos?"_ The user clicks "Allow".

#### Step 2: Authorization Grant  
The AS issues an intermediate claim ticket—the Authorization Code—and sends it back to the client via a browser redirect.  

```http
HTTP/1.1 302 FoundLocation: https://photoapp.com/callback?code=AUTH_CODE_987654321&state=xyz123
```

- The Artifact: `code=AUTH_CODE_987654321` is the Authorization Grant.  
- Key Detail: It cannot be used directly to access photos. It is short-lived (usually expires in 1–10 minutes) and meant strictly to be exchanged at the token endpoint.

#### Step 3: Access Token Exchange  
The client app makes a back-channel POST request to exchange the Authorization Grant for an Access Token.  

```http
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=AUTH_CODE_987654321
&redirect_uri=https%3A%2F%2Fphotoapp.com%2Fcallback
&client_id=my_photo_app_123
&client_secret=SUPER_SECRET_KEY
```

- Notice how `grant_type=authorization_code` tells the AS _what kind_ of grant is being presented.

The AS validates the code, consumes it so it can never be used again, and responds with the Access Token:  

```http
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Cache-Control: no-store

{  
"access_token": "eyJhbGciOi...",  
"token_type": "Bearer",  
"expires_in": 3600,  
"scope": "photos.read"
}
```

Step 4: Consuming the Access Token  
Now that the client has the final key (access_token), it can query the Resource API.  

```http
GET /v1/user/photos HTTP/1.1
Host: api.example.com
Authorization: Bearer eyJhbGciOi...
```