import { API_URL } from "./constants";
import { User } from "./types";

export default async function createUser(user: User): Promise<User | null> {
  const response = await fetch(`${API_URL}/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  const data = await response.json();

  if (data.token) {
    localStorage.setItem("token", data.token);
    const user :User = {
        name : data.user.Name,
        address : data.user.Address
    }

    localStorage.setItem("user", JSON.stringify(user));
  }
  return data.user;
}
