import { API_URL } from "./constants";
import { User } from "./types";

export default async function verifyUser() : Promise<User | null>{
  const token = localStorage.getItem('token');
  if (token) {
    const response = await fetch(`${API_URL}/verify`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });
    const data = await response.json();
    if (data.success) {
      return data.user;
    }
  }
  return null;
}
