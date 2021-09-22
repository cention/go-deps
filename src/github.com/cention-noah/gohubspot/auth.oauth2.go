package gohubspot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// OAuth2 OAuth 2 hubspot authenticator
type OAuth2 struct {
	token string
}

type OAuthsService service

type Settings struct {
	ClientID           string
	ClientSecret       string
	AuthorizationCode  string
	RedirectURI        string
	Scopes             string
	RefreshToken       string
	RequestCredentials *RequestCredentials
}

type OAuthResponse struct {
	AccessToken   string  `json:"access_token"`
	RefreshToken  string  `json:"refresh_token"`
	ExpiresIn     float64 `json:"expires_in"`
	TokenType     string  `json:"token_type"`
	Status        string  `json:"status"`
	Message       string  `json:"message"`
	CorrelationID string  `json:"correlationId"`
	RequestID     string  `json:"requestId"`
}

type RequestCredentials struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
	AccessTokenUpdating  sync.Mutex
}

type TokenUser struct {
	Token     string   `json:"token"`
	UserEmail string   `json:"user"`
	HubDomain string   `json:"hub_domain"`
	Scopes    []string `json:"scopes"`
	HubID     int      `json:"hub_id"`
	AppID     int      `json:"app_id"`
	ExpiresIn float64  `json:"expires_in"`
	UserID    int      `json:"user_id"`
	TokenType string   `json:"token_type"`
}

// NewOAuth2 Create new instance of OAuth2
func NewOAuth2(token string) OAuth2 {
	return OAuth2{token: token}
}

func NewSettings(clientID, clientSecret, redirectURI string, scopes []string) *Settings {
	scp := strings.Join(scopes, ",")
	return &Settings{
		ClientID:           clientID,
		ClientSecret:       clientSecret,
		RedirectURI:        redirectURI,
		RequestCredentials: &RequestCredentials{},
		Scopes:             scp,
	}
}

// Authenticate auth with OAuth2
func (auth OAuth2) Authenticate(request *http.Request) error {

	request.Header.Set("Authorization", "Bearer "+auth.token)

	return nil
}

//Authorization Request an authorization code
//The authorization code flow begins with the client directing the user to the /authorize endpoint.
//
//https://legacydocs.hubspot.com/docs/methods/oauth2/initiate-oauth-integration
func (s *Settings) Authorization() string {
	return fmt.Sprintf("https://app.hubspot.com/oauth/authorize?client_id=%v&redirect_uri=%v&scope=%v", s.ClientID, s.RedirectURI, strings.ReplaceAll(s.Scopes, ",", "%20"))
}

func (s *Settings) InitializeCredentials() error {
	err := s.setAccessToken()
	return err
}

func (s *Settings) setAccessToken() error {
	if s.AuthorizationCode == "" {
		return fmt.Errorf("no access code found in client")
	}
	s.RequestCredentials.AccessTokenUpdating.Lock()
	defer s.RequestCredentials.AccessTokenUpdating.Unlock()
	if s.RequestCredentials.AccessToken != "" && s.RequestCredentials.AccessTokenExpiresAt.After(time.Now()) {
		return nil
	}
	turl := "https://api.hubapi.com/oauth/v1/token"
	tokenURI, err := url.Parse(turl)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(tokenURI.String(), url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {s.ClientID},
		"client_secret": {s.ClientSecret},
		"redirect_uri":  {s.RedirectURI},
		"code":          {s.AuthorizationCode},
	})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	parsedResp := OAuthResponse{}
	err = json.Unmarshal(b, &parsedResp)
	if err != nil {
		return err
	}
	if parsedResp.Status != "" {
		return fmt.Errorf("%v: %v", parsedResp.Status, parsedResp.Message)
	}
	s.RefreshToken = parsedResp.RefreshToken
	expiresAt := time.Now().Add(time.Duration(parsedResp.ExpiresIn) * time.Second)
	s.RequestCredentials.AccessToken = parsedResp.AccessToken
	s.RequestCredentials.AccessTokenExpiresAt = expiresAt
	return nil
}

func (s *Settings) RefreshCredentials() error {
	if s.RefreshToken == "" {
		return fmt.Errorf("no refresh token found in client. call InitializeCredentials to fill this")
	}
	s.RequestCredentials.AccessTokenUpdating.Lock()
	defer s.RequestCredentials.AccessTokenUpdating.Unlock()

	turl := "https://api.hubapi.com/oauth/v1/token"
	tokenURI, err := url.Parse(turl)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(tokenURI.String(), url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {s.ClientID},
		"client_secret": {s.ClientSecret},
		"refresh_token": {s.RefreshToken},
	})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	parsedResp := OAuthResponse{}
	err = json.Unmarshal(b, &parsedResp)
	if err != nil {
		return err
	}
	if parsedResp.Status != "" {
		return fmt.Errorf("%v: %v", parsedResp.Status, parsedResp.Message)
	}
	s.RefreshToken = parsedResp.RefreshToken
	expiresAt := time.Now().Add(time.Duration(parsedResp.ExpiresIn) * time.Second)
	s.RequestCredentials.AccessToken = parsedResp.AccessToken
	s.RequestCredentials.AccessTokenExpiresAt = expiresAt
	return nil
}

func (s *OAuthsService) GetByToken(token string) (*TokenUser, error) {
	url := fmt.Sprintf("/oauth/v1/access-tokens/%s", token)
	res := new(TokenUser)
	err := s.client.RunGet(url, res)
	return res, err
}
