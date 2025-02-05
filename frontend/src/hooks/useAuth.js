import { useState, useEffect, useCallback } from 'react'

export function useAuth() {
    const [user, setUser] = useState(null)
    const [token, setToken] = useState(localStorage.getItem('token'))
    const [loading, setLoading] = useState(true)

    const login = useCallback(async (username, password) => {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        })

        if (!response.ok) {
            throw new Error('Login failed')
        }

        const data = await response.json()
        localStorage.setItem('token', data.token)
        localStorage.setItem('user', JSON.stringify(data.user))
        setToken(data.token)
        setUser(data.user)
    }, [])

    const logout = useCallback(() => {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        setToken(null)
        setUser(null)
    }, [])

    const refreshToken = useCallback(async () => {
        try {
            const response = await fetch('/api/auth/refresh', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                },
            })

            if (!response.ok) {
                throw new Error('Token refresh failed')
            }

            const data = await response.json()
            localStorage.setItem('token', data.token)
            localStorage.setItem('user', JSON.stringify(data.user))
            setToken(data.token)
            setUser(data.user)
        } catch (error) {
            logout()
        }
    }, [token, logout])

    useEffect(() => {
        if (token) {
            const storedUser = localStorage.getItem('user')
            if (storedUser) {
                setUser(JSON.parse(storedUser))
            }
            // Refresh token every 23 hours
            const interval = setInterval(refreshToken, 23 * 60 * 60 * 1000)
            return () => clearInterval(interval)
        }
        setLoading(false)
    }, [token, refreshToken])

    return { user, token, login, logout, refreshToken, loading }
} 