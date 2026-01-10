import axios from "axios";

export const api = axios.create({
  baseURL: "https://cinema.localhost/api",
  withCredentials: true,
});
