import { useState, useEffect, useCallback } from 'react'

export function useWebSocket(url) {
    const [socket, setSocket] = useState(null)
    const [isConnected, setIsConnected] = useState(false)
    const [error, setError] = useState(null)

    useEffect(() => {
        let ws = null
        let reconnectTimer = null

        const connect = () => {
            console.log('Attempting WebSocket connection to:', url)
            ws = new WebSocket(url)

            ws.onopen = () => {
                console.log('WebSocket connected successfully')
                setIsConnected(true)
                setError(null)
            }

            ws.onclose = (event) => {
                console.log('WebSocket disconnected:', event.code, event.reason)
                setIsConnected(false)
                // Try to reconnect after 3 seconds
                reconnectTimer = setTimeout(connect, 3000)
            }

            ws.onerror = (error) => {
                console.error('WebSocket error:', error)
                setError('WebSocket error: ' + error.message)
            }

            setSocket(ws)
        }

        connect()

        // Cleanup on unmount
        return () => {
            if (ws) {
                ws.close()
            }
            if (reconnectTimer) {
                clearTimeout(reconnectTimer)
            }
        }
    }, [url])

    const sendMessage = useCallback((type, payload) => {
        if (socket && isConnected) {
            socket.send(JSON.stringify({ type, payload }))
        }
    }, [socket, isConnected])

    return { socket, isConnected, error, sendMessage }
} 