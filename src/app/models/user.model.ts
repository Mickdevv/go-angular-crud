// user.model.ts
export interface UserWithToken {
  username: string;
  access: string;
  refresh: string;
}

export interface UserLoginRequest {
  username: string;
  password: string;
}

export interface UserRegisterRequest {
  username: string;
  password1: string;
  password2: string;
}
