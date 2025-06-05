    import { useEffect, useRef, useState } from "react";
    import axios from "axios";

    const Chat = () => {
        const token = localStorage.getItem("token");
        const [messages, setMessages] = useState([]);
        const [input, setInput] = useState("");
        const ws = useRef(null);
        const user_id = 2;
        const admin_id = 1;

        const parseToken = () => {
            if (token) {
                try {
                    const decodedToken = JSON.parse(atob(token.split('.')[1]));
                    return decodedToken.role;
                } catch (error) {
                    console.error("Error decoding token:", error);
                    return null;
                }
            }
            return null;
        };

        const currentRole = parseToken();

        useEffect(() => {
            const fetchPreviousMessages = async () => {
                try {
                    const response = await axios.get(`http://localhost:3000/api/chat/user?id=${user_id}`, {
                        headers: { "Authorization": `Bearer ${token}` }
                    });
                    setMessages(response.data.data);
                } catch (error) {
                    console.error("Error fetching messages:", error);
                }
            };

            fetchPreviousMessages();
        }, [user_id, token]);

        useEffect(() => {
            if (ws.current) return;

            ws.current = new WebSocket(`ws://localhost:3000/ws/chat?user_id=${user_id}&admin_id=${admin_id}`);

            ws.current.onopen = () => console.log("WebSocket Connected!");

            ws.current.onmessage = (event) => {
                try {
                    const receivedMessage = JSON.parse(event.data);
                    console.log("Received Data:", receivedMessage, "Current Role:", currentRole);

                    setMessages((prev) => {
                        if (!prev.some(msg => msg.message === receivedMessage.message && msg.user_id === receivedMessage.user_id)) {
                            return [...prev, { ...receivedMessage }];
                        }
                        return [...prev];
                    });

                } catch (error) {
                    console.error("Error parsing message:", error);
                }
            };

            ws.current.onerror = (error) => console.error("WebSocket Error:", error);

            return () => {
                if (ws.current) {
                    ws.current.close();
                    ws.current = null;
                }
            };
        }, [user_id, admin_id]);

        const sendMessage = async () => {
            if (!ws.current || !input.trim()) return;

            const messageData = {
                user_id,
                admin_id,
                role_message: currentRole,
                message: input
            };

            try {
                await axios.post(
                    `http://localhost:3000/api/chat/user?id=${user_id}`,
                    messageData,
                    { headers: { "Authorization": `Bearer ${token}` } }
                );
                setInput("");
            } catch (error) {
                console.error("Error sending message:", error);
            }

        };

        return (
            <div>
                <h2>Chat</h2>
                <div>
                    {messages.map((msg) => (
                        <div key={msg.id || `${msg.user_id}-${msg.message}`} style={{
                            marginBottom: "10px",
                            padding: "5px",
                            border: "1px solid #ddd",
                            display: "flex",
                            justifyContent: String(msg.role) === String(currentRole) ? "flex-end" : "flex-start",
                        }}>
                            <div>
                                <p><strong>{msg.user_id}:</strong> {msg.message}</p>
                                {msg.image_path && (
                                    <img
                                        src={msg.image_path}
                                        alt="Chat Image"
                                        style={{ width: "200px", borderRadius: "5px" }}
                                    />
                                )}
                            </div>
                        </div>
                    ))}
                </div>
                <input
                    type="text"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    placeholder="Type a message..."
                />
                <button onClick={sendMessage}>Send</button>
            </div>
        );
    };

    export default Chat;
