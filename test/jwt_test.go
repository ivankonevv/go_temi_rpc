package test

import (
	"reflect"
	"temi_rpc/pkg/services/models"
	"temi_rpc/tools"
	"testing"
	"time"
)

func TestJWTManager_Generate(t *testing.T) {
	type fields struct {
		secretKey     string
		tokenDuration time.Duration
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test with valid user data",
			fields: fields{
				secretKey:     "secret",
				tokenDuration: 10 * time.Second,
			},
			args: args{user: &models.User{
				Email:     "testuser@example.com",
				FirstName: "test",
				LastName:  "user",
				Role:      "admin",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &tools.JWTManager{
				SecretKey:     tt.fields.secretKey,
				TokenDuration: tt.fields.tokenDuration,
			}
			got, err := manager.Generate(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			claims, err := manager.Verify(got)
			if err != nil {
				t.Errorf("Generate() error = %v, wantErr %v", err, err)
				return
			}
			if claims.Email != tt.args.user.Email || claims.Role != tt.args.user.Role {
				t.Errorf("Generate() got email = %v, want %v. Got role = %v, want %v",
					claims.Email,
					tt.args.user.Email,
					claims.Role,
					tt.args.user.Role,
				)
				return
			}
		})
	}
}

func TestJWTManager_Verify(t *testing.T) {
	// TODO: implement
	manager := &tools.JWTManager{
		SecretKey:     "testKey",
		TokenDuration: 30 * time.Second,
	}
	user := models.User{
		Email: "test@example.com",
		Role:  "test",
	}
	token, _ := manager.Generate(&user)
	type fields struct {
		SecretKey     string
		TokenDuration time.Duration
	}
	type args struct {
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *tools.UserClaims
		wantErr bool
	}{
		{
			name: "Test with valid token",
			fields: fields{
				SecretKey:     "testKey",
				TokenDuration: 30 * time.Second,
			},
			args: args{token},
			want: &tools.UserClaims{
				Email: "test@example.com",
				Role:  "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &tools.JWTManager{
				SecretKey:     tt.fields.SecretKey,
				TokenDuration: tt.fields.TokenDuration,
			}
			got, err := manager.Verify(tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}
