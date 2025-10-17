import axios from 'axios';

const API_URL = '/api';

export interface LoginResponse {
  token: string;
  expiresAt: string;
}

export const login = async (email: string, password: string): Promise<LoginResponse> => {
  const response = await axios.post(`${API_URL}/auth/login`, {
    email,
    password,
  });

  return response.data;
};