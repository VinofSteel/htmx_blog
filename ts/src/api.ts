
import axios, { AxiosInstance } from "axios"

export const API: AxiosInstance = axios.create({
    baseURL: window.location.origin,
    timeout: 5000,
    headers: {
        "Content-Type": "application/json",
        // Authorization: `Bearer ${token}`
    }
});